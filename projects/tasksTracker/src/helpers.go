package src

func Some(tasks []Task, check func(Task) bool) bool {
	for _, task := range tasks {
		if check(task) {
			return true
		}
	}
	return false
}

func filter(tasks []Task, filterfun func(Task) bool) []Task {
	var updatedTasks []Task

	for _, task := range tasks {
		if filterfun(task) {
			updatedTasks = append(updatedTasks, task)
		}
	}
	return updatedTasks
}

func find(tasks *[]Task, findFun func(Task) bool) *Task {

	for i := range *tasks {
		if findFun((*tasks)[i]) {
			return &(*tasks)[i]
		}
	}
	return nil
}
