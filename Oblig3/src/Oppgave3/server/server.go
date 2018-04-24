package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	go tcp()
	go udp()

	for {
		time.Sleep(10000000)
	}
}

const quoteOfTheDay = "Hei hei"

func tcp() {

	fmt.Println("Launching TCP server...")
	ln, err := net.Listen("tcp", ":17")
	checkErr(err)

	for {
		conn, _ := ln.Accept()
		fmt.Println("TCP client connected!")
		conn.Write([]byte(quoteOfTheDay + "\n"))
		conn.Close()
	}
}

func udp() {

	fmt.Println("Launching UDP server...")
	serverAddr, err := net.ResolveUDPAddr("udp", ":17")
	checkErr(err)
	conn, err := net.ListenUDP("udp", serverAddr)
	checkErr(err)

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		_, address, err := conn.ReadFromUDP(buf)
		fmt.Println("UDP client connected!")
		conn.WriteToUDP([]byte(quoteOfTheDay+"\n"), address)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}