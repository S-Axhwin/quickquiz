-- name: CreateSession :one
INSERT INTO sessions (quiz_id, join_code, status)
VALUES ($1, $2, 'WAITING')
RETURNING *;

-- name: GetSessionByJoinCode :one
SELECT * FROM sessions
WHERE join_code = $1;

-- name: GetSessionByID :one
SELECT * FROM sessions
WHERE id = $1;

-- name: StartSession :one
UPDATE sessions
SET status = 'ACTIVE',
    current_question_index = 0,
    question_started_at = now()
WHERE id = $1
RETURNING *;

-- name: AdvanceQuestion :one
UPDATE sessions
SET current_question_index = current_question_index + 1,
    question_started_at = now()
WHERE id = $1
RETURNING *;

-- name: EndSession :one
UPDATE sessions
SET status = 'ENDED'
WHERE id = $1
RETURNING *;

