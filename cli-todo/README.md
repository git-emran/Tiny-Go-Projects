# CLI ToDo App

A simple, fast, and lightweight Command-Line Interface (CLI) task manager written in Go. It persists tasks to a local JSON file, supports interactive input, and accepts environment variable configuration.

## Features

- **Add Tasks:** Quickly add tasks via flags, positional arguments, or interactive standard input (`STDIN`).
- **List Tasks:** Display all pending and completed tasks with 1-based indexing.
- **Mark Complete:** Mark items as completed with a visual indicator (`X`) and automatic timestamping.
- **Configurable Storage:** Automatically saves tasks to `todo.json` by default, customizable via the `TODO_FILENAME` environment variable.
- **Zero External Dependencies:** Built entirely with standard Go packages and internal modules.

## Prerequisites

- [Go](https://go.dev/dl/) 1.18 or higher installed.

## Installation

Clone the repository and build the binary:

```bash
git clone [https://github.com/git-emran/tiny-go-projects.git](https://github.com/git-emran/tiny-go-projects.git)
cd tiny-go-projects/cli-todo
go build -o todo main.go
