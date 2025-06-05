package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	handler     func(args []string)
}

var commands = make(map[string]command)
var history []string

func main() {
	registerCommand("help", "Show help", cmdHelp)
	registerCommand("exit", "Exit the program", cmdExit)
	registerCommand("history", "Show command history", cmdHistory)
	registerCommand("clear", "Clear command history", cmdClear)
	registerCommand("echo", "Echo the input", cmdEcho)
	registerCommand("add", "Add two numbers", cmdAdd)
	registerCommand("multiply", "Multiply two numbers", cmdMultiply)
	registerCommand("subtract", "Subtract two numbers", cmdSubtract)

	rl, err := readline.New("> ")
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		os.Exit(1)
	}
	defer rl.Close()

	color.Cyan("Welcome to goTerminal!!!!!!!!!!!")
	color.Yellow("Type 'help' to see available commands.")

	for {
		input, err := rl.Readline()
		if err != nil { // Handles Ctrl+D (EOF)
			break
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		history = append(history, input)
		parts := strings.Fields(input)
		cmdName := parts[0]
		args := parts[1:]

		if cmd, exists := commands[cmdName]; exists {
			cmd.handler(args)
		} else {
			color.Red("Unknown command: %s", cmdName)
		}
	}
}

func registerCommand(name, description string, handler func(args []string)) {
	commands[name] = command{
		name:        name,
		description: description,
		handler:     handler,
	}
}

func cmdHelp(args []string) {
	color.Green("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
}

func cmdExit(args []string) {
	color.Yellow("Exiting the program. Goodbye!")
	os.Exit(0)
}
func cmdHistory(args []string) {
	if len(history) == 0 {
		color.Yellow("No command history available.")
		return
	}
	color.Green("Command History:")
	for i, cmd := range history {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}
func cmdClear(args []string) {
	history = []string{}
	color.Green("Command history cleared.")
}

func cmdEcho(args []string) {
	if len(args) == 0 {
		color.Red("Usage: echo <message>")
		return
	}
	message := strings.Join(args, " ")
	color.Blue("Echo: %s", message)
}

func cmdAdd(args []string) {
	if len(args) != 2 {
		color.Red("Usage: add <num1> <num2>")
		return
	}
	var num1, num2 float64
	_, err1 := fmt.Sscanf(args[0], "%f", &num1)
	_, err2 := fmt.Sscanf(args[1], "%f", &num2)
	if err1 != nil || err2 != nil {
		color.Red("Both arguments must be numbers.")
		return
	}
	result := num1 + num2
	color.Blue("Result: %f", result)
}

func cmdMultiply(args []string) {
	if len(args) != 2 {
		color.Red("Usage: multiply <num1> <num2>")
		return
	}
	var num1, num2 float64
	_, err1 := fmt.Sscanf(args[0], "%f", &num1)
	_, err2 := fmt.Sscanf(args[1], "%f", &num2)
	if err1 != nil || err2 != nil {
		color.Red("Both arguments must be numbers.")
		return
	}
	result := num1 * num2
	color.Blue("Result: %f", result)
}
func cmdSubtract(args []string) {
	if len(args) != 2 {
		color.Red("Usage: subtract <num1> <num2>")
		return
	}
	var num1, num2 float64
	_, err1 := fmt.Sscanf(args[0], "%f", &num1)
	_, err2 := fmt.Sscanf(args[1], "%f", &num2)
	if err1 != nil || err2 != nil {
		color.Red("Both arguments must be numbers.")
		return
	}
	result := num1 - num2
	color.Blue("Result: %f", result)
}
