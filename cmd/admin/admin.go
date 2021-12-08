package main

import (
	"os"

	"github.com/chenzhiwei/kafka/cmd/admin/app"
)

func main() {
	if err := app.Execute(); err != nil {
		os.Exit(1)
	}
}
