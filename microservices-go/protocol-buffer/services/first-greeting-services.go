package main

import (
	"context"

	hello "github.com/git-emran/microservices-go/protocol-buffer/proto"
)

type Say struct{}

func Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}
