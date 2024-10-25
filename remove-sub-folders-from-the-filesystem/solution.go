package removesubfoldersfromthefilesystem

import "strings"

// time: O(n * m)
// space: O(n * m) where n = num of folders, m = num of path elements
func removeSubfolders(folders []string) []string {
	dirTree := newDirTree()

	for _, folder := range folders {
		add(dirTree, folder)
	}

	return dfs(dirTree)
}

// Trie for detecting prefix

type dir struct {
	Subdirs map[string]*dir
	Name    string
	IsLast  bool
}

func newDir(val string) *dir {
	return &dir{
		Subdirs: make(map[string]*dir),
		Name:    val,
	}
}

func newDirTree() *dir {
	return newDir("")
}

func add(dir *dir, path string) {
	node := dir
	for _, x := range strings.Split(path, "/") {
		if len(x) == 0 {
			continue
		}
		if _, ok := node.Subdirs[x]; !ok {
			node.Subdirs[x] = newDir(x)
		}
		node = node.Subdirs[x]
	}
	node.IsLast = true
}

func dfs(dirTree *dir) []string {
	var xs []string

	type stkFrame struct {
		Dir  *dir
		Path string
	}

	stk := []stkFrame{{dirTree, ""}}

	for len(stk) > 0 {
		top := top(stk)
		stk = pop(stk)

		if top.Dir.IsLast {
			xs = append(xs, top.Path)
			continue
		}

		for _, child := range top.Dir.Subdirs {
			stk = push(stk, stkFrame{
				Dir:  child,
				Path: top.Path + "/" + child.Name,
			})
		}
	}

	return xs
}

// stack functions

func push[T any](stk []T, x T) []T {
	return append(stk, x)
}

func top[T any](stk []T) T {
	return stk[len(stk)-1]
}

func pop[T any](stk []T) []T {
	return stk[:len(stk)-1]
}
