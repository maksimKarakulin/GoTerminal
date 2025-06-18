# 🧮 GoTerminal

**GoTerminal** is a simple yet powerful interactive command-line interface (CLI) built with Go. It features a variety of general and arithmetic commands, colorful output for readability, command history, file I/O, configuration management, and support for piping.

---

## 🚀 Features

### ✅ General Commands

| Command          | Description                      |
| ---------------- | -------------------------------- |
| `help`           | Display all available commands   |
| `exit`           | Exit the terminal gracefully     |
| `history`        | Show history of entered commands |
| `clear`          | Clear command history            |
| `echo <message>` | Echo back the provided message   |

### ➕ Arithmetic Operations

| Command                  | Description                           |
| ------------------------ | ------------------------------------- |
| `arithmetic`             | List all arithmetic commands          |
| `add <num1> <num2>`      | Add two numbers                       |
| `subtract <num1> <num2>` | Subtract second number from the first |
| `multiply <num1> <num2>` | Multiply two numbers                  |
| `division <num1> <num2>` | Divide first number by the second     |
| `modulus <num1> <num2>`  | Return modulus of the two numbers     |

### 📂 File I/O

| Command                       | Description                        |
| ----------------------------- | ---------------------------------- |
| `cat <filename>`              | Display contents of a file         |
| `write <filename> <content>`  | Write content to a file            |
| `append <filename> <content>` | Append content to an existing file |

### 🧩 Pipe Support

* Chain multiple commands using:
  `pipe <cmd1> | <cmd2> | ...`
  Passes output of one command as input to the next.

### 🎨 Colorful Output

* Terminal output is color-coded using [`fatih/color`](https://github.com/fatih/color) for:

  * Prompts
  * Errors
  * Standard output

### ⌨️ Readline Integration

Powered by [`chzyer/readline`](https://github.com/chzyer/readline):

* Line editing
* Arrow key history navigation
* Command autocompletion *(coming soon)*

### ⚙️ Configuration & Themes

* `config`: View current config
* `set <property> <value>`: Update config
* `theme <theme-name>`: Change terminal theme
  Available themes: `default`, `dark`, `light`, `solarized`

---

## 🛠 Installation & Usage

### Prerequisites

* Go installed (version ≥ 1.16)

### Install Dependencies

```bash
go get github.com/chzyer/readline
go get github.com/fatih/color
```

### Run the Application

```bash
go run main.go
```

---

## 🧠 Future Improvements

* 🔮 AI Command Suggestions / Prompts
* 🌐 Frontend with database integration
* ☁️ Host online for remote terminal access
* 📂 `.gitignore` `.env` and other sensitive files (for security)
