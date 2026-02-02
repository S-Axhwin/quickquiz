-- name: SubmitAnswer :one
INSERT INTO answers (
    session_id,
    student_id,
    question_id,
    selected_option
) VALUES ($1, $2, $3, $4)
ON CONFLICT (student_id, question_id)
DO UPDATE SET
    selected_option = EXCLUDED.selected_option,
    answered_at = now()
RETURNING *;

-- name: GetAnswersByQuestion :many
SELECT * FROM answers
WHERE session_id = $1 AND question_id = $2;

