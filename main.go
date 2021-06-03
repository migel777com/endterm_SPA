package main

import (
	"bytes"
	"fmt"
)

func main() {
	fastOut := new(bytes.Buffer)
	Fast(fastOut)
	fastResult := fastOut.String()

	fmt.Println(fastResult)

	/*var b byte
	b = 'B'


	if 'A' <= b && b <= 'Z' {
		b = b + 32
		fmt.Println(string(b))
	}*/

}
