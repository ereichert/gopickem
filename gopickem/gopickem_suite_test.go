package gopickem_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoPickem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gopickem spec suite")
}
