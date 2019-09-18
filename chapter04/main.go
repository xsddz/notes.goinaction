package main

import (
	"fmt"

	"notes.goinaction/chapter04/array"
	"notes.goinaction/chapter04/mmap"
	"notes.goinaction/chapter04/slice"
)

func main() {
	fmt.Printf("[[array test]]:\n")
	array.Run()

	fmt.Printf("\n[[slice test]]:\n")
	slice.Run()

	fmt.Printf("\n[[map test]]:\n")
	mmap.Run()
}
