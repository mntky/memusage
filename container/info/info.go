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
func (i *Info) MemAvailable() int {
	return i.info(2)
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
		lines = append(lines, i)
	}
	return lines[num]
}

func (i *Info) Uptime() string {
	var secondstr string
	file, err := os.Open("/proc/uptime")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	line := bufio.NewScanner(file)
	for line.Scan() {
		secondstr = line.Text()
	}
	split := strings.Split(secondstr, " ")
	second, _ := strconv.ParseFloat(split[0], 32)

	upseconds := int(second)%60
	upminutes := int(second)/60
	uphours := upminutes/60
	updays := uphours/24
	uphours = uphours%24
	upminutes = upminutes % 60
	ret := fmt.Sprintf("%ddays %d:%d:%d",updays, uphours, upminutes, upseconds)
	return ret
}

