package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tasks_api/pkg/models"
	"github.com/tasks_api/pkg/storage"
	"github.com/tasks_api/pkg/tasks"
)

// postTask godoc
// @Summary Создать новую задачу
// @Description Создает новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.CreateTaskRequest true "Данные задачи"
// @Success 200 {object} map[string]interface{} "Задача успешно создана"
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks [post]
func postTask(c *fiber.Ctx) error {

	var newTaskRequest models.CreateTaskRequest

	// парсим json в структуру newTask , ошибку вернем так же в формате JSON
	if err := c.BodyParser(&newTaskRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	// проверим указан ли title
	if newTaskRequest.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title не может быть пустым"})
	}

	// проверим, существует ли уже задача с таким title
	exist, err := storage.CheckTaskExistByTitle(newTaskRequest.Title)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if exist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "задача с таким title уже имеется"})
	}

	// формируем основную структуру Task для добавления в БД
	newTask := &models.Task{
		Title:       newTaskRequest.Title,
		Description: newTaskRequest.Description,
		Created_at:  time.Now(),
		Status:      models.StatusInProgress,
	}

	// добавляем в БД
	if err := storage.CreateTask(newTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// запускаем задачу
	go tasks.ProcessingTask(newTask)

	// успешный ответ
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Задача успешно добавлена",
	})

}

// getTask godoc
// @Summary Получить задачу
// @Description Получает задачу по title
// @Tags tasks
// @Accept json
// @Produce json
// @Param title path string true "title задачи"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks/{title} [get]
func getTask(c *fiber.Ctx) error {
	// провалидируем title
	title := c.Params("title")
	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный title"})
	}

	//проверим,существует ли задача с таким title
	exist, err := storage.CheckTaskExistByTitle(title)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if !exist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "задачи с таким title не содержится"})
	}

	// запрос к storage
	task, err := storage.GetTaskByTitle(title)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	tasks.SetDuration(task)

	return c.JSON(task)
}

// getAllTasks godoc
// @Summary Получить все задачи
// @Description Получает все имеющиеся задачи
// @Tags tasks
// @Accept json
// @Produce json
// Success 200 {array} models.Task
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks [get]
func getAllTasks(c *fiber.Ctx) error {
	allTasks := []*models.Task{}

	// получим список задач
	if err := storage.GetAllTasksFromStorage(&allTasks); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for _, task := range allTasks {
		tasks.SetDuration(task)
	}
	return c.JSON(allTasks)
}

// deleteTask godoc
// @Summary Удалить задачу
// @Description Удаляет задачу по Title
// @Tags tasks
// @Accept json
// @Produce json
// @Param title path string true "title задачи"
// @Success 200 {object} map[string]interface{} "Задача успешно удалена"
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks/{title} [delete]
func deleteTask(c *fiber.Ctx) error {
	// провалидируем title
	title := c.Params("title")
	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный title"})
	}

	//проверим,существует ли задача с таким title
	exist, err := storage.CheckTaskExistByTitle(title)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if !exist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "задачи с таким title не содержится"})
	}

	// удаление
	if err := storage.DeleteTaskFromStorage(title); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// успешный ответ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Задача успешно удалена",
	})
}
