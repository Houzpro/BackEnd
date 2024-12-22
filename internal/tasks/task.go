package tasks

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID       string `json:"id"`
	Status   string `json:"status"` // pending, in_progress, done, error, cancelled
	Filename string `json:"filename"`
}

var tasks = make(map[string]*Task)
var mutex = &sync.Mutex{}
var maxConcurrentTasks = 5
var semaphore = make(chan struct{}, maxConcurrentTasks)

func CreateTask() string {
	taskID := generateID()
	task := &Task{
		ID:     taskID,
		Status: "pending",
	}
	mutex.Lock()
	tasks[taskID] = task
	mutex.Unlock()
	log.Printf("Task %s created", taskID)
	return taskID
}

func GetTask(taskID string) *Task {
	mutex.Lock()
	defer mutex.Unlock()
	return tasks[taskID]
}

func RunTask(taskID string, cancelChan chan struct{}) {
	semaphore <- struct{}{}        // Acquire a semaphore slot
	defer func() { <-semaphore }() // Release the semaphore slot

	task := GetTask(taskID)
	if task == nil {
		return
	}

	log.Printf("Task %s started", taskID)
	task.Status = "in_progress"
	filename := "export_" + taskID + ".json"
	task.Filename = filename

	select {
	case <-time.After(5 * time.Second):
		log.Printf("Task %s is writing to file", taskID)
		file, err := os.Create(filename)
		if err != nil {
			log.Printf("Task %s failed to create file: %v", taskID, err)
			task.Status = "error"
			return
		}
		defer file.Close()

		data := map[string]string{"message": "Data exported successfully"}
		json.NewEncoder(file).Encode(data)

		task.Status = "done"
		log.Printf("Task %s completed successfully", taskID)
	case <-cancelChan:
		log.Printf("Task %s was cancelled", taskID)
		task.Status = "cancelled"
	}
}

func CancelTask(taskID string) {
	task := GetTask(taskID)
	if task == nil {
		return
	}

	log.Printf("Cancelling task %s", taskID)
	task.Status = "cancelled"
}

func generateID() string {
	return time.Now().Format("20060102150405")
}

func CreateTaskHandler(c *gin.Context) {
	taskID := CreateTask()
	c.JSON(200, gin.H{"task_id": taskID})
}

func GetTaskHandler(c *gin.Context) {
	taskID := c.Param("id")
	task := GetTask(taskID)
	if task == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}

func CancelTaskHandler(c *gin.Context) {
	taskID := c.Param("id")
	CancelTask(taskID)
	c.JSON(200, gin.H{"status": "Task cancelled"})
}
