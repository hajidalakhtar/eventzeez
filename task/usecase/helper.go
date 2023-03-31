package usecase

import "eventzezz_backend/domain"

func ToTaskResponses(task []domain.Task, taskCategoryAssignments []domain.TaskCategoryAssignment) []domain.TaskResponse {

	var taskResponses []domain.TaskResponse
	for _, task := range task {
		taskResponse := ToTaskResponse(task, taskCategoryAssignments)
		taskResponses = append(taskResponses, taskResponse)
	}

	return taskResponses

}

func ToTaskResponse(task domain.Task, taskCategoryAssignments []domain.TaskCategoryAssignment) domain.TaskResponse {
	taskResponse := domain.TaskResponse{
		ID:          task.ID,
		TaskName:    task.TaskName,
		Description: task.Description,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	for _, taskCategoryAssignment := range taskCategoryAssignments {
		taskResponse.TaskCategory = append(taskResponse.TaskCategory, taskCategoryAssignment.TaskCategory)
	}

	return taskResponse
}
