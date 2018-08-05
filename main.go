package main

import (
	"fmt"
	"os"

	"github.com/Zac-Garby/db/server"
)

func main() {
	s := &server.Server{
		Addr: ":7913",
	}

	fmt.Println("listening on :7913")
	if err := s.Listen(); err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
