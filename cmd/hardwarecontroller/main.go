// cmd/hardwarecontroller/main.go
package main

import (
	"flag"
	"log"
	"os"

	"hardwarecontroller/internal/hardwarecontroller"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := hardwarecontroller.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
