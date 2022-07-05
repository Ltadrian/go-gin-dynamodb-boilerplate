package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ltadrian/test-dynamo-db-api/config"
	"github.com/ltadrian/test-dynamo-db-api/db"
	"github.com/ltadrian/test-dynamo-db-api/server"
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
