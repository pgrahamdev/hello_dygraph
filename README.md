# hello_dygraph

Simple illustration of how to use Go templates with the Dygraphs
JavaScript library

Note that if `dygraph-cdn.html` is used as the template instead of
`dygraph.html` in the Go source file, then `http.Handle()` call can be
commented out since the program doesn't need to serve up the
JavaScript and CSS assets.
