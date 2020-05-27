package main

import (
	"net"
	"testing"
	"time"
)

func TestEcho(t *testing.T) {
	go startServer(cfg{port: "8080"})
	var conn net.Conn
	var err error
	for i := 0; i < 3; i++ {
		conn, err = net.Dial("tcp", ":8080")
		if err != nil && i == 2 {
			t.Fatal(err)
		} else {
			time.Sleep(1 * time.Second)
		}
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
