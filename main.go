package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//data := make([]byte, 100) //this is called a byte slice

	err := writeData(`my name is Miracle`)
	if err != nil {
		log.Fatal(err)
	}

	//file, theOpenFileError := os.Open(`C:\Users\DELL\Desktop\itemData.txt`)
	//if theOpenFileError != nil {
	//	log.Fatal("error opening file")
	//}

	theData, err := os.ReadFile(`C:\Users\DELL\Desktop\itemData.txt`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(theData))
}

func writeData(content string) error {
	err := os.WriteFile(`C:\Users\DELL\Desktop\itemData.txt`, []byte(content), 0666)
	if err != nil {
		return err
	}
	return nil
}
