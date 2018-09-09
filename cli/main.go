package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: cli <address>")
	}

	repl(os.Stdin, os.Args[1])
}

func repl(in io.Reader, addr string) {
	r := bufio.NewReader(in)

	for {
		fmt.Print("? ")
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		parts := strings.SplitN(line, " ", 2)

		if len(parts) == 1 {
			parts = append([]string{"json"}, parts[0])
		}

		if len(parts) != 2 {
			fmt.Println("input should be in the format '<action> <selector>' or just '<selector>'")
			continue
		}

		var (
			action   = strings.ToLower(parts[0])
			selector = parts[1]
			data     string
		)

		if action == "set" || action == "append" || action == "prepend" || action == "key" {
			for {
				fmt.Print("| ")
				line, err := r.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}

				data += line

				if json.Valid([]byte(data)) {
					break
				}
			}
		}

		if err := request(addr, action, selector, data); err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println()
	}
}

func request(addr, action, selector, data string) error {
	if !(action == "json" || action == "set" || action == "append" || action == "prepend" || action == "key" || action == "delete" || action == "empty") {
		return fmt.Errorf("invalid request action: %s", action)
	}

	u, err := url.Parse(addr)
	if err != nil {
		return fmt.Errorf("addr parser: %s", err.Error())
	}

	u.Path = "/" + action

	q := u.Query()
	q.Set("selector", url.PathEscape(selector))
	u.RawQuery = q.Encode()

	var resp *http.Response
	if action == "json" {
		resp, err = http.Get(u.String())
	} else {
		resp, err = http.Post(u.String(), "text/json", strings.NewReader(data))
	}
	if err != nil {
		return fmt.Errorf("http: %s", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		return fmt.Errorf("http err: %s", resp.Status)
	}

	if len(body) > 0 {
		var response interface{}
		if err := json.Unmarshal(body, &response); err != nil {
			return fmt.Errorf("json decode: %s\n%s", err.Error(), string(body))
		}

		out, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return fmt.Errorf("json marshal: %s", err.Error())
		}

		fmt.Println(string(out))
	}

	fmt.Println(http.StatusText(resp.StatusCode))

	return nil
}
