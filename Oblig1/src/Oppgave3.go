
package main

import "fmt"
import "os"


func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)



	go func() {
		for {

		}
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()


	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
