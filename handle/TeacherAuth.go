package api

import (
	db "github/prac-soc/db/store"
	"net/http"
)

type Handler struct {
	Queries *db.Queries
}

func (h *Handler) RegisterTeacher(w http.ResponseWriter, r *http.Request) {

}
