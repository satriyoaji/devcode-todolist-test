package dto

type CreateActivityPayload struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:""` //required,email
}
type UpdateActivityPayload struct {
	Title string `json:"title" validate:"required"`
}
