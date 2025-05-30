package task

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// GetTasksByUserID - Возвращаем массив задач пользователя по его ID
	GetTasksByUserID(userID uint) ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// CreateTask создает новую задачу в базе данных
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

// GetAllTasks возвращает все задачи из базы данных
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	result := r.db.Model(&Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	var updatedTask Task
	err := r.db.First(&updatedTask, id).Error
	if err != nil {
		return Task{}, err
	}
	return updatedTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
