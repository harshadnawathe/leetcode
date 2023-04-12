package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Mising argument: Problem Url")
		os.Exit(1)
	}

	url := os.Args[1]
	problemSpec, err := problemSpecFrom(url)
	if err != nil {
		fmt.Println("Cannot parse url: ", err)
		os.Exit(1)
	}

	err = os.Mkdir(problemSpec.TitleSlug, 0o750)
	if err != nil {
		fmt.Println("Cannot create solution directory: ", err)
		os.Exit(1)
	}

	if err = createReadMeFile(problemSpec); err != nil {
		fmt.Println("Cannot create README: ", err)
		os.Exit(1)
	}
	if err = createSolutionFile(problemSpec); err != nil {
		fmt.Println("Cannot create solution.go: ", err)
		os.Exit(1)
	}
	if err = createSolutionTestFile(problemSpec); err != nil {
		fmt.Println("Cannot create solution_test.go: ", err)
		os.Exit(1)
	}
}

// ProblemSpec
type ProblemSpec struct {
	Id          string `json:"questionFrontendId"`
	Title       string `json:"title"`
	TitleSlug   string `json:"titleSlug"`
	DefaultCode string `json:"code"`
}

func (ps *ProblemSpec) PackageName() string {
	re := regexp.MustCompile(`\W|^\d`)
	packageName := re.ReplaceAllString(ps.TitleSlug, "")
	return packageName
}

func problemSpecFrom(url string) (*ProblemSpec, error) {
	var err error

	var response *http.Response
	response, err = http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var responseBody []byte
	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	questionRe := regexp.MustCompile(`{"questionId":.*?}`)
	questionJson := questionRe.Find(responseBody)
	codeRe := regexp.MustCompile(`{"lang":"Go",.*?"}`)
	codeJson := codeRe.Find(responseBody)

	problemSpec := &ProblemSpec{}

	err = json.Unmarshal(questionJson, problemSpec)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(codeJson, problemSpec)
	if err != nil {
		return nil, err
	}

	return problemSpec, nil
}

// ReadMeFile
var (
	readMeFileName     = "README.md"
	readMeTemplateText = `# {{ .Id }}. {{ .Title }}

Click [here](https://leetcode.com/problems/{{.TitleSlug}}/) for the leetcode problem statement.
`

	readMeTemplate = template.Must(template.New(readMeFileName).Parse(readMeTemplateText))
)

func createReadMeFile(problemSpec *ProblemSpec) error {
	readMeFile, err := os.Create(filepath.Join(problemSpec.TitleSlug, readMeFileName))
	if err != nil {
		return err
	}
	defer readMeFile.Close()

	err = readMeTemplate.Execute(readMeFile, problemSpec)
	if err != nil {
		return err
	}

	return nil
}

// Solution File
var (
	solutionFileName     = "solution.go"
	solutionTemplateText = `package {{ .PackageName }}

{{ .DefaultCode }}
`
	solutionTemplate = template.Must(template.New(solutionFileName).Parse(solutionTemplateText))
)

func createSolutionFile(problemSpec *ProblemSpec) error {
	solutionFile, err := os.Create(filepath.Join(problemSpec.TitleSlug, solutionFileName))
	if err != nil {
		return err
	}
	defer solutionFile.Close()

	err = solutionTemplate.Execute(solutionFile, problemSpec)
	if err != nil {
		return nil
	}

	return nil
}

// Solution spec
type SolutionSpec struct {
	Package string
	Func    *FuncSpec
}

type FuncSpec struct {
	Name   string
	Result string
	Params []struct {
		Name, Type string
	}
}

func (fs *FuncSpec) TestName() string {
	return fmt.Sprintf("Test%s%s", strings.ToUpper(fs.Name[0:1]), fs.Name[1:])
}

func solutionSpecFrom(solutionFileName string) (*SolutionSpec, error) {
	solutionFile, err := os.Open(solutionFileName)
	if err != nil {
		return nil, err
	}
	defer solutionFile.Close()

	src, err := io.ReadAll(solutionFile)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, solutionFileName, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	spec := &SolutionSpec{}

	spec.Package = f.Name.Name

	for _, decl := range f.Decls {
		if spec.Func != nil {
			break
		}

		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv != nil {
			// not a function declaration or a method
			continue
		}

		// Result check
		if fn.Type.Results == nil || len(fn.Type.Results.List) > 1 {
			// Returns no result or more than one result
			continue
		}

		// Params check
		if len(fn.Type.Params.List) == 0 {
			// Has no parameters
			continue
		}

		// Func name capture
		spec.Func = &FuncSpec{Name: fn.Name.Name}

		// Params capture
		for i, param := range fn.Type.Params.List {
			var paramName string
			if param.Names == nil {
				paramName = fmt.Sprintf("arg%v", i)
			} else {
				paramName = param.Names[0].Name
			}

			typeStart := fset.Position(param.Type.Pos()).Offset
			typeEnd := fset.Position(param.Type.End()).Offset
			paramType := string(src[typeStart:typeEnd])

			spec.Func.Params = append(spec.Func.Params, struct{ Name, Type string }{paramName, paramType})
		}

		// Result capture
		result := fn.Type.Results.List[0].Type

		typeStart := fset.Position(result.Pos()).Offset
		typeEnd := fset.Position(result.End()).Offset
		spec.Func.Result = string(src[typeStart:typeEnd])
	}

	return spec, nil
}

// Solution Test File

var (
	solutionTestFileName     = "solution_test.go"
	solutionTestTemplateText = `package {{ .Package }}

import (
	"fmt"
	"reflect"
	"testing"
)

{{- if .Func }}
func {{ .Func.TestName }}(t *testing.T) {
{{- else }}
func TestSolution(t *testing.T) {
{{- end }}

  tests := []struct {
		{{- if .Func }}
		{{- range .Func.Params }}
		{{ .Name }} {{ .Type }}
		{{- end }}
		{{- else }}
		input int
		{{- end}}
		{{- if .Func }}
		want {{ .Func.Result }}
		{{- else}}
		want  int
		{{- end }}
	}{}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			{{- if .Func }}
			got := {{ .Func.Name }}(
				{{- range .Func.Params }}
				test.{{ .Name }},
				{{- end }}
			)
			{{- else }}
			got := test.input
			{{- end }}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
`

	solutionTestTemplate = template.Must(template.New(solutionTestFileName).Parse(solutionTestTemplateText))
)

func createSolutionTestFile(problemSpec *ProblemSpec) error {
	solutionFilePath := filepath.Join(problemSpec.TitleSlug, solutionFileName)
	spec, err := solutionSpecFrom(solutionFilePath)
	if err != nil {
		return err
	}

	solutionTestFile, err := os.Create(filepath.Join(problemSpec.TitleSlug, solutionTestFileName))
	if err != nil {
		return err
	}
	defer solutionTestFile.Close()

	err = solutionTestTemplate.Execute(solutionTestFile, spec)
	if err != nil {
		return nil
	}

	return nil
}
