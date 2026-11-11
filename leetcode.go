package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"
)

// The script exists to remove boilerplate, so it degrades instead of giving
// up: whatever cannot be fetched or parsed is replaced by a placeholder and
// reported on stderr. The only fatal errors are the ones that leave us
// without a directory to write into.
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing argument: problem url or title slug")
		fmt.Fprintln(os.Stderr, "Usage: leetcode <https://leetcode.com/problems/two-sum/|two-sum>")
		os.Exit(1)
	}

	titleSlug, err := titleSlugFrom(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot determine problem slug:", err)
		os.Exit(1)
	}

	problemSpec := problemSpecFor(titleSlug)

	err = os.Mkdir(problemSpec.TitleSlug, 0o750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		fmt.Fprintln(os.Stderr, "Cannot create solution directory:", err)
		os.Exit(1)
	}

	// One failing file must not cost us the other two.
	steps := []struct {
		fileName string
		create   func(*ProblemSpec) error
	}{
		{readMeFileName, createReadMeFile},
		{solutionFileName, createSolutionFile},
		{solutionTestFileName, createSolutionTestFile},
	}

	failed := false
	for _, step := range steps {
		if err := step.create(problemSpec); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot create %s: %v\n", step.fileName, err)
			failed = true
		}
	}

	if failed {
		os.Exit(1)
	}
}

// createFile opens a file in the solution directory, refusing to clobber what
// is already there. A file that exists is skipped, not an error: rerunning on
// a half-done problem should fill in the gaps and leave solved code alone.
// A nil file with a nil error means "skipped".
func createFile(problemSpec *ProblemSpec, fileName string) (*os.File, error) {
	path := filepath.Join(problemSpec.TitleSlug, fileName)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o640)
	if errors.Is(err, fs.ErrExist) {
		fmt.Println("Keeping existing", path)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	fmt.Println("Created", path)

	return file, nil
}

// writeGoFile renders a template into a go file through gofmt, so the whole
// scaffolding comes out formatted the way the toolchain would write it.
// Output that cannot be formatted is still worth writing.
func writeGoFile(file *os.File, goTemplate *template.Template, data any) error {
	var rendered bytes.Buffer
	if err := goTemplate.Execute(&rendered, data); err != nil {
		return err
	}

	formatted, err := format.Source(rendered.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot gofmt %s: %v\n", file.Name(), err)
		formatted = rendered.Bytes()
	}

	_, err = file.Write(formatted)

	return err
}

// ProblemSpec
type ProblemSpec struct {
	Id          string
	Title       string
	TitleSlug   string
	DefaultCode string
}

func (ps *ProblemSpec) PackageName() string {
	re := regexp.MustCompile(`\W|^\d`)
	packageName := re.ReplaceAllString(ps.TitleSlug, "")
	return packageName
}

func (ps *ProblemSpec) Url() string {
	return problemUrlPrefix + ps.TitleSlug + "/"
}

// Stub is the body of solution.go: the snippet leetcode gave us, or somewhere
// to paste it when the fetch degraded.
func (ps *ProblemSpec) Stub() string {
	if ps.DefaultCode != "" {
		return ps.DefaultCode
	}
	return "// TODO: paste the Go code snippet from " + ps.Url()
}

const problemUrlPrefix = "https://leetcode.com/problems/"

var titleSlugRe = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

// titleSlugFrom takes the slug from a problem url, or accepts a bare slug so
// the scaffolding can still be made when leetcode is unreachable.
func titleSlugFrom(arg string) (string, error) {
	if titleSlugRe.MatchString(arg) {
		return arg, nil
	}

	parsedUrl, err := url.Parse(arg)
	if err != nil {
		return "", err
	}

	// .../problems/<slug>[/description/...]
	parts := strings.Split(strings.Trim(parsedUrl.Path, "/"), "/")
	if len(parts) < 2 || parts[0] != "problems" || !titleSlugRe.MatchString(parts[1]) {
		return "", fmt.Errorf("not a leetcode problem url: %s", arg)
	}

	return parts[1], nil
}

var romanNumeralRe = regexp.MustCompile(`^[ivx]+$`)

// titleFrom reconstructs a readable title from the slug alone, for when the
// real one could not be fetched: "meeting-rooms-iii" -> "Meeting Rooms III".
func titleFrom(titleSlug string) string {
	words := strings.Split(titleSlug, "-")
	for i, word := range words {
		switch {
		case word == "":
		case romanNumeralRe.MatchString(word):
			words[i] = strings.ToUpper(word)
		default:
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// problemSpecFor always returns a usable spec. If leetcode cannot be reached
// the spec falls back to what the slug alone tells us, so the directory, the
// README and the test skeleton still get written and only the stub has to be
// pasted in by hand.
func problemSpecFor(titleSlug string) *ProblemSpec {
	problemSpec, err := fetchProblemSpec(titleSlug)
	if err == nil {
		return problemSpec
	}

	fmt.Fprintln(os.Stderr, "Cannot fetch problem from leetcode:", err)
	fmt.Fprintln(os.Stderr, "Falling back to minimal scaffolding.")

	return &ProblemSpec{
		Title:     titleFrom(titleSlug),
		TitleSlug: titleSlug,
	}
}

// Leetcode serves the problem pages behind a bot challenge, so the question
// is read from the GraphQL api instead of scraped out of the html.
const (
	graphqlUrl   = "https://leetcode.com/graphql/"
	graphqlQuery = `query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionFrontendId
    title
    titleSlug
    isPaidOnly
    codeSnippets {
      langSlug
      code
    }
  }
}`
	goLangSlug = "golang"
)

type graphqlRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

type graphqlResponse struct {
	Data struct {
		Question *struct {
			Id           string `json:"questionFrontendId"`
			Title        string `json:"title"`
			TitleSlug    string `json:"titleSlug"`
			IsPaidOnly   bool   `json:"isPaidOnly"`
			CodeSnippets []struct {
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
			} `json:"codeSnippets"`
		} `json:"question"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func fetchProblemSpec(titleSlug string) (*ProblemSpec, error) {
	requestBody, err := json.Marshal(graphqlRequest{
		Query:     graphqlQuery,
		Variables: map[string]any{"titleSlug": titleSlug},
	})
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, graphqlUrl, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Referer", problemUrlPrefix+titleSlug+"/")

	client := &http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("leetcode graphql returned %s", response.Status)
	}

	var result graphqlResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("leetcode graphql error: %s", result.Errors[0].Message)
	}

	question := result.Data.Question
	if question == nil {
		return nil, fmt.Errorf("no such problem: %s", titleSlug)
	}

	problemSpec := &ProblemSpec{
		Id:        question.Id,
		Title:     question.Title,
		TitleSlug: question.TitleSlug,
	}

	for _, snippet := range question.CodeSnippets {
		if snippet.LangSlug == goLangSlug {
			problemSpec.DefaultCode = snippet.Code
			return problemSpec, nil
		}
	}

	// Everything but the stub was fetched, so report the gap and carry on
	// with what we have rather than throwing the question away.
	if question.IsPaidOnly {
		fmt.Fprintf(os.Stderr, "%s is a premium problem, no code snippet available.\n", titleSlug)
	} else {
		fmt.Fprintf(os.Stderr, "No Go code snippet for %s.\n", titleSlug)
	}

	return problemSpec, nil
}

// ReadMeFile
var (
	readMeFileName     = "README.md"
	readMeTemplateText = `# {{ with .Id }}{{ . }}. {{ end }}{{ .Title }}

Click [here]({{ .Url }}) for the leetcode problem statement.
`

	readMeTemplate = template.Must(template.New(readMeFileName).Parse(readMeTemplateText))
)

func createReadMeFile(problemSpec *ProblemSpec) error {
	readMeFile, err := createFile(problemSpec, readMeFileName)
	if err != nil || readMeFile == nil {
		return err
	}
	defer readMeFile.Close()

	return readMeTemplate.Execute(readMeFile, problemSpec)
}

// Solution File
var (
	solutionFileName     = "solution.go"
	solutionTemplateText = `package {{ .PackageName }}

{{ .Stub }}
`
	solutionTemplate = template.Must(template.New(solutionFileName).Parse(solutionTemplateText))
)

func createSolutionFile(problemSpec *ProblemSpec) error {
	solutionFile, err := createFile(problemSpec, solutionFileName)
	if err != nil || solutionFile == nil {
		return err
	}
	defer solutionFile.Close()

	return writeGoFile(solutionFile, solutionTemplate, problemSpec)
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

// TestName follows the gopls convention: Test_<name> for an unexported
// function, Test<Name> for an exported one. Leetcode stubs are always
// unexported, so in practice this is Test_mostBooked.
func (fs *FuncSpec) TestName() string {
	if ast.IsExported(fs.Name) {
		return "Test" + fs.Name
	}
	return "Test_" + fs.Name
}

// CallArgs renders the call under test the way gopls writes it, on one line:
// "tt.n, tt.meetings".
func (fs *FuncSpec) CallArgs() string {
	args := make([]string, len(fs.Params))
	for i, param := range fs.Params {
		args[i] = "tt." + param.Name
	}
	return strings.Join(args, ", ")
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
	"reflect"
	"testing"
)
{{ if .Func }}
func {{ .Func.TestName }}(t *testing.T) {
	tests := []struct {
		name string
{{- range .Func.Params }}
		{{ .Name }} {{ .Type }}
{{- end }}
		want {{ .Func.Result }}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := {{ .Func.Name }}({{ .Func.CallArgs }})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("{{ .Func.Name }}() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{- else }}
func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
{{- end }}
`

	solutionTestTemplate = template.Must(template.New(solutionTestFileName).Parse(solutionTestTemplateText))
)

func createSolutionTestFile(problemSpec *ProblemSpec) error {
	solutionTestFile, err := createFile(problemSpec, solutionTestFileName)
	if err != nil || solutionTestFile == nil {
		return err
	}
	defer solutionTestFile.Close()

	solutionFilePath := filepath.Join(problemSpec.TitleSlug, solutionFileName)
	spec, err := solutionSpecFrom(solutionFilePath)
	if err != nil {
		// Without a signature to mirror there is nothing to generate from,
		// but the generic table-driven skeleton is still worth having.
		fmt.Fprintln(os.Stderr, "Cannot read solution.go for test generation:", err)
		spec = &SolutionSpec{Package: problemSpec.PackageName()}
	}

	return writeGoFile(solutionTestFile, solutionTestTemplate, spec)
}
