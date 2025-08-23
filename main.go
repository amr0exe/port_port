package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func scanAport(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		fmt.Printf("    closed :%d\n", port)
		return
	}

	fmt.Printf(" port is open :%d\n", port)
	defer conn.Close()
}

func loadPortsFromFile(filename string) []int {
	var ports []int

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		p, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Error converting text-to-integer: %v", err)
			continue
		}

		ports = append(ports, p)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from the file: %v", err)
	}

	return ports
}

func main() {
	target := "scanme.nmap.org"

	ports := loadPortsFromFile("top100.txt")

	var wg sync.WaitGroup

	for _, p := range ports {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()
			scanAport(target, p)
		}(p)
	}

	wg.Wait()
}
