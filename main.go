package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	cfg := ReadConfig()
	fmt.Println(cfg)

	for port, values := range cfg {
		var address string
		address = "localhost:" + strconv.FormatUint(port, 10)

		if values[0] == "path" {
			// handle file servers
			if _, err := os.Stat(values[1]); err == nil {
				fmt.Println("Found directory", values[1])
			} else if errors.Is(err, os.ErrNotExist) {
				fmt.Println("Directory", values[1], "not found.")
				fmt.Println("Creating", values[1]+"...")
				err := os.Mkdir(values[1], 0644)
				if err != nil {
					panic(err)
				}
			}

			fmt.Println("Starting a file server on", address, "using directory", values[1])
			log.Fatalln(address, http.FileServer(http.Dir(values[1])))
		} else {
			// handle redirect (reverse proxy in future)
		}
	}
}
