package main

import (
	"fmt"
	"os"

	"learn.reboot01.com/git/moadwan/go-reloaded/autocorrect"
)

func main() {
	//declare the edited text from the beginning so it has its memory allocation and prevent errors
	var content []byte

	// check the os arguments
	if len(os.Args) == 1 {
		fmt.Println("please provide file names\nUsage: go run main.go <input_file> <output_file>")
		return
	}
	if len(os.Args) == 2 {
		fmt.Println("please provide another file name\nUsage: go run main.go <input_file> <output_file>")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("too many arguments\nUsage: go run main.go <input_file> <output_file>")
		return
	}

	//open the given files
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	//check if the output file exists
	fileStatus, writeerr := os.Stat(outputFile)
	if os.IsNotExist(writeerr) {
		fmt.Println("Output file doesn't exist, it will be created as", outputFile)
		writeerr = os.WriteFile(outputFile, content, 0644)
	}

	//check if the output file is empty
	fileStatus, writeerr = os.Stat(outputFile)
	if writeerr != nil {
		fmt.Println(writeerr)
	}
	if fileStatus.Size() != 0 {
		fmt.Println("Output file is not empty, overwriting will occur!")
	}

	//read the contents of the input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		return
	}

	//do the modifications
	toEdit := string(content)
	toEdit = autocorrect.ModifyLine(toEdit)

	//write the contents to the output file
	err = os.WriteFile(outputFile, []byte(toEdit), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %s\n", err)
		return
	}
}
