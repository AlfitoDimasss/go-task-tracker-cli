# Task Tracker CLI

A simple Task Tracker application built with Go that runs in the terminal. This project allows users to manage tasks by adding, updating, deleting, and tracking their progress. Tasks are stored locally in a JSON file without using any external libraries or frameworks.

## Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as:
  - Todo
  - In Progress
  - Done
- List all tasks
- Persist data in a JSON file
- Interactive CLI (REPL-style)

## Tech Stack

- Go
- Standard Library only
- JSON for data persistence

## Project Structure

```text
internal/
├── handler/     # Business logic
├── helper/      # Storage and utility functions
├── model/       # Data models
└── tasks.json   # Task storage

main.go          # CLI entry point
```

## Architecture

The application is organized into four layers:

| Layer   | Responsibility                      |
| ------- | ----------------------------------- |
| main    | Command parsing and CLI interaction |
| handler | Application/business logic          |
| helper  | File storage and utility functions  |
| model   | Data structures and constants       |

Flow:

```text
User Input
    ↓
main
    ↓
handler
    ↓
helper
    ↓
tasks.json
```

## Installation

Clone the repository:

```bash
git clone <your-repository-url>
cd go-task-tracker-cli
```

Run the application:

```bash
go run .
```

Or build the executable:

```bash
go build -o task-cli
./task-cli
```

## Usage

When the application starts, you will see:

```text
=== Go Task Tracker CLI ===
Type 'help' to see command
Type 'exit' or 'quit' to close
=================================
```

### Add a Task

```text
task-cli add Buy groceries
```

Output:

```text
Task added successfully (ID: 1)
```

### List Tasks

```text
task-cli list
```

Example:

```text
ID   Status       Description
[1]  todo         Buy groceries
[2]  done         Learn Go
```

### Update a Task

```text
task-cli update 1 Buy groceries and cook dinner
```

### Delete a Task

```text
task-cli delete 1
```

### Mark Task as In Progress

```text
task-cli mark-in-progress 1
```

### Mark Task as Done

```text
task-cli mark-done 1
```

### Mark Task as Todo

```text
task-cli mark-todo 1
```

### Exit Application

```text
task-cli exit
```

## Data Storage

Tasks are stored in:

```text
internal/tasks.json
```

Example:

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "created_at": "2025-05-26T10:00:00Z",
    "updated_at": "2025-05-26T10:00:00Z"
  }
]
```

## Learning Objectives

This project was built to practice:

- Go structs and custom types
- File handling
- JSON encoding and decoding
- Error handling
- CLI application development
- Project organization and separation of concerns

## Future Improvements

- Better command validation
- Colored terminal output
- Search tasks
- Due dates and priorities
- Export tasks to CSV
