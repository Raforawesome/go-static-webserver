package main

import (
	"errors"
	"fmt"
	"os"
)

func ReadConfig() map[string]uint16 {
	m := make(map[string]uint16)

	dirName, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat("gss-config"); err == nil {
		fmt.Println("Found config file")
		fmt.Println("Loading config...")
		file, _ := os.ReadFile("gss-config")
		fmt.Println(string(file))

	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No config file exists")
		fmt.Println(fmt.Sprintf("Writing default config to %v/gss-config", dirName))
		err := os.WriteFile("gss-config", []byte("[./views]=443"), 0644)
		if err != nil {
			panic(err)
		}
	}

	return m
}
