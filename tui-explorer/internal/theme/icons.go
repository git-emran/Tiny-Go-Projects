package theme

import (
	"path/filepath"
	"strings"

	"github.com/git-emran/tiny-go-project/tui-explorer/internal/fs"
)

type IconSet int

const (
	IconSetPlain = iota
	IconSetnerd
)

var nerdIconsByExt = map[string]string{
	".go":   "\uE627",
	".mod":  "\uE627",
	".py":   "\uE73C",
	".js":   "\uE74E",
	".ts":   "\uE628",
	".rs":   "\uE7A8",
	".md":   "\uF48A",
	".json": "\uE60B",
	".yaml": "\uF481",
	".yml":  "\uF481",
	".toml": "\uF481",
	".sh":   "\uF489",
	".png":  "\uF1C5",
	".jpg":  "\uF1C5",
	".jpeg": "\uF1C5",
	".gif":  "\uF1C5",
	".pdf":  "\uF1C1",
	".zip":  "\uF410",
	".tar":  "\uF410",
	".gz":   "\uF410",
}

const (
	nerdFolder      = "\uF07B"
	nerdFolderOpen  = "\uF07C"
	nerdFileDefault = "\uF15B"

	plainFolder      = "[D]"
	plainFiledefault = "[F]"
)

func IconFor(set IconSet, e fs.Entry, selected bool) string {
	if set == IconSetPlain {
		if e.IsDir {
			return plainFolder + " "
		}
		return plainFiledefault + " "
	}

	if e.IsDir {
		if selected {
			return nerdFolderOpen + " "
		}

		return nerdFolder + " "
	}
	ext := strings.ToLower(filepath.Ext(e.Name))
	if icon, ok := nerdIconsByExt[ext]; ok {
		return icon + " "
	}
	return nerdFileDefault + " "
}
