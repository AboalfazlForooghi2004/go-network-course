package main

import "fmt"

func main() {
	host := "127.0.0.1"
	port := 8080
	protocol := "TCP"

	letancy := 12.5
	isActiv := true

	connectionCount := 0

	fmt.Println("network Service Info")
	fmt.Println("----------------------")
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Protocol: %s\n", protocol)

	fmt.Printf("Latency: %.1f ms\n", letancy)
	fmt.Printf("Active: %t\n", isActiv)
	fmt.Printf("connections: %d\n", connectionCount)

	fmt.Println("--------------------------------------------")

	networkService()
}
