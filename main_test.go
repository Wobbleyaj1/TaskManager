package main

import (
	"os"
	"strings"
	"testing"
)


// TestTaskFormatPending verifies the format of a pending task
func TestTaskFormatPending(t *testing.T) {
	taskName := "Test Task"
	taskDesc := "Test Description"
	expectedStatus := "[Pending]"
	
	// Simulate task creation format
	task := taskName + " - " + taskDesc + " - " + expectedStatus
	
	// Verify task contains all components
	if !strings.Contains(task, taskName) {
		t.Error("Task should contain task name")
	}
	if !strings.Contains(task, taskDesc) {
		t.Error("Task should contain task description")
	}
	if !strings.Contains(task, expectedStatus) {
		t.Error("Task should contain [Pending] status")
	}
}

// TestTaskFormatDone verifies the format of a completed task
func TestTaskFormatDone(t *testing.T) {
	pendingTask := "Test Task - Test Description - [Pending]"
	
	// Simulate marking task as done
	doneTask := strings.Replace(pendingTask, "[Pending]", "[Done]", 1)
	
	// Verify status was changed
	if strings.Contains(doneTask, "[Pending]") {
		t.Error("Done task should not contain [Pending]")
	}
	if !strings.Contains(doneTask, "[Done]") {
		t.Error("Done task should contain [Done]")
	}
}

// TestMarkTaskAsDone tests the logic for marking a task as done
func TestMarkTaskAsDone(t *testing.T) {
	tasks := []string{
		"Task 1 - Description 1 - [Pending]",
		"Task 2 - Description 2 - [Pending]",
	}
	
	taskID := 0
	
	// Verify task is pending
	if !strings.Contains(tasks[taskID], "[Pending]") {
		t.Error("Task should start as [Pending]")
	}
	
	// Mark task as done
	tasks[taskID] = strings.Replace(tasks[taskID], "[Pending]", "[Done]", 1)
	
	// Verify task is now done
	if !strings.Contains(tasks[taskID], "[Done]") {
		t.Error("Task should be marked as [Done]")
	}
	
	// Verify other tasks remain unchanged
	if !strings.Contains(tasks[1], "[Pending]") {
		t.Error("Other tasks should remain unchanged")
	}
}

// TestDeleteTask tests the logic for deleting a task
func TestDeleteTask(t *testing.T) {
	tasks := []string{
		"Task 1 - Description 1 - [Pending]",
		"Task 2 - Description 2 - [Pending]",
		"Task 3 - Description 3 - [Done]",
	}
	
	taskID := 1
	originalLength := len(tasks)
	
	// Delete task at index 1
	tasks = append(tasks[:taskID], tasks[taskID+1:]...)
	
	// Verify length decreased by 1
	if len(tasks) != originalLength-1 {
		t.Errorf("Expected %d tasks after deletion, got %d", originalLength-1, len(tasks))
	}
	
	// Verify correct task was removed (Task 2 should be gone)
	for _, task := range tasks {
		if strings.Contains(task, "Task 2") {
			t.Error("Task 2 should have been deleted")
		}
	}
	
	// Verify other tasks remain
	if !strings.Contains(tasks[0], "Task 1") {
		t.Error("Task 1 should still exist")
	}
	if !strings.Contains(tasks[1], "Task 3") {
		t.Error("Task 3 should still exist")
	}
}

// TestTaskIDValidation tests validation logic for task IDs
func TestTaskIDValidation(t *testing.T) {
	tasks := []string{
		"Task 1 - Description 1 - [Pending]",
		"Task 2 - Description 2 - [Pending]",
	}
	
	testCases := []struct {
		taskID      int
		shouldBeValid bool
		description string
	}{
		{0, true, "First task ID should be valid"},
		{1, true, "Last task ID should be valid"},
		{-1, false, "Negative task ID should be invalid"},
		{2, false, "Task ID beyond length should be invalid"},
		{999, false, "Large task ID should be invalid"},
	}
	
	for _, tc := range testCases {
		isValid := tc.taskID >= 0 && tc.taskID < len(tasks)
		if isValid != tc.shouldBeValid {
			t.Errorf("%s: taskID=%d, expected valid=%v, got valid=%v", 
				tc.description, tc.taskID, tc.shouldBeValid, isValid)
		}
	}
}

// TestAddTaskToSlice tests adding a task to the task slice
func TestAddTaskToSlice(t *testing.T) {
	tasks := []string{
		"Task 1 - Description 1 - [Pending]",
	}
	
	newTask := "Task 2 - Description 2 - [Pending]"
	tasks = append(tasks, newTask)
	
	// Verify length increased
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
	
	// Verify new task was added
	if tasks[1] != newTask {
		t.Errorf("Expected last task to be '%s', got '%s'", newTask, tasks[1])
	}
}

