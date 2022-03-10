package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cfg := ReadConfig()
	fmt.Println(cfg)

	for port, values := range cfg {
		var address string
		address = "localhost:" + strconv.FormatUint(port, 10)

		if values.Components[0] == "path" {
			//fmt.Println("Starting server on " + address)
			StartServer(address, values.Components[1])
		}
	}
}

func StartServer(addr string, path string) {
	fmt.Println("Starting server on " + addr)
	fmt.Println("Serving " + path)
	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Fatal(http.ListenAndServe(addr, nil))
}

/*
func main() {
	cfg := ReadConfig()
	fmt.Println(cfg)

	for port, values := range cfg {
		var address string
		address = "localhost:" + strconv.FormatUint(port, 10)

		if values.Components[0] == "path" {
			// handle file servers
			if _, err := os.Stat(values.Components[1]); err == nil {
				fmt.Println("Found directory", values.Components[1])
			} else if errors.Is(err, os.ErrNotExist) {
				fmt.Println("Directory", values.Components[1], "not found.")
				fmt.Println("Creating", values.Components[1]+"...")
				err := os.Mkdir(values.Components[1], 0644)
				if err != nil {
					panic(err)
				}
			}

			fmt.Println("Starting a file server on", address, "using directory", values.Components[1])
			log.Fatalln(address, http.FileServer(http.Dir(values.Components[1])))
		} else {
			// handle redirect (reverse proxy in future)
		}
	}
}
*/
