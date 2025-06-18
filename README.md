Here's a `README.md` describing the features of your Go terminal application:

---

# GoTerminal

`GoTerminal` is a pretty simple interactive command-line interface built in Go. It supports basic commands, math operations, and command history features. It uses the `readline` library for input handling and `color` for colorful terminal output.

## 🚀 Features

### ✅ General Commands

* **`help`**
  Display all available commands with their descriptions.

* **`exit`**
  Exit the terminal gracefully.

* **`history`**
  Show a history of all entered commands during the session.

* **`clear`**
  Clear the command history.

* **`echo <message>`**
  Repeats (echoes) the message back to the user.

### ➕ Math Operations

* **`add <num1> <num2>`**
  Adds two numbers and displays the result.

* **`subtract <num1> <num2>`**
  Subtracts the second number from the first and displays the result.

* **`multiply <num1> <num2>`**
  Multiplies two numbers and displays the result.

### 🎨 Colorful Output

All responses and messages are color-coded using the [`fatih/color`](https://github.com/fatih/color) library for better readability and a more engaging CLI experience.

### ⌨️ Readline Support

Powered by [`chzyer/readline`](https://github.com/chzyer/readline), the terminal supports:

* Line editing
* Command history navigation using arrow keys
* Auto line completion (will be added in next commit)

## 🛠 Installation & Usage

1. Make sure you have Go installed.
2. Install required packages:

   ```bash
   go get github.com/chzyer/readline
   go get github.com/fatih/color
   ```
3. Run the application:

   ```bash
   go run main.go
   ```

## 🧩 Future Improvements

* AI Prompts
* Frontend with database implementation
* Hosting online
* .gitignore the .env files because obviously you need to do that. 

---

