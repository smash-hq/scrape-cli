# 🧰 scrape-cli

> A command-line tool to quickly scaffold [Scrapeless](https://github.com/smash-hq) actor templates using Golang or
> Node.js.

---

## ✨ Features

- ⚡ Instant project scaffolding
- 🤖 Built-in templates: Golang & Node.js
- 🎯 Interactive CLI like `npm init`
- 🔧 Extendable template system

---

## 📦 Installation

### Using the install script (recommended)
You can quickly install the latest release via our bash install script (Linux/macOS/Windows Git Bash):
```bash
bash <(curl -fsSL https://raw.githubusercontent.com/smash-hq/scrape-cli/main/install-scrape-cli.sh)
```
This script will:
- Automatically detect your OS and architecture
- Download the latest scrape-cli release
- Extract the binary and install it to ~/.local/bin
- Add ~/.local/bin to your PATH if it’s not already set (for Linux/macOS)
- Show instructions on how to add to PATH manually on Windows

### From source (requires Go 1.18+)

```bash
go install github.com/smash-hq/scrape-cli@latest
```

## 🚀 Usage

### 📌 1. Interactive Mode

Just run:

```bash
scrape-cli --create
```

You’ll be guided to:

- Select a template (e.g. start_with_golang, start_with_node)
- Input a project folder name

**Useful when you're unsure about flags or want a guided experience.**

### 📌 2. Non-interactive Mode

Fully automate project generation with flags:

```bash
scrape-cli --tmpl start_with_golang --name my-actor
```

Creates a folder **my-actor** using the Golang actor template.
Run your actor:
```bash
cd my-actor
scrape-cli --run
```

### 📌 3. Show Version

```bash
scrape-cli --version
```

## 🧩 Flags

| Flag        | Short | Description                                                |
|-------------|-------|------------------------------------------------------------|
| `--create`  | `-c`  | Launch interactive template selection and naming           |
| `--tmpl`    | `-t`  | Choose a template (`start_with_golang`, `start_with_node`) |
| `--name`    | `-n`  | Set the project folder name (default: `my-actor`)          |
| `--version` | `-v`  | Print the version number of `scrape-cli`                   |
| `--run`     | `-r`  | Quickly launch your actor                                  |

## 📸 Example

```
scrape-cli --create
# Use the arrow keys to navigate: ↓ ↑ → ←
# ? Select a template:  [Use arrows]
#  > start_with_golang
#    start_with_node
#
# ? Project name: my-actor

# Output:
# Template Source: https://github.com/scrapeless-ai/actor-template-go.git
# Template generated in \your_workbase\my_actor
# Project 'my_actor' created using 'start_with_golang' template.

cd my-actor
scrape-cli --run
# Output:
# Launch my-actor logs...
```

## 🛠️ Development

```bash
go run main.go --create
```