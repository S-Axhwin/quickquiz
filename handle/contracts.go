package api

import "github.com/google/uuid"

type registerteacherrequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterTeacherResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
