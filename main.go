package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ltadrian/go-gin-dynamodb-boilerplate/config"
	"github.com/ltadrian/go-gin-dynamodb-boilerplate/db"
	"github.com/ltadrian/go-gin-dynamodb-boilerplate/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	fmt.Printf("env: %s\n", *environment)
	config.Init(*environment)
	db.Init()

	r := server.NewRouter()
	r.Run(config.GetConfig().GetString("server.port"))
}
