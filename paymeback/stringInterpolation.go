// https://stackoverflow.com/questions/50095616/go-string-interpolation
// https://gobyexample.com/string-formatting

package main

import (
	"fmt"
)

func main() {
	source := "./source.json"
	fmt.Printf("Successfully opened \"%v\"\n", source)
}
