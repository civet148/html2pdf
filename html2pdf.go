package html2pdf

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/civet148/templatex"
)

type Option struct {
	PageWidth    uint   //页宽（默认A4宽度）
	PageHeight   uint   //页高（默认A4高度）
	PageSize     string //页面尺寸（例如A4)
	MarginTop    uint   //页边距上
	MarginBottom uint   //页边距下
	MarginLeft   uint   //页边距左
	MarginRight  uint   //页边距右
}

type Html2PDF struct {
	option Option
	wkpdf  *wkhtmltopdf.PDFGenerator
}

func NewHtml2PDF(options ...Option) *Html2PDF {
	var opt Option
	wkpdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}
	opt = getOption(wkpdf, options...)
	return &Html2PDF{
		wkpdf:  wkpdf,
		option: opt,
	}
}

// 通过模板文件生成PDF
func (h *Html2PDF) GenerateWithFile(strTemplatePath, strPdfPath string, data any) (err error) {
	var strTemplateContent string
	strTemplateContent, err = templatex.Generate(strTemplatePath, data)
	if err != nil {
		fmt.Printf("templatex generate error：%s\n", err)
		return
	}
	return h.GenerateWithContent(strTemplateContent, strPdfPath)
}

// 通过模板内容生成PDF
func (h *Html2PDF) GenerateWithContent(strTemplateContent string, strPdfPath string) (err error) {
	var htmlBuf = bytes.NewBufferString(strTemplateContent)
	// 从 HTML 内容创建页面
	page := wkhtmltopdf.NewPageReader(htmlBuf)

	// 将页面添加到 PDF 生成器
	h.wkpdf.AddPage(page)

	// 生成 PDF
	err = h.wkpdf.Create()
	if err != nil {
		return fmt.Errorf("create PDF error：%s", err)
	}

	// 写入 PDF 文件
	err = h.wkpdf.WriteFile(strPdfPath)
	if err != nil {
		return fmt.Errorf("write PDF to %s error：%s", strPdfPath, err)
	}
	return nil
}

func getOption(wkpdf *wkhtmltopdf.PDFGenerator, options ...Option) (opt Option) {
	if len(options) == 0 {
		wkpdf.PageSize.Set(wkhtmltopdf.PageSizeA4)
		opt.PageSize = wkhtmltopdf.PageSizeA4
	} else {
		opt = options[0]
		if opt.PageWidth != 0 && opt.PageHeight != 0 {
			wkpdf.PageWidth.Set(opt.PageWidth)
			wkpdf.PageHeight.Set(opt.PageHeight)
		}
		if opt.MarginTop != 0 {
			wkpdf.MarginTop.Set(opt.MarginTop)
		}
		if opt.MarginBottom != 0 {
			wkpdf.MarginBottom.Set(opt.MarginBottom)
		}
		if opt.MarginLeft != 0 {
			wkpdf.MarginLeft.Set(opt.MarginLeft)
		}
		if opt.MarginRight != 0 {
			wkpdf.MarginRight.Set(opt.MarginRight)
		}
		if opt.PageSize != "" {
			wkpdf.PageSize.Set(opt.PageSize)
		}
	}
	return opt
}

