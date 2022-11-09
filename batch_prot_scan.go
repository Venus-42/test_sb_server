package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

func processPortItem(port string) []string {
	fmt.Println("processPortItem", port)
	var ports []string
	arr := strings.Split(port, ",")
	for _, p := range arr {
		if strings.Contains(p, "-") {
			ports = append(ports, rangeToArr(p)...)
		} else {
			ports = append(ports, p)
		}
	}
	return ports
}

// convert "1-3" to ["1", "2", "3"]
func rangeToArr(s string) []string {
	if strings.Contains(s, "-") {
		var arr []string
		from, _ := strconv.Atoi(strings.Split(s, "-")[0])
		to, _ := strconv.Atoi(strings.Split(s, "-")[1])
		if from == 0 {
			from = 1
		}
		if to == 0 {
			to = 65535
		}
		for i := 0; i < to; i++ {
			arr = append(arr, strconv.Itoa(i))
		}
		return arr
	} else {
		return []string{s}
	}
}

func scan(ip, port string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", ip+":"+port, time.Second*3)
	if err != nil {
		//log.Println(port, "is not open")
		return
	}
	defer conn.Close()
	log.Printf("%s:%s is open", ip, port)
}
func main() {
	ip := flag.String("h", "192.168.5.110", "")
	port := flag.String("p", "0-65535", "")
	flag.Parse()
	log.Println("scan target:", *ip, *port)
	// thread synchronization
	wg := &sync.WaitGroup{}
	for _, p := range processPortItem(*port) {
		wg.Add(1)
		//log.Println("start scan ", p)
		go scan(*ip, p, wg)
	}
	wg.Wait()

}
