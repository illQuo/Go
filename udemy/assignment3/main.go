package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		f, err := os.Open(args[1])
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, f)
	} else {
		fmt.Println("too many arguments")
		os.Exit(1)
	}
}
