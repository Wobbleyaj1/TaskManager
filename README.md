# Task Management Application

## Project Overview

GoWorld is a command-line task management application written in Go. It provides a simple and intuitive interface for managing your daily tasks with features including task creation, listing, status tracking, and deletion.

### Features

- **Add Tasks**: Create new tasks with a name and description
- **List Tasks**: View all tasks with their current status (Pending/Done)
- **Mark as Done**: Update task status from Pending to Done
- **Delete Tasks**: Remove tasks from your list
- **Interactive Menu**: Easy-to-use command-line interface

### Design Choices

The application uses a simple slice-based data structure to store tasks, with persistent storage in a text file. Each task is stored as a formatted string containing:

- Task Name
- Task Description
- Status (Pending or Done)

This design provides:

- Simplicity and ease of implementation
- Fast in-memory operations
- Clear separation of concerns with dedicated functions for each operation
- Persistent storage across program runs via text file
- Human-readable task storage format

## Usage Instructions

When you run the application, you'll see a command prompt. Type one of the following commands:

### Available Commands

| Command | Action              |
| ------- | ------------------- |
| `a`     | Add a new task      |
| `l`     | List all tasks      |
| `m`     | Mark a task as done |
| `d`     | Delete a task       |
| `e`     | Exit the program    |
| `h`     | Show help menu      |

### Workflow Examples

#### Adding a Task

1. Type `a` and press Enter
2. Enter the task name when prompted
3. Enter the task description when prompted
4. The task will be added with status `[Pending]`

#### Listing Tasks

1. Type `l` and press Enter
2. All tasks will be displayed with their ID, name, description, and status

#### Marking a Task as Done

1. Type `m` and press Enter
2. Enter the task ID number when prompted
3. The task status will change from `[Pending]` to `[Done]`

#### Deleting a Task

1. Type `d` and press Enter
2. Enter the task ID number when prompted
3. The task will be removed from the list

## How to Build and Run the Application

### Prerequisites

- Go 1.16 or higher installed on your system
- Basic familiarity with the command line

### Installation

1. Clone or download the project to your local machine
2. Navigate to the project directory:
   ```bash
   cd /path/to/GoWorld
   ```

### Running the Application

#### Option 1: Run Directly (Recommended for Development)

```bash
go run main.go
```

#### Option 2: Build and Execute

```bash
# Build the executable
go build -o goworld

# Run the executable
./goworld
```

### Running Tests

The project includes comprehensive unit tests for core functionality:

```bash
# Run all tests
go test

# Run tests with coverage
go test -cover

# Run tests with detailed coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Code Structure

```
GoWorld/
├── main.go          # Main application code
├── main_test.go     # Unit tests for core functionality
├── tasks.txt        # Persistent task storage (created automatically)
├── go.mod           # Go module definition
├── go.sum           # Dependency checksums
└── README.md        # This file
```

## Technical Details

### Core Functions

- `loadTasks()`: Reads tasks from the text file on program startup
- `saveTasks()`: Writes all tasks to the text file after any modification
- `addTask()`: Handles task creation with name and description input
- `listTasks()`: Displays all tasks in a formatted list
- `markTask()`: Updates task status from Pending to Done
- `deleteTask()`: Removes a task by its ID
- `helpMenu()`: Displays available commands
- `run()`: Main program loop handling user input and command routing

### Data Flow

1. **On Startup**: Program loads existing tasks from `tasks.txt` file
2. User enters a command
3. `run()` function processes the command through a switch statement
4. Appropriate function is called with current tasks slice
5. Function performs operation and returns updated tasks slice
6. **After Modifications**: Tasks are automatically saved to `tasks.txt`
7. Program loops back to accept next command

### File Storage Format

Tasks are stored in `tasks.txt` with one task per line in the format:

```
Task Name - Task Description - [Status]
```

Example:

```
Buy groceries - Get milk, eggs, and bread - [Pending]
Finish homework - Complete math assignment - [Done]
```

## Future Enhancements

Potential features for future versions:

- Task priorities and due dates
- Task categories or tags
- Search and filter capabilities
- Export tasks to different formats (JSON, CSV)
- Undo/Redo functionality
- Task archiving for completed tasks

## Author

Created as a learning project for Go programming.

## License

This project is available for educational purposes.
