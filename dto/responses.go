package dto

type CreateTaskResponse struct {
	Task *Task `json:"task"`
}

type GetTasksResponse struct {
	Task []Task `json:"tasks"`
}

type GetTaskByIdResponse struct {
	Task Task `json:"task"`
}

type UpdateTaskByIdResponse struct {
	Task Task `json:"task"`
}
