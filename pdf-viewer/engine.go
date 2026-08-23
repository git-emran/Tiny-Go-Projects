package main

import (
	"fmt"
	"image"
	"time"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
	"github.com/klippa-app/go-pdfium/webassembly"
)

type renderedPage struct {
	Image  image.Image
	Width  int
	Height int
}

type Engine struct {
	pool      pdfium.Pool
	insatance pdfium.Pdfium
	doc       *responses.OpenDocument
}

const maxRenderPixels = 8000 * 8000

func newEngine() (*Engine, error) {
	pool, err := webassembly.Init(webassembly.Config{
		MinIdle:  1,
		MaxIdle:  2,
		MaxTotal: 4,
	})
	if err != nil {
		return nil, fmt.Errorf("init pdfium pool: %w", err)
	}
	inst, err := pool.GetInstance(30 * time.Second)
	if err != nil {
		return nil, fmt.Errorf("pdfium instance: %w", err)
	}

	return &Engine{pool: pool, insatance: inst}, nil
}

func (e *Engine) OpenFile(path string) (pageCount int, err error) {
	data, err := readFileBytes(path)
	if err != nil {
		return 0, err
	}

	doc, err := e.insatance.OpenDocument(&requests.OpenDocument{File: &data})
	if err != nil {
		return 0, fmt.Errorf("open document : %w", err)
	}
	e.doc = doc
	pageInfo, err := e.insatance.FPDF_GetPageCount(&requests.FPDF_GetPageCount{Document: doc.Document})
	if err != nil {
		return 0, fmt.Errorf("page count : %w", err)
	}
	return pageInfo.PageCount, nil
}

func (e *Engine) RenderPage(pageIndex int, scale float32) (*renderedPage, error) {
	page := requests.Page{
		ByIndex: &requests.PageByIndex{
			Document: e.doc.Document,
			Index:    pageIndex,
		},
	}
	sizeRes, err := e.insatance.GetPageSize(&requests.GetPageSize{Page: page})
	if err != nil {
		return nil, fmt.Errorf("get page size %d: %w", pageIndex, err)
	}

	pageHeightPt := sizeRes.Height
	pageWidthPt := sizeRes.Width

	effectivePixels := pageWidthPt * float64(scale) * pageHeightPt * float64(scale)
	if effectivePixels > maxRenderPixels {
		return nil, fmt.Errorf("requested render size exceeds limit")
	}

	res, err := e.insatance.RenderPageInDPI(&requests.RenderPageInDPI{
		Page: page,
		DPI:  int(72 * scale),
	})
	if err != nil {
		return nil, fmt.Errorf("render page %d: %w", pageIndex, err)
	}
	defer res.Cleanup()
	return &renderedPage{
		Image:  res.Result.RenderedImage,
		Width:  res.Result.RenderedImage.Bounds().Dx(),
		Height: res.Result.RenderedImage.Bounds().Dy(),
	}, nil
}

func (e *Engine) Close() {
	if e.doc != nil {
		e.insatance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{Document: e.doc.Document})
		e.insatance.Close()
		e.pool.Close()
	}
}
