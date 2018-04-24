package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":17")
	if err != nil {
		fmt.Println(err)
	}

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
	conn.Close()
}