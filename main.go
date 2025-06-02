package main

import (
	"fmt"
	"sync"
)

// Task представляет собой структуру задачи
type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

// TaskManager управляет задачами
type TaskManager struct {
	tasks  []Task
	nextID int
	mutex  sync.RWMutex
}

// NewTaskManager создает новый менеджер задач
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

// AddTask добавляет новую задачу
func (tm *TaskManager) AddTask(title, description string) Task {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	task := Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	tm.tasks = append(tm.tasks, task)
	tm.nextID++
	return task
}

// GetTask возвращает задачу по ID
func (tm *TaskManager) GetTask(id int) (Task, error) {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()

	for _, task := range tm.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

// UpdateTask обновляет существующую задачу
func (tm *TaskManager) UpdateTask(id int, title, description string) (Task, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Title = title
			tm.tasks[i].Description = description
			return tm.tasks[i], nil
		}
	}
	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

// DeleteTask удаляет задачу по ID
func (tm *TaskManager) DeleteTask(id int) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// ToggleTaskStatus переключает статус выполнения задачи
func (tm *TaskManager) ToggleTaskStatus(id int) (Task, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Completed = !tm.tasks[i].Completed
			return tm.tasks[i], nil
		}
	}
	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

// GetAllTasks возвращает все задачи
func (tm *TaskManager) GetAllTasks() []Task {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()

	tasks := make([]Task, len(tm.tasks))
	copy(tasks, tm.tasks)
	return tasks
}

func main() {
	tm := NewTaskManager()
	// Пример использования
	task := tm.AddTask("Тестовая задача", "Описание тестовой задачи")
	fmt.Printf("Создана задача: %+v\n", task)
	
	tasks := tm.GetAllTasks()
	fmt.Printf("Всего задач: %d\n", len(tasks))

	updatedTask, err := tm.UpdateTask(task.ID, "Обновленная задача", "Описание обновленной задачи")
	if err != nil {
		fmt.Printf("Ошибка при обновлении задачи: %v\n", err)
	} else {
		fmt.Printf("Задача обновлена: %+v\n", updatedTask)
	}

	err = tm.DeleteTask(task.ID)
	if err != nil {
		fmt.Printf("Ошибка при удалении задачи: %v\n", err)
	} else {
		fmt.Println("Задача удалена")
	}
}