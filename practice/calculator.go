package calculator

import (
	"bufio"
	"fmt"
	"os"
)

func calculator() {
	fmt.Println("Hi! Welcome to the calculator Application")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please give your first input: ")
	num1, _ := reader.ReadString('\n')

}
