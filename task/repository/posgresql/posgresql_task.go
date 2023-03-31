package posgresql

import (
	"context"
	"eventzezz_backend/domain"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type posgresqlTaskRepository struct {
	conn *gorm.DB
}

func NewPosgresqlTaskRepository(conn *gorm.DB) domain.TaskRepository {
	return &posgresqlTaskRepository{conn}
}

func (p posgresqlTaskRepository) GetTasks(ctx context.Context, cursor string, num int64) ([]domain.Task, string, error) {
	var tasks []domain.Task
	result := p.conn.Find(&tasks)
	return tasks, "", result.Error
}

func (p posgresqlTaskRepository) GetTaskByID(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	var task domain.Task
	result := p.conn.Joins("Event").First(&task, id)
	return task, result.Error
}

func (p posgresqlTaskRepository) GetTaskCategoryAssignmentsByTaskID(ctx context.Context, id uuid.UUID) ([]domain.TaskCategoryAssignment, error) {
	var taskCategoryAssignments []domain.TaskCategoryAssignment
	result := p.conn.Joins("TaskCategory").Joins("Task").Where("task_id = ?", id).Find(&taskCategoryAssignments)
	fmt.Println(taskCategoryAssignments)
	return taskCategoryAssignments, result.Error
}

func (p posgresqlTaskRepository) UpdateTask(ctx context.Context, u *domain.Task, id uuid.UUID) error {
	result := p.conn.Model(&domain.Task{}).Where("id = ?", id).Updates(u)
	return result.Error
}

func (p posgresqlTaskRepository) CreateTask(ctx context.Context, task *domain.Task) error {
	fmt.Println(task.EventID)
	result := p.conn.Create(task)
	return result.Error
}

func (p posgresqlTaskRepository) DeleteTaskByID(ctx context.Context, id uuid.UUID) error {
	result := p.conn.Delete(&domain.Task{}, id)
	return result.Error
}
