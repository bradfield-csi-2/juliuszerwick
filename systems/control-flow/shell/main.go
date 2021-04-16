package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("Entering shell program")

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

		// Take input variable and parse for commands instead of echoing back to user.
		userArgs := strings.Split(input, " ")
		//fmt.Printf("userArgs: %v\n", userArgs)
		processArgs(userArgs)

		// Echo back user input on a separate line.
		//fmt.Println(input)
	}
}

func processArgs(args []string) {
	var cmd *exec.Cmd

	if len(args) > 1 {
		input := strings.Join(args[1:], " ")
		//fmt.Printf("input to cmd: %v\n", input)
		cmd = exec.Command(args[0], input)
	} else {
		cmd = exec.Command(args[0])
	}

	//cmd.Stdin = strings.NewReader(input)

	var out bytes.Buffer
	cmd.Stdout = &out

	//err := cmd.Run()
	if err := cmd.Start(); err != nil {
		fmt.Printf("%s: command not found\n", args[0])
	}

	go signalHandler(cmd.Process)

	if err := cmd.Wait(); err != nil {
		fmt.Println()
	}
	//fmt.Printf("%q\n", out.String())
	fmt.Print(out.String())
}

func signalHandler(process *os.Process) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	process.Signal(sig)
}
