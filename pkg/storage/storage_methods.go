package storage

import (
	"github.com/tasks_api/pkg/models"
)

func CreateTask(task *models.Task) error {
	DB.Mu.Lock()
	defer DB.Mu.Unlock()
	DB.Db[task.Title] = task

	return nil
}

func GetAllTasksFromStorage(tasksArr *[]*models.Task) error {
	DB.Mu.Lock()
	defer DB.Mu.Unlock()
	for _, task := range DB.Db {
		*tasksArr = append(*tasksArr, task)
	}
	return nil
}

func GetTaskByTitle(title string) (*models.Task, error) {
	DB.Mu.Lock()
	defer DB.Mu.Unlock()
	return DB.Db[title], nil
}

func DeleteTaskFromStorage(title string) error {
	DB.Mu.Lock()
	defer DB.Mu.Unlock()
	delete(DB.Db, title)
	return nil
}

func CheckTaskExistByTitle(title string) (bool, error) {
	DB.Mu.Lock()
	defer DB.Mu.Unlock()
	if _, ok := DB.Db[title]; !ok {
		return false, nil
	}

	return true, nil
}
