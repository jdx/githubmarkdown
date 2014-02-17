GitHub Markdown in Go
=====================

[![wercker status](https://app.wercker.com/status/b19fd4a4c797d9b5b10a75f85f1d158e/m/master "wercker status")](https://app.wercker.com/project/bykey/b19fd4a4c797d9b5b10a75f85f1d158e)
[![GoDoc](https://godoc.org/github.com/dickeyxxx/githubmarkdown?status.png)](https://godoc.org/github.com/dickeyxxx/githubmarkdown)

Parses markdown via the Github API.

Usage
-----

```go
html, err := githubmarkdown.Parse("Hello world github/linguist#1 **cool**, and #1!")
```
