package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"

)

func main() {
	var lines []int
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()


	line := bufio.NewScanner(file)

	for line.Scan() {
		split := strings.Split(line.Text(), ":")
		trim := strings.Trim(split[1], " ")
		split = strings.Split(trim, " ")
		i, _ := strconv.Atoi(split[0])
		fmt.Printf("%d	%s\n", i, "kB")
		lines = append(lines, i)
		//fmt.Println(strings.Trim(sp[1], " "))
	}
	fmt.Println(lines[0])
}
