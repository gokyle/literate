package main

// pandoc.go adds pandoc support for converting listings to PDF format.

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// pdfWriter transforms the listing to a PDF. First, the markdown is
// written to a temporary file (which is removed when the function
// returns); this temporary file is then passed to pandoc for conversion.
func pdfWriter(markdown, filename string, args []string) (err error) {
	tmp, err := ioutil.TempFile("", "literate_pandoc")
	if err != nil {
		return
	}
	defer tmp.Close()
	defer os.Remove(tmp.Name())

	_, err = tmp.Write([]byte(markdown))
	if err != nil {
		return
	}

	tmptpl, err := ioutil.TempFile("", "literate_pandoc")
	if err != nil {
		return
	}
	tempName := tmptpl.Name() + ".latex"
	tmptpl.Close()
	os.Remove(tmptpl.Name())
	defer os.Remove(tempName)

	err = ioutil.WriteFile(tempName, []byte(ltxTemplate), 0644)
	if err != nil {
		return
	}

	outFile := GetOutFile(filename + ".pdf")
	args = append(args, "-o")
	args = append(args, outFile)
	args = append(args, "--listing")
	args = append(args, "--template")
	args = append(args, tmptpl.Name())
	args = append(args, tmp.Name())
	pandoc := exec.Command("pandoc", args...)
	pandocOut, err := pandoc.CombinedOutput()
	if err != nil {
		fmt.Printf("[!] pandoc: '%s'\n", string(pandocOut))
	}
	return
}

// PdfWriter transforms the listing to a PDF. First, the markdown is
// written to a temporary file (which is removed when the function
// returns); this temporary file is then passed to pandoc for
// conversion. It is a wrapper around pdfWriter for non-unified
// output.
func PdfWriter(markdown, filename string) (err error) {
	return pdfWriter(markdown, filename, nil)
}

// PdfWriter transforms the listing to a PDF. First, the markdown is
// written to a temporary file (which is removed when the function
// returns); this temporary file is then passed to pandoc for conversion.
func UnifiedPdfWriter(markdown, filename string) (err error) {
	return pdfWriter(markdown, filename, []string{"-V", "documentclass=book", "--toc"})
}

// PandocTexWriter uses pandoc to convert the markdown output to a
// TeX file.
func PandocTexWriter(markdown, filename string) (err error) {
	tmp, err := ioutil.TempFile("", "literate_pandoc")
	if err != nil {
		return
	}
	defer tmp.Close()
	defer os.Remove(tmp.Name())

	_, err = tmp.Write([]byte(markdown))
	if err != nil {
		return
	}

	outFile := GetOutFile(filename + ".ltx")
	pandoc := exec.Command("pandoc", "-s", "-o", outFile, tmp.Name())
	err = pandoc.Run()
	return
}
