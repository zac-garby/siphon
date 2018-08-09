package main

import (
	"fmt"
	"os"

	"github.com/Zac-Garby/db/server"
)

func main() {
	s, err := server.NewServer(":7913", `
	me: user
	
	struct user {
		name: string
		email: string
		age: uint8
		friends: [user]
	}`)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error making database:", err)
		os.Exit(1)
	}

	fmt.Println("listening on :7913")
	if err := s.Listen(); err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
