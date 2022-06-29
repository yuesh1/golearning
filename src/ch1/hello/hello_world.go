package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 0 {
		fmt.Println("hello my args " + os.Args[1])
	}
}
