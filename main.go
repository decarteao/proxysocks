package main

import (
	"fmt"
	"log"
	"os"

	"github.com/things-go/go-socks5"
)

const (
	listen_ip   = "127.0.0.1"
	listen_port = 1234
)

func main() {
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)

	log.Printf("[!] SOCKS Server iniciado na porta %d\n", listen_port)
	if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%d", listen_ip, listen_port)); err != nil {
		panic(err)
	}
}
