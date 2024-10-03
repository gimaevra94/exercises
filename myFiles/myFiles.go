package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	FileIsExist()
	ImmediatelyTextReader()
	LineByLineTextReader()
	TextWriter()
	FileRename()
}

func FileIsExist() {
	fmt.Println("\n\n\n")
	if _, err := os.Stat("forMyFiles.txt"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
	fmt.Println("\n\n\n")
}

func ImmediatelyTextReader() {
	b, err := ioutil.ReadFile("forMyFiles.txt")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	fmt.Println("\n\n\n")
	fmt.Println(str,"\n\n\n")
}

func LineByLineTextReader() {
	file, err := os.Open("forMyFiles.txt")

	if err != nil {
		return
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		fmt.Println(i, line)
	}
	fmt.Println("\n\n\n")
}

func TextWriter()  {
	file,err:=os.Create("text.txt")

	if err!=nil{
		return
	}

	defer file.Close()

	file.WriteString("пошел нахуй")
}

func FileRename()  {
	os.Rename("text.txt","пошелнахуй.txt")
}