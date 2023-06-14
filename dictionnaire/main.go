package main

import (
	"fmt"
	"dict/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")

	handleErr(err)

	defer d.Close()
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
	}
}