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
		processArgs(userArgs)
	}
}

/* For part 5 - Multiple Commands:
- extract cmd variable creation and operations into a separate function
- have the processArgs function handle tokenizing the user input
- after splitting input by strings, iterate through array of input
		- set first element to the "cmd" -> Ex: echo, whoami, cat
		- set following elements [1:] to args to "cmd"
		- if the following characters are found, mark as starting point for new cmd:
				- &&, ||, ;
		- For &&, the following "cmd" + args will be started after the previous cmd
			returns with a zero exit code.
		- For ||, the following "cmd" + args will be started after the previous cmd,
			but the previous cmd's stdout and stdin will need to be connected to the
			following cmd's stdout and stdin. Each subsequent cmd will run after the
			previous cmd returns a non-zero exit status.
		- For ;, simply execute each process sequentially regardless of exit code
			returned from each cmd.
*/
func processArgs(args []string) {
	var cmd *exec.Cmd

	if len(args) > 1 {
		input := strings.Join(args[1:], " ")
		cmd = exec.Command(args[0], input)
	} else {
		cmd = exec.Command(args[0])
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Start(); err != nil {
		fmt.Printf("%s: command not found\n", args[0])
	}

	go signalHandler(cmd.Process)

	if err := cmd.Wait(); err != nil {
		fmt.Println()
	}
	fmt.Print(out.String())
}

func signalHandler(process *os.Process) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	process.Signal(sig)
}
