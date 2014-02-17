// Package githubmarkdown provides a function to parse markdown via GitHub's API
// http://developer.github.com/v3/markdown/
package githubmarkdown

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Parses markdown with an HTTP request to GitHub's api.
// It retuns an html string.
func Parse(md string) (string, error) {
	reader := strings.NewReader(md)
	resp, err := http.Post("https://api.github.com/markdown/raw", "text/x-markdown", reader)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(html), nil
}
