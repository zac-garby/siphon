package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Zac-Garby/siphon/server"
)

var schemaFile = flag.String("schema", "schema.sip", "the location of the file containing the database schema")
var port = flag.Int("port", 7913, "the port on which to listen")

func main() {
	flag.Parse()

	if *port >= 65536 {
		log.Fatal("port cannot be larger than 65,536")
	}

	schemaBytes, err := ioutil.ReadFile(*schemaFile)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewServer(fmt.Sprintf(":%d", *port), string(schemaBytes))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("listening on :%d...\n", *port)
	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
