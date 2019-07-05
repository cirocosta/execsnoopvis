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
		err     error
		node    parser.Node
		scanner parser.Scanner
	)

	Describe("ScanAll", func() {

		var (
			nodes []*parser.Node
		)

		JustBeforeEach(func() {
			scanner = parser.NewScanner(bytes.NewReader([]byte(content)))
			nodes, err = scanner.ScanAll()
		})

		Context("empty reader", func() {

			BeforeEach(func() {
				content = ""
			})

			It("succeeds", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns no nodes", func() {
				Expect(nodes).To(BeEmpty())
			})
		})

		Context("entire file", func() {

			BeforeEach(func() {
				content = `PCOMM            PID    PPID   RET ARGS
go               17123  16232    0 /usr/local/go/bin/go build -v .
go               17124  16232    0 /usr/local/go/bin/go build -v .`
			})

			It("succeeds", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns node nodes", func() {
				Expect(nodes).To(HaveLen(2))
			})

			It("has fields properly set", func() {
				Expect(nodes[0].Pid).To(Equal(uint64(17123)))
				Expect(nodes[1].Pid).To(Equal(uint64(17124)))
			})
		})

	})

	Describe("Scan", func() {

		var (
			done bool
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

				Context("not having argv", func() {

					BeforeEach(func() {
						content = "go               17123  16232    0"
					})

					It("succeeds", func() {
						Expect(err).ToNot(HaveOccurred())
					})

					It("retrieves the fields accordingly", func() {
						Expect(node.Command).To(Equal("go"))
						Expect(node.Pid).To(BeNumerically("==", 17123))
						Expect(node.Ppid).To(BeNumerically("==", 16232))
						Expect(node.ExitCode).To(Equal(0))
					})

				})

				Context("having everything", func() {

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

		Context("having header + regular line", func() {

			BeforeEach(func() {
				content = `PCOMM            PID    PPID   RET ARGS
go               17123  16232    0 /usr/local/go/bin/go build -v .`
			})

			It("detects the header", func() {
				Expect(errors.Cause(err)).To(Equal(parser.ErrIsHeader))
			})

			It("is able to scan the next line", func() {
				node, done, err = scanner.Scan()
				Expect(err).NotTo(HaveOccurred())
				Expect(done).NotTo(BeTrue())
				Expect(node.Command).To(Equal("go"))
			})

		})

	})

})
