package github_markdown_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoGithubMarkdown(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GithubMarkdown Suite")
}
