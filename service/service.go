package service

import (
	"fmt"
)

type Service struct{}

type Args struct {
	A int64
	B string
}

func (s *Service) Method(args *Args, reply *string) error {
	*reply = fmt.Sprintf("%s: %d", args.B, args.A)
	return nil
}
