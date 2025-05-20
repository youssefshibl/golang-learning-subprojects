package src

import (
	"errors"
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
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("task description is required")
			}

			return AddTask(args[0])

		},
	}
	return cmd
}

func DeleteTaskCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a task from the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("task id is required")
			}

			taskIDInt, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				return err
			}
			return DeleteTask(taskIDInt)

		},
	}
	return cmd
}

func ListAllTasksCMD() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return ListAllTasks(TaskStatus(args[0]))
			}
			return ListAllTasks("All")

		},
	}
	return cmd
}

func UpdateStatusCMD() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "updateStatus",
		Short: "update status of task by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return ListAllTasks(TaskStatus(args[0]))
			}
			taskIDInt, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				return err
			}
			return UpdateTaskStatus(taskIDInt, TaskStatus(args[1]))

		},
	}
	return cmd
}

func UpdateDesCMD() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "updateDes",
		Short: "update description of task by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return ListAllTasks(TaskStatus(args[0]))
			}
			taskIDInt, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				return err
			}
			return UpdateTaskSDescription(taskIDInt, args[1])

		},
	}
	return cmd
}
