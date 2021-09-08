package subprocess_test

import (
	. "github.com/mexisme/go-subprocess-mon/subprocess"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("subprocess.run", func() {
	var command *Subprocess

	BeforeEach(func() {
		command = New()
	})

	It("running 'true' succeeds", func() {
		Expect(command.SetCommand([]string{"true"}).Run()).NotTo(HaveOccurred())
	})

	It("running 'false' fails", func() {
		Expect(command.SetCommand([]string{"false"}).Run()).To(HaveOccurred())
	})
})
