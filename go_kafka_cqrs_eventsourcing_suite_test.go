package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoKafkaCqrsEventsourcing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoKafkaCqrsEventsourcing Suite")
}
