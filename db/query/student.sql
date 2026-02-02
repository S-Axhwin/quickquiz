-- name: CreateStudent :one
INSERT INTO students (session_id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetStudentsBySession :many
SELECT * FROM students
WHERE session_id = $1
ORDER BY joined_at ASC;

