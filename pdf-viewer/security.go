package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	maxFileSizeBytes = 500 << 20 // 500 MB
	pdfMagicBytes    = "%PDF-"
)

func readFileBytes(path string) ([]byte, error) {
	clean := filepath.Clean(path)
	info, err := os.Stat(clean)
	if err != nil {
		return nil, fmt.Errorf("stat file %w", err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("path is a directory")
	}
	if info.Size() > maxFileSizeBytes {
		return nil, fmt.Errorf("file exceeds the maximum size of %d bytes", maxFileSizeBytes)
	}

	data, err := os.ReadFile(clean)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
}
