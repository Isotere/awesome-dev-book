package main

import "fmt"

type Service struct {
	config Config
}

func NewService(opts ...ConfigOption) *Service {
	return &Service{config: NewConfig(opts...)}
}

func main() {
	service := NewService(WithDayHours(12, 24))

	fmt.Printf("Service: %+v", *service)
}
