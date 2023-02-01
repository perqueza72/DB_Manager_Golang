package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Add folder path.")
		return
	}

	fmt.Print(os.Args[1])
}
