package main

import (
	"net"
	"testing"
)

func TestEcho(t *testing.T) {
	go startServer(cfg{port: "8080"})
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	req := "hello world"
	_, err = conn.Write([]byte(req))
	if err != nil {
		t.Error(err)
	}
	answer := make([]byte, 128)
	n, err := conn.Read(answer)
	if err != nil {
		t.Error(err)
	}
	resp := string(answer[:n])
	if resp != req {
		t.Errorf("expected 'hello world', got: %s", req)
	}
}
