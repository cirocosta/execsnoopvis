package parser_test

import (
	"bytes"

	"github.com/cirocosta/execsnoopvis/parser"
	"github.com/pkg/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scanner", func() {

	var (
		content string
		done    bool
		err     error
		node    parser.Node
		scanner parser.Scanner
	)

	JustBeforeEach(func() {
		scanner = parser.NewScanner(bytes.NewReader([]byte(content)))
		node, done, err = scanner.Scan()
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

			BeforeEach(func() {
				content = "PCOMM            PID    PPID   RET ARGS"
			})

			It("returns an error that indicates that it's a header", func() {
				Expect(errors.Cause(err)).To(Equal(parser.ErrIsHeader))
			})

		})

		Context("not being a header", func() {

			BeforeEach(func() {
				content = "go               17123  16232    0 /usr/local/go/bin/go build -v ."
			})

			It("succeeds", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("retrieves the fields accordingly", func() {
				Expect(node.Command).To(Equal("go"))
				Expect(node.Pid).To(BeNumerically("==", 17123))
				Expect(node.Ppid).To(BeNumerically("==", 16232))
				Expect(node.ExitCode).To(Equal(0))
				Expect(node.Argv).To(ConsistOf([]string{
					"/usr/local/go/bin/go",
					"build",
					"-v",
					".",
				}))
			})

		})

	})

})
