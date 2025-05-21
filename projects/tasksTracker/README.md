# Task Tracker CLI

A command-line interface tool for managing your tasks efficiently.

## Overview

Task Tracker is a simple yet powerful CLI tool built in Go that helps you manage your tasks directly from the terminal. It provides functionality to create, list, update, and delete tasks with ease.

## Features

- **Add Tasks**: Create new tasks with descriptions
- **List Tasks**: View all tasks or filter by status
- **Update Tasks**: Change task descriptions or status
- **Delete Tasks**: Remove tasks from your list

## Installation

```bash
# Clone the repository
git clone https://github.com/youssefshibl/golang-learning-subprojects.git

# Navigate to the project directory
cd golang-learning-subprojects

# Build the project
go build -o task-tracker

# Optional: Move the binary to your PATH
mv task-tracker /usr/local/bin/
```

## Usage

### Basic Commands

```bash
# Get help
task-tracker --help

# Add a new task
task-tracker add "Complete project documentation"

# List all tasks
task-tracker list

# List tasks by status (All, Pending, Completed, In Progress)
task-tracker list Pending

# Update task status (by ID)
task-tracker updateStatus 1 Completed

# Update task description (by ID)
task-tracker updateDes 1 "Updated task description"

# Delete a task (by ID)
task-tracker delete 1
```

### Command Details

| Command | Description | Usage |
|---------|-------------|-------|
| `add` | Add a new task | `task-tracker add [description]` |
| `list` | List all tasks or filter by status | `task-tracker list [status]` |
| `updateStatus` | Update task status by ID | `task-tracker updateStatus [id] [status]` |
| `updateDes` | Update task description by ID | `task-tracker updateDes [id] [new description]` |
| `delete` | Delete a task by ID | `task-tracker delete [id]` |

## Task Statuses

The following task statuses are supported:
- `All` (used for filtering only)
- `Pending`
- `Done`
- `Todo`

## Project Structure

The project follows a clean architecture with commands defined using Cobra framework:

```
.
└── tasksTracker 
    ├── go.mod                          # Module file
    ├── go.sum                          # Dependency file
    ├── main.go                         # Entry point
    ├── README.md                       # Project documentation
    └── src
        ├── cmd.go                     # Command definitions
        ├── enums.go                   # Enum definitions
        ├── errors.go                  # Error handling
        ├── fileController.go          # File operations
        ├── helpers.go                 # Helper functions
        ├── interfaces.go              # Interface definitions
        └── taskController.go          # Task management

```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions