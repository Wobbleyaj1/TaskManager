package main

import (
	"testing"
)

// TestFormatTask verifies task formatting
func TestFormatTask(t *testing.T) {
	result := formatTask("Buy groceries", "Get milk and eggs")
	expected := "Buy groceries - Get milk and eggs - [Pending]"
	
	if result != expected {
		t.Errorf("formatTask() = %q, want %q", result, expected)
	}
}

// TestMarkTaskDone verifies marking a task as done
func TestMarkTaskDone(t *testing.T) {
	task := "Buy groceries - Get milk and eggs - [Pending]"
	result := markTaskDone(task)
	expected := "Buy groceries - Get milk and eggs - [Done]"
	
	if result != expected {
		t.Errorf("markTaskDone() = %q, want %q", result, expected)
	}
}

// TestIsTaskDone verifies checking task completion status
func TestIsTaskDone(t *testing.T) {
	tests := []struct {
		task     string
		expected bool
	}{
		{"Task 1 - Description - [Pending]", false},
		{"Task 1 - Description - [Done]", true},
	}
	
	for _, test := range tests {
		result := isTaskDone(test.task)
		if result != test.expected {
			t.Errorf("isTaskDone(%q) = %v, want %v", test.task, result, test.expected)
		}
	}
}

// TestMarkTaskAsDone tests the logic for marking a task as done
func TestMarkTaskAsDone(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Task 1 - Description - [Pending]", "Task 1 - Description - [Done]"},
		{"Task 2 - Another task - [Pending]", "Task 2 - Another task - [Done]"},
	}
	
	for _, test := range tests {
		result := markTaskDone(test.input)
		if result != test.expected {
			t.Errorf("markTaskDone(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}

