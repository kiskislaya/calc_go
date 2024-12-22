package main

import (
	"log"

	"github.com/kiskislaya/calc_go/internal/application"
)

func main() {
	app := application.New()
	log.Fatal(app.RunServer())
}
