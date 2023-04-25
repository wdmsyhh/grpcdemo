package main

import (
	"context"
	hello "grpcdemo/proto"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.StringMessage) (*hello.StringMessage, error) {
	reply := &hello.StringMessage{Value: "hello:" + args.GetValue()}
	return reply, nil
}
