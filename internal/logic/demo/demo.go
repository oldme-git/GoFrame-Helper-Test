package demo

import (
	"GoFrame-Helper-Test/internal/service"
	"context"
)

type sDemo struct {
}

func init() {
	service.RegisterDemo(New())
}

func New() *sDemo {
	return &sDemo{}
}

func (s *sDemo) DemoA(ctx context.Context) (err error) {
	return nil
}
