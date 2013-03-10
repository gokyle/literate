package main

// pandoc.go adds pandoc support for converting listings to PDF format.

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// PdfWriter transforms the listing to a PDF. First, the markdown is
// written to a temporary file (which is removed when the function
// returns); this temporary file is then passed to pandoc for conversion.
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

	outName := filename + ".pdf"
	pandoc := exec.Command("pandoc", "-o", outName, tmp.Name())
	err = pandoc.Run()
	return
}
