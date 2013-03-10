package main

// pandoc.go adds pandoc support for converting listings to PDF format.

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// PdfWriter transforms the listing to a PDF.
func PdfWriter(markdown, filename string) (err error) {
	tmp, err := ioutil.TempFile("", "golst_pandoc")
	if err != nil {
		return
	}
	defer tmp.Close()
	defer os.Remove(tmp.Name())
	_, err = tmp.Write([]byte(markdown))
	if err != nil {
		return
	}

        outName := "-o " + filename + ".pdf"
	pandoc := exec.Command("pandoc", outName, tmp.Name())
	err = pandoc.Run()
	return
}
