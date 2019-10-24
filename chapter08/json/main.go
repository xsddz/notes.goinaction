package main

import (
	"notes.goinaction/chapter08/json/decode"
	"notes.goinaction/chapter08/json/encode"
)

func main() {
	decode.RunAPIDemo()
	decode.RunStringDemo()

	encode.RunMapDemo()
}
