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

func title() {
	fmt.Printf("\033[37m+----------------------+\n")
	fmt.Printf("\033[37m|       HostInfo       |\n")
	fmt.Printf("\033[37m+----------------------+\n")
}

func graph(total,free int) string {
	i := 0
	grp := "["
	for i < (total-free)/(total/20) {
		grp = grp + "|"
		i += 1
	}
	for i < 20 {
		grp = grp + "-"
		i += 1
	}
	grp += "]"
	return grp
}

func main() {
	i := &info.Info{}
	clear()
	var total, free, avail int

	for {
		total = i.MemTotal()
		free =	i.MemFree()
		avail = i.MemAvailable()
		title()
		fmt.Printf("\033[32m[Uptime]\033[39m%s",i.Uptime())
		fmt.Printf("\033[1E")
		fmt.Printf("\033[32m[MemTotal]\033[39m:	%d mB",total/1024)
		fmt.Printf("\033[1E")
		fmt.Printf("\033[32m[MemFree]\033[39m:	%d mB",free/1024)
		fmt.Printf("\033[1E")
		fmt.Printf("\033[32m[MemAvailable]\033[39m:	%d mB",avail/1024)
		fmt.Printf("\033[1E")
		fmt.Printf("\033[32m[use]\033[39m:		%d mB",(total-free)/1024)
		fmt.Printf("\033[1E")
		fmt.Printf("\033[36m%s\n",graph(total/1024,free/1024))
		time.Sleep(1000*time.Millisecond)
		clear()
	}
}
