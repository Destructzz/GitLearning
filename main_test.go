package main

import (
	"testing"
)

func TestNewTaskManager(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Error("NewTaskManager() returned nil")
	}
	if len(tm.tasks) != 0 {
		t.Error("New TaskManager should have empty tasks slice")
	}
	if tm.nextID != 1 {
		t.Error("New TaskManager should start with nextID = 1")
	}
}

func TestAddTask(t *testing.T) {
	tm := NewTaskManager()
	task := tm.AddTask("Test Task", "Test Description")

	if task.ID != 1 {
		t.Errorf("Expected task ID to be 1, got %d", task.ID)
	}
	if task.Title != "Test Task" {
		t.Errorf("Expected task title to be 'Test Task', got '%s'", task.Title)
	}
	if task.Description != "Test Description" {
		t.Errorf("Expected task description to be 'Test Description', got '%s'", task.Description)
	}
	if task.Completed {
		t.Error("New task should not be completed")
	}
}

func TestGetTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Test Task", "Test Description")

	task, err := tm.GetTask(1)
	if err != nil {
		t.Errorf("GetTask failed: %v", err)
	}
	if task.ID != 1 {
		t.Errorf("Expected task ID to be 1, got %d", task.ID)
	}

	_, err = tm.GetTask(999)
	if err == nil {
		t.Error("Expected error when getting non-existent task")
	}
}

func TestUpdateTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Original Title", "Original Description")

	task, err := tm.UpdateTask(1, "Updated Title", "Updated Description")
	if err != nil {
		t.Errorf("UpdateTask failed: %v", err)
	}
	if task.Title != "Updated Title" {
		t.Errorf("Expected updated title to be 'Updated Title', got '%s'", task.Title)
	}
	if task.Description != "Updated Description" {
		t.Errorf("Expected updated description to be 'Updated Description', got '%s'", task.Description)
	}

	_, err = tm.UpdateTask(999, "Title", "Description")
	if err == nil {
		t.Error("Expected error when updating non-existent task")
	}
}

func TestDeleteTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Test Task", "Test Description")

	err := tm.DeleteTask(1)
	if err != nil {
		t.Errorf("DeleteTask failed: %v", err)
	}

	tasks := tm.GetAllTasks()
	if len(tasks) != 0 {
		t.Error("Task was not deleted")
	}

	err = tm.DeleteTask(999)
	if err == nil {
		t.Error("Expected error when deleting non-existent task")
	}
}

func TestToggleTaskStatus(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Test Task", "Test Description")

	task, err := tm.ToggleTaskStatus(1)
	if err != nil {
		t.Errorf("ToggleTaskStatus failed: %v", err)
	}
	if !task.Completed {
		t.Error("Task status should be toggled to completed")
	}

	task, err = tm.ToggleTaskStatus(1)
	if err != nil {
		t.Errorf("ToggleTaskStatus failed: %v", err)
	}
	if task.Completed {
		t.Error("Task status should be toggled back to not completed")
	}

	_, err = tm.ToggleTaskStatus(999)
	if err == nil {
		t.Error("Expected error when toggling non-existent task")
	}
}

func TestGetAllTasks(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task 1", "Description 1")
	tm.AddTask("Task 2", "Description 2")

	tasks := tm.GetAllTasks()
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
} 