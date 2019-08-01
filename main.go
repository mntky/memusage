package main

import (
	"fmt"
	"time"
	"os/exec"
	"os"

	"./info"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	i := &info.Info{}
	clear()

	for {
		fmt.Printf("MemTotal:%d",i.MemTotal())
		fmt.Printf("\033[1E")
		fmt.Printf("MemFree:%d",i.MemFree())
		time.Sleep(1000*time.Millisecond)
		clear()
	}
}
