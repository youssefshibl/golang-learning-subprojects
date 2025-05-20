package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var taskFileName = "tasks.json"

func getCurrentPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return ""
	}
	return filepath.Join(cwd, taskFileName)
}

type fileInitArgs struct {
	filePath       string
	initialContent string
}

func createFileWithContent(args fileInitArgs) error {
	if err := os.WriteFile(args.filePath, []byte(args.initialContent), 0644); err != nil {
		return fmt.Errorf("failed to create file with content: %w", err)
	}
	return nil
}

// ReadTasksFromFile loads tasks from the JSON file, or creates it if missing.
func ReadTasksFromFile() ([]Task, error) {
	filePath := getCurrentPath()
	if filePath == "" {
		return nil, errors.New("could not determine task file path")
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s does not exist. Creating it...\n", taskFileName)
		if err := createFileWithContent(fileInitArgs{filePath: filePath, initialContent: "[]"}); err != nil {
			return nil, fmt.Errorf("failed to initialize task file: %w", err)
		}
		return []Task{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("error checking file status: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open task file: %w", err)
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("failed to decode task data: %w", err)
	}

	return tasks, nil
}

// WriteTasksToFile saves the tasks to the JSON file.
func WriteTasksToFile(tasks []Task) (bool, error) {
	filePath := getCurrentPath()
	if filePath == "" {
		return false, errors.New("could not determine task file path")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to create task file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(tasks); err != nil {
		return false, fmt.Errorf("failed to encode tasks to file: %w", err)
	}

	return true, nil
}
