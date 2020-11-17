package main

import (
	"fmt"
	"os"
)

func main() {

	argc := len(os.Args)
	fmt.Printf("argc=%d\n", argc)
	for i := 0; i < argc; i++ {
		fmt.Printf("argv[%d]=%s\n", i, os.Args[i])
	}
	os.Exit(0)
}
