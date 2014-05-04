## literate

`literate` is a Go command line tool for generating readable program
listings.  It was inspired by Luke Gorries' post 
["Readable Programs"](http://blog.lukego.com/blog/2012/10/24/readable-programs/)
and, in particular, his `pbook` example. It works by treating all text that have
a line comment as Markdown text, and putting code in code blocks.

Listings that are written to a file are written as &lt;file.go>.&lt;format>; for
example, producing listings for `listing.go`:

* markdown -> `listing.go.md`
* HTML -> `listing.go.html`
* PDF -> `listing.go.pdf`

Multiple languages are supported; for the list, pass the `-l help`
flag to `literate`. `literate` is a simple tool that only recognises line
comments right now; you can tell it what a line comment starts with and
it will use that.

### Examples

* To produce an HTML listing for `listing.go`:

        literate -o html listing.go

  This will produce `listing.go.html`, which is a standalone page. The
  template for this page is contained in the `html.go` source file, and
  is taken from [my site tyrfingr](http://tyrfingr.is).

* To produce a PDF listing for `pandoc.go`:

        literate -o pdf pandoc.go

  Similar to the previous example, this will produce the PDF file
  `pandoc.go.pdf`.

* To produce a markdown listing for `html.go` and have this printed
  to standard output:

        literate html.go

  Or, alternatively,

        literate -o - html.go

* You can produce listings for multiple files at the same time:

        literate -o pdf *.go

  If you ran this on the `literate` directory, you would have `listing.go.pdf`,
  `html.go.pdf`, `pandoc.go.pdf`, and `texout.go.pdf`.

### Supported Output Formats

Standalone listings are generated; the following formats (selected with
the `-o` flag) are supported:

* html - generate HTML listing
* latex - uses [pandoc](http://www.johnmacfarlane.net/pandoc/) to convert the
  markdown to a LaTeX listing.
* md - generate markdown listing
* pdf - uses [pandoc](http://www.johnmacfarlane.net/pandoc/) to convert the
  markdown to a PDF listing.
* tex - rudimentary pure-Go TeX listing.

### Dependencies

* `literate` uses [`blackfriday`](https://github.com/russross/blackfriday)
for producing HTML from the listing. This will automatically be pulled in
by `go get`.

* `pandoc` is required for conversion to PDF.

### Why?

Ultimately, I wanted a way to take notes while reading a book and have
a way to extract better-looking notes from a source file while still
retaining the ability to run the files. This program does (pretty much)
just what I want it to, although block comments would be nice. Originally,
I had [golst](https://github.com/gokyle/golst), but wanted to be able
to use it with Lisp and Haskell.

### License

`literate` is released under an ISC license. For details, see
[LICENSE](./LICENSE) in the source repository.
