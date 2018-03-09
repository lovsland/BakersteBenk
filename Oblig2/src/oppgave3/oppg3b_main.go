package main

import (
	"./oppg3b"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

func main() {
	signal.Notify(sigintCh, os.Interrupt, syscall.SIGINT)
	go ctrlc()
	oppg3b.WriteNumbersToFile(oppg3b.Input())
	oppg3b.WriteSum(oppg3b.ReadFromFile())
	oppg3b.PrintSum()
	time.Sleep(1*time.Second)
}

var sigintCh = make(chan os.Signal, 2)

func ctrlc() {
	<- sigintCh
	fmt.Println("CTRL+C stoppet programmet")
	os.Exit(1)
}