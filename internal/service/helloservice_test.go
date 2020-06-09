package service

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HelloService", func() {
	var helloService *HelloService

	BeforeEach(func() {
		helloService = &HelloService{}
	})

	Context("SayHello", func() {
		It("returns correct message", func() {
			helloMessage := helloService.SayHello("test")
			Expect(helloMessage).To(Equal("Hello, test!"))
		})
	})
})
