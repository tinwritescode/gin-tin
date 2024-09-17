package main

import (
	"log"

	"github.com/tinwritescode/gin-tin/pkg/api"
	"github.com/tinwritescode/gin-tin/pkg/config"
)

func main() {
	a, err := api.NewApp(
		&config.Config{
			ServerAddress: ":8080",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
