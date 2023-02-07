package dto

type CreateTodoPayload struct {
	Title           string `json:"title" validate:"required" binding:"required"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required,numeric" binding:"required,numeric"`
	IsActive        *bool  `json:"is_active" validate:"required,boolean" binding:"required,boolean"`
}
type UpdateTodoPayload struct {
	Title    string `json:"title" validate:"required" binding:"required"`
	Priority string `json:"priority" validate:"required" binding:"required"`
	IsActive *bool  `json:"is_active" validate:"required,boolean" binding:"required,boolean"`
}
