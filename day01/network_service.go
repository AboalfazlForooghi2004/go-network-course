package main

import "fmt"

func networkService() {

	serviceName := "ssh"
	host := "192.168.1.10"
	port := 22
	isRunning := true
	latency := 4.75
	maxConnections := 100

	fmt.Println("network Service Info")
	fmt.Println("----------------------")
	fmt.Printf("Server: %s\n", serviceName)
	fmt.Printf("Server type: %T\n", serviceName)

	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Host type: %T\n", host)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Port type: %T\n", port)
	fmt.Printf("Running: %t\n", isRunning)
	fmt.Printf("Running Type: %T\n", isRunning)
	fmt.Printf("Latency: %.2f\n", latency)
	fmt.Printf("Latency type: %T\n", latency)
	fmt.Printf("Max Connections: %d\n", maxConnections)
	fmt.Printf("Max Connections type: %T\n", maxConnections)
}
