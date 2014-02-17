package github_markdown

import (
	"strings"
	"io/ioutil"
	"net/http"
)

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
