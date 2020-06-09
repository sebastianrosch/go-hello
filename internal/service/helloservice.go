package service

import (
	"fmt"
)

type HelloService struct {
}

func (h *HelloService) SayHello(word string) string {
	return fmt.Sprintf("Hello, %s!", word)
}