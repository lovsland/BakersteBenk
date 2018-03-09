package Oppgave_1

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	fileMode os.FileMode
	err      error
)

type FileMode uint32

func fileIsRegular() {

	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {

		fmt.Println("The file does not exist")
	}

	if _, err := os.Stat("test.txt"); err == nil {

		fmt.Println("This is a Regular file")
	}

}


func ReadLink() {

	if _, err := os.Lstat("test.txt"); err == nil {

		fmt.Println("This is a symbolic link")

	} else {

		fmt.Println("This is not a symbolic link")
	}

}

func main() {

	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)


	}


	filestat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes := float64(filestat.Size())
	fmt.Printf("Information about file %s:\n", filestat.Name())
	fmt.Printf("Size: %.0f bytes\n", bytes)
	fmt.Printf("Size: %f kilobytes\n", bytes / 1024)
	fmt.Printf("Size: %f megabytes\n", bytes / (1024 * 1024))
	fmt.Printf("Size: %f gigabytes\n", bytes / (1024 * 1024 * 1024))

	if filestat.Mode() & os.ModeDevice != 0 {
		fmt.Println(" Is a device file")
	} else {
		fmt.Println(" Is not a device file")
	}



	file.Close()

	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Is not append only",)

	fileIsRegular()
	ReadLink()

}
