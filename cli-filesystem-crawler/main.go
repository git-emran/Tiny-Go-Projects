package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/git-emran/tiny-go-projects/cli-filesystem-crawler/walk"
)

type config struct {
	ext  string
	size int64
	list bool
}

func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if walk.FilterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			if cfg.list {
				return walk.ListFile(path, out)
			}

			return walk.ListFile(path, out)
		})
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "file extension to filter out")
	size := flag.Int64("size", 0, "minimum file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
