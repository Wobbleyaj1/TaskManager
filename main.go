package main

import (
	"bufio" // For reading user input
	"fmt" // For standard library
	"os" // For os operations such as exit
	"strings"
	"strconv" // Convert string to int
)

// File path constant for storing tasks persistently
const tasksFile = "tasks.txt"

// loadTasks reads tasks from the text file and returns them as a slice
func loadTasks() []string {
	tasks := []string{}
	
	// Open the file for reading
	file, err := os.Open(tasksFile)
	if err != nil {
		// If file doesn't exist, return empty slice (first run)
		if os.IsNotExist(err) {
			return tasks
		}
		fmt.Println("Error opening tasks file:", err)
		return tasks
	}
	defer file.Close() // Ensure file is closed when function returns
	
	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Only add non-empty lines
			tasks = append(tasks, line)
		}
	}
	
	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading tasks file:", err)
	}
	
	return tasks
}

// saveTasks writes all tasks to the text file
func saveTasks(tasks []string) error {
	// Create or truncate the file for writing
	file, err := os.Create(tasksFile)
	if err != nil {
		fmt.Println("Error creating tasks file:", err)
		return err
	}
	defer file.Close() // Ensure file is closed when function returns
	
	// Write each task on a new line
	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		_, err := writer.WriteString(task + "\n")
		if err != nil {
			fmt.Println("Error writing task:", err)
			return err
		}
	}
	
	// Flush the buffer to ensure all data is written
	return writer.Flush()
}

// Add Task: Allow users to add new tasks with a brief description.
func addTask(tasks []string) []string {
	// Prompt For Task To Be Added
	fmt.Print("Task Name: ")
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	taskName := strings.TrimSpace(input1)	// Trim whitespace/newline from input

	fmt.Print("Task Description: ")
	input2, _ := reader.ReadString('\n')
	taskDescription := strings.TrimSpace(input2) // Trim whitespace/newline from input

	// Combine name and description with Pending status
	newTask := taskName + " - " + taskDescription + " - [Pending]"

	tasks = append(tasks, newTask)
	
	// Save tasks to file after adding
	saveTasks(tasks)

	return tasks
}

// List Tasks: Display all existing tasks with their status (Pending/Done).
func listTasks(tasks []string) {

	fmt.Print("ID: Name - Description - Status")
	// Iterate over the slice
    for index, value := range tasks {
        fmt.Printf("\n %d: %s", index, value)
    }
	fmt.Println()
}

// Mark Task as Done: Enable users to mark a specific task as done by its task number.
func markTask(tasks []string) []string {
	// Prompt For Task ID To Be Marked As Done
	fmt.Print("Task ID To Mark As Done: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	taskID, err := strconv.Atoi(strings.TrimSpace(input)) // Trim whitespace/newline from input

	if err != nil {
		// Handle the error if the string is not a valid integer
		fmt.Println("Error during conversion:", err)
		return tasks
	}

	// Check if taskID is within valid range
	if taskID < 0 || taskID >= len(tasks) {
		fmt.Println("Error: Invalid task ID")
		return tasks
	}

	// Check if task is already marked as done
	if strings.Contains(tasks[taskID], " [Done]") {
		fmt.Println("Task is already marked as done")
		return tasks
	}

	// Mark task as done by replacing [Pending] with [Done]
	tasks[taskID] = strings.Replace(tasks[taskID], " [Pending]", " [Done]", 1)
	
	// Save tasks to file after marking as done
	saveTasks(tasks)
	
	listTasks(tasks)

	return tasks
}

// Delete Task: Permit deletion of a task by its task number.
func deleteTask(tasks []string) []string {
	// Prompt For Task ID TO Be Removed
	fmt.Print("Task ID To Be Removed: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	taskID, err := strconv.Atoi(strings.TrimSpace(input)) // Trim whitespace/newline from input

	if err != nil {
		// Handle the error if the string is not a valid integer
		fmt.Println("Error during conversion:", err)
		return tasks
	}

	// Check if taskID is within valid range
	if taskID < 0 || taskID >= len(tasks) {
		fmt.Println("Error: Invalid task ID")
		return tasks
	}

	//Remove taskID from tasks
	tasks = append(tasks[:taskID], tasks[taskID+1:]...)
	
	// Save tasks to file after deletion
	saveTasks(tasks)

	listTasks(tasks)

	return tasks
}

// Help/Usage: Provide a help menu that outlines how to use each command.
func helpMenu() {
	fmt.Println("Help Menu")
	fmt.Println("Button: Action\n a: Add task\n l: List Tasks\n m: Mark Task\n d: Delete Task\n e: Exit Program")
}

// Recursive program running loop
func run(tasks []string) {
	fmt.Print("------------------\nInput: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	command := strings.TrimSpace(input)	// Trim whitespace/newline from input

	switch command {
	// Add Task
	case "a":
		tasks := addTask(tasks)
		run(tasks)
	// List Tasks
	case "l":
		listTasks(tasks)
		run(tasks)
	// Mark Task As Done
	case "m":
		tasks := markTask(tasks)
		run(tasks)
	// Delete Task
	case "d":
		tasks := deleteTask(tasks)
		run(tasks)
	// Exit program
	case "e":
		fmt.Println("\nGoodbye!")
	default:
		helpMenu()
		run(tasks)
	}
}

func main() {
	// Load existing tasks from file on startup
	tasks := loadTasks()
	
	fmt.Printf("Loaded %d task(s) from file\n", len(tasks))
	fmt.Print("Type 'h' for help menu\n")
	run(tasks)
}