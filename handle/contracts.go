package api

type RegisterTeacherRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterTeacherResponse struct {
	ID    int32  `json:"id"`
	Email string `json:"email"`
}
