package parser_test

import (
	"bytes"

	"github.com/cirocosta/execsnoopvis/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scanner", func() {

	var (
		content string
		done    bool
		err     error
		// node    parser.Node
		scanner parser.Scanner
	)

	JustBeforeEach(func() {
		scanner = parser.NewScanner(bytes.NewReader([]byte(content)))
		_, done, err = scanner.Scan()
	})

	Context("on empty reader", func() {

		BeforeEach(func() {
			content = ""
		})

		It("succeeds", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("is done", func() {
			Expect(done).To(BeTrue())
		})

	})

	Context("having a line to parse", func() {

		Context("not having enough fields", func() {

			BeforeEach(func() {
				content = "A"
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})

		})

		Context("being the line a header", func() {

			It("returns an error that indicates that it's a header", func() {

			})

		})

	})

})
