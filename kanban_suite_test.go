package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKanban(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kanban Suite")
}
