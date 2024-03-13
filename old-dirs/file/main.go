package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")

	file, err := os.Create("sss.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	output := "message"
	file.Write(([]byte)(output))
}
