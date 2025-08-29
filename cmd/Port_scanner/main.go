package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var allowTarget = map[string]bool{
	"scanme.nmap.org": true,
	"127.0.0.1":       true,
	"ya.ru":           false,
}

func isTargetAllowed(target string) bool {
	ok := allowTarget[target]
	return ok
}

func parsePorts(portsStr string) (portsInt []int, err error) {
	var ports []string

	if strings.Contains(portsStr, "-") { //Обрабатываем как диапазон (например, "1-100")
		var startPortInt, endPortInt int

		ports = strings.Split(portsStr, "-")
		if len(ports) != 2 {
			return nil, errors.New("неверный формат диапазона")
		}

		startPortInt, err = strconv.Atoi(ports[0])
		if err != nil {
			return nil, errors.New("неверный начальный порт")
		}

		endPortInt, err = strconv.Atoi(ports[1])
		if err != nil {
			return nil, errors.New("неверный конечный порт")
		}

		if startPortInt < 1 || endPortInt > 65535 {
			return nil, errors.New("неверный диапазон портов")
		}

		for i := startPortInt; i <= endPortInt; i++ {
			portsInt = append(portsInt, i)
		}

	} else if strings.Contains(portsStr, ",") { //Обрабатываем как перечисление (например, "1,80,433")

		ports = strings.Split(portsStr, ",")

		for _, portStr := range ports {
			port, err := strconv.Atoi(portStr)
			if err != nil {
				return nil, fmt.Errorf("неверный порт: %s", portStr)
			}

			if port < 1 || port > 65535 {
				return nil, errors.New("неверный диапазон портов")
			}

			portsInt = append(portsInt, port)

		}

		return
	} else { //Обрабатываем как единичный порт
		port, err := strconv.Atoi(portsStr)
		if err != nil {
			return nil, fmt.Errorf("неверный порт: %s", portsStr)
		}
		if port < 1 || port > 65535 {
			return nil, fmt.Errorf("порт %d вне диапазона 1-65535", port)
		}

		portsInt = append(portsInt, port)

	}

	if len(portsInt) == 0 {
		return nil, errors.New("не указаны порты для сканирования")
	}

	return
}

func findOpenPort(openPort chan<- int, checkPort <-chan int, target string, wg *sync.WaitGroup, timeout int) {
	defer wg.Done()

	for port := range checkPort {
		if scanPort(target, port, timeout) {
			openPort <- port
		}
	}
}

func scanPort(host string, port int, timeout int) bool {

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}

func getServiceName(port int) string {
	services := map[int]string{
		21: "ftp", 22: "ssh", 23: "telnet", 25: "smtp", 53: "dns",
		80: "http", 110: "pop3", 143: "imap", 443: "https",
		993: "imaps", 995: "pop3s", 3306: "mysql", 3389: "rdp",
		5432: "postgresql", 6379: "redis", 27017: "mongodb",
	}

	if service, exists := services[port]; exists {
		return service
	}
	return "unknown"
}

func main() {

	var wg sync.WaitGroup
	target := "scanme.nmap.org"
	portsStr := "1-100"
	workers := 10
	timeout := 1

	fmt.Println("Введите целевой хост (Например scanme.nmap.org):")
	fmt.Scanln(&target)

	if !isTargetAllowed(target) {
		fmt.Println("Цель не разрешена для сканирования")
		fmt.Scanln()
		return
	}

	fmt.Println("Введите целевые порты (Например 1,80,433 или 1-100, или 80):")
	fmt.Scanln(&portsStr)

	portsInt, err := parsePorts(portsStr)
	if err != nil {
		fmt.Printf("Ошибка получения портов: %v\n", err)
		fmt.Scanln()
		return
	}

	fmt.Println("Введите количество параллельных соединений:")
	fmt.Scanln(&workers)
	fmt.Println("Введите время ожидания ответа от порта:")
	fmt.Scanln(&timeout)

	openPort := make(chan int, len(portsInt))
	checkPort := make(chan int, len(portsInt))

	wg.Add(workers)
	for range workers {
		go findOpenPort(openPort, checkPort, target, &wg, timeout)
	}

	for _, i := range portsInt {
		checkPort <- i
	}
	close(checkPort)

	wg.Wait()
	close(openPort)

	fmt.Println("Список открытых портов:")
	for port := range openPort {
		fmt.Printf("Порт:%d, Сервис: %s \n", port, getServiceName(port))
	}
	fmt.Scanln()
}
