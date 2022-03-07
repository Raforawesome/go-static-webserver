package main

import (
	"errors"
	"fmt"
	"git.raforaweso.me/Raforawesome/go-lists"
	"os"
	"strconv"
	"strings"
)

func ReadConfig() map[uint64]lists.StrList {
	m := make(map[uint64]lists.StrList)

	dirName, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat("gss-config"); err == nil {
		fmt.Println("Found config file")
		fmt.Println("Loading config...")
		fmt.Println("")
		fileRaw, _ := os.ReadFile("gss-config")
		lines := strings.Split(string(fileRaw), "\n")

		for _, line := range lines {
			fields := strings.Split(line, " ")
			target := fields[0]
			target = strings.ReplaceAll(target, "[", "")
			target = strings.ReplaceAll(target, "]", "")
			portStr := fields[1]
			port, err := strconv.ParseUint(portStr, 10, 16)
			if err != nil {
				panic(err)
			}

			ls := lists.NewStrList()
			ls.Append(Ternary(strings.HasPrefix(target, "./") || strings.HasPrefix(target, "/"), "path", "url"))
			ls.Append(target)
			m[port] = ls
		}

	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No config file exists")
		fmt.Println(fmt.Sprintf("Writing default config to %v/gss-config", dirName))
		fmt.Println("")
		err := os.WriteFile("gss-config", []byte("[./views] 443"), 0644)
		if err != nil {
			panic(err)
		}
		//m[443] = [2]string{"path", "./views"}
		ls := lists.NewStrList()
		ls.Append("path")
		ls.Append("./views")
		m[443] = ls
	}

	return m
}
