package main

import (
	"time"
	"log"
	"fmt"
	"os"
	"./info"
)

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
	var total, free, avail, use  int
	var upt, grph string
	hname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	for {
		f, err :=os.Create("/tmp/info")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		total = i.MemTotal()/1024
		free =	i.MemFree()/1024
		avail = i.MemAvailable()/1024
		use = (total-free)
		upt = i.Uptime()
		grph = graph(total,free)
		sdata := fmt.Sprintf("%s\n%s\n%d\n%d\n%d\n%d\n%s\n", hname,upt,total,free,avail,use,grph)
		//fmt.Println(sdata)
		f.Write([]byte(sdata))
		time.Sleep(1000*time.Millisecond)
	}
}
