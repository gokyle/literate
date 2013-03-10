## golst

`golst` is a Go command line tool for generating readable program
listings.  It was inspired by Luke Gorries' post 
["Readable Programs"](http://blog.lukego.com/blog/2012/10/24/readable-programs/)
and, in particular, his `pbook` example.

Listings that are written to a file are written as <file.go>.<format>; for
example, producing listings for `listing.go`:

* markdown -> `listing.go.md`
* HTML -> `listing.go.html`
* PDF -> `listing.go.pdf`

### Dependencies

* `golst` uses [`blackfriday`](https://github.com/russross/blackfriday)
for producing HTML from the listing. This will automatically be pulled in
by `go get`.

* `pandoc` is required for conversion to PDF.

### License

`golst` is released under an ISC license.
