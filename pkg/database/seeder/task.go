package seeder

import (
	"eventzezz_backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func CreateTask(
	db *gorm.DB,
	id uuid.UUID,
	eventID uuid.UUID,
	taskName string,
	isDone bool,
	description string,
	startDate time.Time,
	endDate time.Time,
) error {
	return db.Create(&domain.Task{
		ID:          id,
		EventID:     eventID,
		TaskName:    taskName,
		IsDone:      isDone,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
	}).Error
}

func CreateTaskCategoryAssignment(
	db *gorm.DB,
	id uuid.UUID,
	taskID uuid.UUID,
	taskCategoryID uuid.UUID,
) error {
	return db.Create(&domain.TaskCategoryAssignment{
		ID:             id,
		TaskID:         taskID,
		TaskCategoryID: taskCategoryID,
	}).Error
}

func CreateTaskCategory(
	db *gorm.DB,
	id uuid.UUID,
	name string,
	color string,
) error {
	return db.Create(&domain.TaskCategory{
		ID:    id,
		Name:  name,
		Color: color,
	}).Error
}
