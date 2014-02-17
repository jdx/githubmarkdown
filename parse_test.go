package githubmarkdown_test

import (
	. "github.com/dickeyxxx/githubmarkdown"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {
	It("renders markdown", func() {
		html, _ := Parse("Hello world github/linguist#1 **cool**, and #1!")
		Ω(html).Should(Equal("<p>Hello world github/linguist#1 <strong>cool</strong>, and #1!</p>"))
	})
})
