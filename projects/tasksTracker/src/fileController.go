package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

func getCurrentPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error current directory:", err)
	}
	return path.Join(cwd, string(TaskFileName))
}

type makeEmptyFileArgs struct {
	filePath       string
	initialContent string
}

func makeEmptyFile(m makeEmptyFileArgs) {
	content := []byte(m.initialContent)
	f, err := os.Create(m.filePath)
	if err != nil {
		fmt.Print(err)
	} else {
		os.WriteFile(m.filePath, content, os.ModeAppend.Perm())
	}
	f.Close()

}

func ReadTasskFromFile() ([]Task, error) {
	path := getCurrentPath()
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("%s file not exist ... will create it ... \n", TaskFileName)
			makeEmptyFile(makeEmptyFileArgs{filePath: path, initialContent: "[]"})
			return []Task{}, nil

		} else {
			fmt.Println(err)
			return nil, err
		}
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}
	return tasks, nil

}

func WriteTasksToFile(tasks []Task) (bool, error) {
	path := getCurrentPath()
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error when try to make file ", err)
		return false, err
	}
	// closed file after function execute
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Error when try to write in file ", err)
		return false, err
	}

	return true, nil

}
