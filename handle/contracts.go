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

type CreateRoomRequest struct {
	Title string `json:"title"`
}

type CreateRoomResponse struct {
	Quiz_id uuid.UUID `json:"quiz_id"`
	Title   string    `json:"title"`
}
