package dto

type CreateTodoPayload struct {
	Title           string `json:"title" validate:"required" binding:"required"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required,numeric" binding:"required,numeric"`
	IsActive        *bool  `json:"is_active" validate:"" binding:""`
}
type UpdateTodoPayload struct {
	Title    string `json:"title" validate:"" binding:""`
	Priority string `json:"priority" validate:"" binding:""`
	IsActive *bool  `json:"is_active" validate:"" binding:""`
}
