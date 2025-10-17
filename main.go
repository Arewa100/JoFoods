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
	errTwo := writeData(`my name is Miracle ope`)
	if errTwo != nil {
		log.Fatal(errTwo)
	}

	theData, err := os.ReadFile(`C:\Users\DELL\Desktop\itemData.txt`)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(theData)
	//fmt.Println(string(theData))

	//var lines []string = []string{"opeyemi", "loveth"}
	//addDataToNewLine(lines)
}

func writeData(content string) error {
	err := os.WriteFile(`C:\Users\DELL\Desktop\itemData.txt`, []byte(content), 0666)
	if err != nil {
		return err
	}
	return nil
}

func addDataToNewLine(lines []string) {
	fmt.Println(lines)
}
