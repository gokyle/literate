package main

// texout handles TeX output.

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func SourceToLatex(filename string) (tex string, err error) {
	tex = `\documentclass[11pt]{article}
\usepackage{parskip}
\setlength{\parindent}{11pt}
\setlength{\parindent}{0cm}
\usepackage[margin=0.75in]{geometry}
\usepackage{fancyvrb}
\usepackage{iwona,palatino}
\usepackage[OT1]{fontenc}
\usepackage{textcomp}
\usepackage{lmodern}
\usepackage[hidelinks]{hyperref}
\usepackage{xltxtra}
\usepackage{graphicx}
\usepackage{algpseudocode}
\usepackage{amssymb}
\usepackage{listings}
\usepackage{framed}


\title{%s}
\author{literate listing}

\begin{document}
\maketitle

`
	tex = fmt.Sprintf(tex, filename)

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	var (
		line      string
		longLine  bool
		lineBytes []byte
		isPrefix  bool
		comment   = true
	)

	for {
		err = nil
		lineBytes, isPrefix, err = buf.ReadLine()
		if io.EOF == err {
			err = nil
			break
		} else if err != nil {
			break
		} else if isPrefix {
			line += string(lineBytes)

			longLine = true
			continue
		} else if longLine {
			line += string(lineBytes)
			longLine = false
		} else {
			line = string(lineBytes)
		}

		if CommentLine.MatchString(line) {
			if !comment {
				tex += "\\end{lstlisting}\n\n"
			}
			tex += CommentLine.ReplaceAllString(line, "")
			tex += "\n"
			comment = true
		} else {
			if comment {
				tex += "\n\n\\begin{lstlisting}[frame=single]\n"
				comment = false
			}
			tex += line + "\n"
		}
	}
	if !comment {
		tex += "\\end{lstlisting}\n\n"
	}
	tex += "\\end{document}\n"
	return
}

// TexWriter writes the transformed listing to a TeX file.
func TexWriter(listing string, filename string) (err error) {
	outFile := GetOutFile(filename + ".tex")
	err = ioutil.WriteFile(outFile, []byte(listing), 0644)
	return
}
