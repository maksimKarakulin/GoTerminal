package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		
		// Remove newline character
		input = strings.TrimSpace(input)
		
		// Handle commands
		switch input {
		case "exit", "quit":
			fmt.Println("Goodbye!")
			os.Exit(0)
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  help - Show this help message")
			fmt.Println("  exit/quit - Exit the program")
			fmt.Println("  echo - Echo your input")
		case "echo":
			fmt.Print("Enter text to echo: ")
			echoInput, _ := reader.ReadString('\n')
			fmt.Println("You said:", strings.TrimSpace(echoInput))
		default:
			fmt.Printf("Unknown command: %q. Type 'help' for available commands.\n", input)
		}
	}
}
