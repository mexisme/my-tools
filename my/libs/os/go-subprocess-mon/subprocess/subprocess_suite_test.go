package subprocess_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSubprocess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Subprocess Suite")
}
