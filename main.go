package main

import (
	"fmt"

	"github.com/jakescript/rubix/cube"
)

func main() {
	c := cube.GenerateCube()

	for i := range c {
		c[i][1][1] = 8
		fmt.Println(c[i])
	}

}
