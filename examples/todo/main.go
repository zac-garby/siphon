package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type todo struct {
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("Welcome to the example TODO app")
	fmt.Println("Enter either 'ls', 'add', 'done', 'rm', 'clear', or 'quit'")

	r := bufio.NewReader(os.Stdin)

outer:
	for {
		line := input("todo> ", r)

		switch line {
		case "ls":
			resp, err := request("GET", "json?selector=todos", "")
			if err != nil {
				fmt.Println(err)
				continue outer
			}

			todos := make([]*todo, 4)

			if err := json.Unmarshal(resp, &todos); err != nil {
				fmt.Println(err)
				continue outer
			}

			for i, todo := range todos {
				done := "     "
				if todo.Completed {
					done = "DONE "
				}

				fmt.Printf("%s%d) %s\n", done, i, todo.Description)
			}

		case "add":
			description := input("description> ", r)

			todo := &todo{
				Completed:   false,
				Description: description,
			}

			encoded, err := json.Marshal(todo)
			if err != nil {
				fmt.Println(err)
				continue outer
			}

			if _, err = request("POST", "append?selector=todos", string(encoded)); err != nil {
				fmt.Println(err)
				continue outer
			}

		case "done":
			index := input("todo index> ", r)

			if _, err := strconv.ParseInt(index, 10, 64); err != nil {
				fmt.Println("invalid index -- not an integer")
				continue outer
			}

			resp, err := request("GET", fmt.Sprintf("json?selector=todos[%s].completed", index), "")
			if err != nil {
				fmt.Println(err)
				continue outer
			}

			wasCompleted := false

			if err := json.Unmarshal(resp, &wasCompleted); err != nil {
				fmt.Println(err)
				continue outer
			}

			if _, err = request("POST", fmt.Sprintf("set?selector=todos[%s].completed", index), fmt.Sprint(!wasCompleted)); err != nil {
				fmt.Println(err)
				continue outer
			}

		case "rm":
			index := input("todo index> ", r)

			if _, err := strconv.ParseInt(index, 10, 64); err != nil {
				fmt.Println("invalid index -- not an integer")
				continue outer
			}

			if _, err := request("POST", "unset?selector=todos", index); err != nil {
				fmt.Println(err)
				continue outer
			}

		case "clear":
			if _, err := request("POST", "empty?selector=todos", ""); err != nil {
				fmt.Println(err)
				continue outer
			}

		case "quit":
			break outer

		default:
			fmt.Println("Valid commands are 'ls', 'add', 'done', 'rm', 'clear', or 'quit'")
		}
	}
}

func request(method, path, body string) ([]byte, error) {
	var resp *http.Response
	if method == "POST" {
		response, err := http.Post("http://localhost:7913/"+path, "text/json", strings.NewReader(body))
		if err != nil {
			return nil, err
		}

		resp = response
	} else if method == "GET" {
		response, err := http.Get("http://localhost:7913/" + path)
		if err != nil {
			return nil, err
		}

		resp = response
	} else {
		return nil, errors.New("only POST and GET requests are supported")
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errorStruct := new(struct {
			Error string `json:"err"`
		})

		if err := json.Unmarshal(r, &errorStruct); err != nil {
			return nil, errors.New("unknown error, response not JSON parseable." + resp.Status)
		}

		return nil, errors.New(errorStruct.Error)
	}

	return r, nil
}

func input(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)

	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return strings.TrimSpace(line)
}
