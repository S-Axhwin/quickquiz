-- name: CreateTeacher :one
INSERT INTO teachers (email, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: GetTeacherByEmail :one
SELECT * FROM teachers
WHERE email = $1;

