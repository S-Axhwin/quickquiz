package api

import (
	"encoding/json"
	"fmt"
	db "github/prac-soc/db/store"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Queries *db.Queries
}

// TODO: wanna to comp this
func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Came inside create room")
	ctx := r.Context()
	var req CreateRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "required title", http.StatusBadRequest)
		return
	}

	room, err := h.Queries.CreateQuiz(ctx, db.CreateQuizParams{
		Title: req.Title,
	})
	if err != nil {
		http.Error(w, "Error while creating room. may be already room exisits", http.StatusInternalServerError)
		return
	}

	resp := CreateRoomResponse{
		Quiz_id: room.ID,
		Title:   room.Title,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
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
		http.Error(w, "User Alreay Exisits", http.StatusBadRequest)
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

func (h *Handler) LoginTeacher(w http.ResponseWriter, r *http.Request) {

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

	teacher, err := h.Queries.GetTeacherByEmail(ctx, req.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(teacher.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	claims := jwt.RegisteredClaims{
		Subject:   teacher.ID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "host.com",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenRes, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "token missing", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenRes,
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   false, //TODO: true in prod (HTTPS)
		Path:     "/",
	})

	w.Write([]byte("Login successful"))

}
