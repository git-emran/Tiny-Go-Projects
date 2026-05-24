package main

import (
	"fmt"
	"os"
	"strings"

	clitodo "github.com/git-emran/tiny-go-projects/cli-todo"
)

const todoFileName = ".todo.json"

func main() {
	l := &clitodo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)

	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
