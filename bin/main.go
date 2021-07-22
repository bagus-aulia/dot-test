package main

import (
	"fmt"

	"github.com/bagus-aulia/dot-test/app/server"
	"github.com/bagus-aulia/dot-test/config"
)

func main() {
	fmt.Println(
		`=========================
     Test DOT server
=========================`)
	server.Start()
	fmt.Println("get config", config.Get("db.host"))
}
