package main

import (
	"fmt"
	"time"
	"os/exec"
	"os"
	"bufio"

	"./info"
)

type datastr struct {
	Uptime		string
	Memtotal	string
	MemFree		string
	MemAvail	string
	Use			string
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func title(hostname string) {
	i := 0
	var line string
	for i < len(hostname) {
		line += "-"
		i += 1
	}
	fmt.Printf("\033[37m+%s+\n",line)
	fmt.Printf("\033[37m|%s|\n", hostname)
	fmt.Printf("\033[37m+%s+\n",line)
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

func cont() {
	var lines []string
	var line *bufio.Scanner
	for {
		file, err := os.Open("/var/snap/lxd/common/lxd/containers/definite-muskox/rootfs/tmp/info")
		if err != nil {
			return
			fmt.Println(err)
		}

		line = bufio.NewScanner(file)

		for line.Scan() {
			lines = append(lines, line.Text())
		}
		if len(lines) == 0 {
			continue
		}
		file.Close()
		//fmt.Println(len(lines))
		break
	}
	title(lines[0])
	fmt.Printf("\033[32m[Uptime]\033[39m%s",lines[1])
	fmt.Printf("\033[1E")
	fmt.Printf("\033[32m[MemTotal]\033[39m:	%s mB",lines[2])
	fmt.Printf("\033[1E")
	fmt.Printf("\033[32m[MemFree]\033[39m:	%s mB",lines[3])
	fmt.Printf("\033[1E")
	fmt.Printf("\033[32m[MemAvailable]\033[39m:	%s mB",lines[4])
	fmt.Printf("\033[1E")
	fmt.Printf("\033[32m[use]\033[39m:		%s mB",lines[5])
	fmt.Printf("\033[1E")
	fmt.Printf("\033[36m%s\n",lines[6])
}

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\033[37m+----------------------+\n")
	i := &info.Info{}
	clear()
	var total, free, avail int

	for {
		total = i.MemTotal()
		free =	i.MemFree()
		avail = i.MemAvailable()
		title(name)
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
		fmt.Printf("\033[1E")
		cont()
		time.Sleep(1000*time.Millisecond)
		clear()
	}
}
