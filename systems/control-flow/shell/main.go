package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Entering shell program")

	/* The main loop
	- create an infinite loop
	- at the start of the loop, print the shell prompt
	- allow user input on the same line
	- once user writes the enter character (carriage return) '\r'
	- if user enters the EOF terminal control character, exit loop and program
	*/

	reader := bufio.NewReader(os.Stdin)

	for {
		// Shell prompt
		fmt.Printf("~> ")

		// Read user input until a newline is entered.
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("\nExiting shell, bye!")
			break
		}

		// Replace newline with nothing.
		input = strings.Replace(input, "\n", "", -1)

		if input == "exit" {
			fmt.Println("Exiting shell, bye!")
			break
		}

		// Echo back user input on a separate line.
		fmt.Println(input)
	}

}
