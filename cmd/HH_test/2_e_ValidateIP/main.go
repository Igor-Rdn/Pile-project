package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidIP(ip string) bool {

	count := 0

	for num := range strings.SplitSeq(ip, ".") {

		count++

		if i, err := strconv.Atoi(num); err != nil || i < 0 || i > 255 {
			return false
		}

	}

	return count == 4

}

func processIPs(ipInput string, threshold int) {

	ipMap := make(map[string]int)
	ipSlice := make([]string, 0) //Для вывода в порядке появления

	for ip := range strings.SplitSeq(ipInput, " ") {

		if isValidIP(ip) {
			if _, exists := ipMap[ip]; !exists {
				ipSlice = append(ipSlice, ip)
			}

			ipMap[ip]++
		}

	}

	found := false

	for _, ip := range ipSlice {
		count := ipMap[ip]
		if count > threshold {
			fmt.Printf("%s: %d\n", ip, count)
			found = true
		}
	}

	if !found {
		fmt.Println("NO")
	}

}

func main() {
	//192.168.1.1 10.0.0.1 192.168.1.1 172.16.0.1 10.0.0.1 8.8.8.8 invalid_ip 256.0.0.1 192.168.-1.1 10.0.0.1 192.168.1.
	fmt.Println("Введите список ip адресов:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ipList := scanner.Text()

	fmt.Println("Введите пороговое значение:")
	scanner.Scan()
	threshold, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Пороговое значение указано неверно")
		return
	}

	processIPs(ipList, threshold)

}
