package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var aString string = "This is Go"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
		fmt.Printf("Variable type %T", aFloat)
	}

}
