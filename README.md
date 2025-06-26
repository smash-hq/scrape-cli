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

### From source (requires Go 1.18+)

```bash
go install github.com/smash-hq/scrape-cli@latest
```

## 🚀 Usage

### 📌 1. Interactive Mode

Just run:

```
scrape-cli --create
```

You’ll be guided to:

- Select a template (e.g. start_with_golang, start_with_node)
- Input a project folder name

**Useful when you're unsure about flags or want a guided experience.**

### 📌 2. Non-interactive Mode

Fully automate project generation with flags:

```
scrape-cli --tmpl start_with_golang --name my-actor
```

Creates a folder **my-actor** using the Golang actor template.
Run your actor:
```
cd my-actor
scrape-cli --run
```

### 📌 3. Show Version

```
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

```
go run main.go --create
```