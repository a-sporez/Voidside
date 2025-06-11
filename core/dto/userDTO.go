// /core/dto/userDTO.go
package dto

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}
