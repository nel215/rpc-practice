package client

import (
	"github.com/nel215/rpc-practice/service"
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

func TestMethod(t *testing.T) {
	s := new(service.Service)
	err := rpc.Register(s)
	if err != nil {
		t.Error(err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		t.Error(err)
	}
	go func() {
		err := http.Serve(l, nil)
		if err != nil {
			t.Error(err)
		}
	}()

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		t.Error(err)
	}

	args := &service.Args{
		A: 32,
		B: "test",
	}
	var res string
	err = client.Call("Service.Method", args, &res)
	if err != nil {
		t.Error(err)
	}
	if res != "test: 32" {
		t.Errorf("expected test: 32, but got %s", res)
	}
}
