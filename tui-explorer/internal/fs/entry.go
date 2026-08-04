// Package fs is an entry point for the application where directories are read.
package fs

import (
	"os"
	"sort"
	"strings"
)

type Entry struct {
	Name  string
	IsDir bool
}

func ReadDir(path string) []Entry {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var out []Entry
	for _, f := range files {
		out = append(out, Entry{Name: f.Name(), IsDir: f.IsDir()})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].IsDir != out[j].IsDir {
			return out[i].IsDir
		}
		return strings.ToLower(out[i].Name) < strings.ToLower(out[j].Name)
	})

	return out
}
