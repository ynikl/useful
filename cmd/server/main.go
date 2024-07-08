package main

import (
	"fmt"
	"ft/internal"
)

func main() {
	if err := internal.Init(".", "config"); err != nil {
		panic(err)
	}
	fmt.Println("vim-go")
}
