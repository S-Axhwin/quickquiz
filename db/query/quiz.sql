-- name: CreateQuiz :one
INSERT INTO quizzes (teacher_id, title)
VALUES ($1, $2)
RETURNING *;

-- name: GetQuizByID :one
SELECT * FROM quizzes
WHERE id = $1;

-- name: CreateQuestion :one
INSERT INTO questions (
    quiz_id,
    text,
    options,
    correct_option,
    time_limit_seconds,
    order_index
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetQuestionsByQuiz :many
SELECT * FROM questions
WHERE quiz_id = $1
ORDER BY order_index ASC;

