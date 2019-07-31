package info

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"

)

type Info struct {}

func (i *Info) MemTotal() int {
	return i.info(0)
}
func (i *Info) MemFree() int {
	return i.info(1)
}

func (i *Info) info(num int) int {
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
		//fmt.Printf("%d	%s\n", i, "kB")
		lines = append(lines, i)
	}
	return lines[num]
}
