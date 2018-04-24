package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("udp", ":17")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(conn, "\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
}