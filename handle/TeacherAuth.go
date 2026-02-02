package api

import (
	"encoding/json"
	db "github/prac-soc/db/store"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Queries *db.Queries
}

func (h *Handler) RegisterTeacher(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var req registerteacherrequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if req.Email == "" || req.Password == "" {
		http.Error(w, "invalid inputs", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "server side err", http.StatusInternalServerError)
		return
	}

	teacher, err := h.Queries.CreateTeacher(ctx, db.CreateTeacherParams{
		Email:        req.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		http.Error(w, "server side err", http.StatusInternalServerError)
		return
	}

	resp := RegisterTeacherResponse{
		Email: teacher.Email,
		ID:    teacher.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
