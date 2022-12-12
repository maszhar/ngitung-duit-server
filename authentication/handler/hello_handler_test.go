package handler

import (
	"context"
	"testing"

	pb "github.com/djeniusinvfest/inventora/auth/proto"
)

func TestHello(t *testing.T) {
	s := NewHandler(nil)
	wantedResp := "authentication server is running"

	req := &pb.HelloRequest{}
	resp, err := s.Hello(context.Background(), req)
	if err != nil {
		t.Errorf("HelloTest() got unexpected error")
	}
	if resp.Message != wantedResp {
		t.Errorf("Hello()='%v', wanted '%v'", resp.Message, wantedResp)
	}
}
