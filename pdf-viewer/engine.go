package main

import (
	"fmt"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/responses"
	pdfium_single "github.com/klippa-app/go-pdfium/single_threaded"
)

type Engine struct {
	pool      pdfium.Pool
	insatance pdfium.Pdfium
	doc       *responses.OpenDocument
}

func newEngine() (*Engine, error) {
	pool := pdfium_single.Init(pdfium_single.Config{})
	inst, err := pool.GetInstance(0)
	if err != nil {
		return nil, fmt.Errorf("pdfium instance: %w", err)
	}

	return &Engine{pool: pool, insatance: inst}, nil
}

func (e *Engine) OpenFile(path string) (pageCount int, err error) {
	data, err := readFileBytes(path)
}
