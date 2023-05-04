package pwd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPwdHash(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PwdHash Suite")
}
