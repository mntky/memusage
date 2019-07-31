package main

import (
	"fmt"

	"./info"
)

func main() {
	i := &info.Info{}

	fmt.Println(i.MemTotal())
}
