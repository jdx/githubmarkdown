GitHub Markdown in Go
=====================

[![wercker status](https://app.wercker.com/status/fde87e53c1f60cb4f5b24f16fc091e6d/m/master "wercker status")](https://app.wercker.com/project/bykey/fde87e53c1f60cb4f5b24f16fc091e6d)

[![GoDoc](https://godoc.org/github.com/dickeyxxx/go-github-markdown?status.png)](https://godoc.org/github.com/dickeyxxx/go-github-markdown)

Parses markdown via the Github API.

Usage
-----

```go
html, err := githubmarkdown.Parse("Hello world github/linguist#1 **cool**, and #1!")
```
