package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ruifeng377/pRinter/api"
	"github.com/ruifeng377/pRinter/db"
	"github.com/ruifeng377/pRinter/printer"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	filesDir := flag.String("dir", "", "the dir to store files")
	flag.Parse()

	fileStore := db.NewFileStore(*filesDir)
	printer := printer.NewPrinterForWin()
	server, err := api.NewServer(fileStore, printer)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(fmt.Sprintf("%s:%d", "0.0.0.0", *port))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
