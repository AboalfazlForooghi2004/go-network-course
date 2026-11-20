package main


import	"fmt"

func  main(){
	host:="127.0.0.1"
	port:=8080
	protocol:="TCP"

	fmt.Println("hello,go networl programming")
	fmt.Println("Port:", port)
	fmt.Println("Host:", host)
	fmt.Println("Protocol:", protocol)


	fmt.Printf("server address: %s:%d\n", host,port)

}
