package tasks

import (
	"math/rand/v2"
	"time"

	"github.com/tasks_api/pkg/models"
	"github.com/tasks_api/pkg/storage"
)

func SetDuration(task *models.Task) {
	var dur int

	if task.Completed_at == nil {
		dur = int(time.Duration(time.Since(task.Created_at).Minutes()))
	} else {
		dur = int(time.Duration(task.Completed_at.Sub(task.Created_at).Minutes()))
	}

	task.Duration = &dur
}

// имитация I/O bound задачи
func ProcessingTask(task *models.Task) {
	// выберем рандомно сколько будет длиться фиктивная задача
	duration := rand.IntN(3) // от 0 до 2 минут для примера
	time.Sleep(time.Minute * time.Duration(duration))

	// выберем рандомно число для имитации фейла задачи
	randFail := rand.IntN(8)

	storage.DB.Mu.Lock()
	defer storage.DB.Mu.Unlock()
	// имитация фейла задачи
	if randFail == 6 {
		storage.DB.Db[task.Title].Status = models.StatusFailed
	} else {
		storage.DB.Db[task.Title].Status = models.StatusCompleted
		result := "Some result"
		storage.DB.Db[task.Title].Result = &result
	}

	timeCompleted := time.Now()
	storage.DB.Db[task.Title].Completed_at = &timeCompleted

}
