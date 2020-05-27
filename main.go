package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

type cfg struct {
	host string
	port string
}

func parseConfig() cfg {
	host := flag.String("host", "", "host")
	port := flag.String("port", "8080", "port")
	flag.Parse()

	return cfg{host: *host, port: *port}
}

func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Printf("unable to echo: %v", err)
	}
}

func startServer(config cfg) error {
	lis, err := net.Listen("tcp", net.JoinHostPort(config.host, config.port))
	if err != nil {
		return fmt.Errorf("echo server can't listen: %v", err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("can't accept connection: %v", err)
		}
		go echo(conn)
	}
}

func main() {
	cfg := parseConfig()
	log.Printf("starting echo server: %+v", cfg)
	err := startServer(cfg)
	if err != nil {
		log.Panicf("can't run echo server: %v", err)
	}
}
