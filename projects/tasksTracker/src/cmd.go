package src

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
		Long: `Task Tracker is a CLI tool for managing tasks. It allows you to create, list, and delete tasks.
    
You can also mark tasks as completed and update their status.
Complete code available at https://github.com/arikchakma/backend-projects`,
	}

	cmd.AddCommand(NewTaskCMD())
	cmd.AddCommand(DeleteTaskCMD())
	cmd.AddCommand(ListAllTasksCMD())
	cmd.AddCommand(UpdateStatusCMD())
	cmd.AddCommand(UpdateDesCMD())

	return cmd
}

func NewTaskCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "add [description]",
		Short: "Add a task to the task list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return AddTask(args[0])
		},
	}
}

func DeleteTaskCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task from the task list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			taskID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid task ID: %w", err)
			}
			return DeleteTask(taskID)
		},
	}
}

func ListAllTasksCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks (optionally filtered by status)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			status := "All"
			if len(args) == 1 {
				status = args[0]
			}
			return ListAllTasks(TaskStatus(status))
		},
	}
}

func UpdateStatusCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "updateStatus [id] [status]",
		Short: "Update the status of a task by ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			taskID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid task ID: %w", err)
			}
			return UpdateTaskStatus(taskID, TaskStatus(args[1]))
		},
	}
}

func UpdateDesCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "updateDes [id] [new description]",
		Short: "Update the description of a task by ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			taskID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid task ID: %w", err)
			}
			return UpdateTaskSDescription(taskID, args[1])
		},
	}
}
