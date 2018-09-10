package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/NewbieCoder99/go-night-framework/config"
	"github.com/NewbieCoder99/go-night-framework/database"
	"github.com/NewbieCoder99/go-night-framework/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	Database.Init()
	server.Init()
}
