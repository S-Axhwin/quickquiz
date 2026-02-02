CREATE INDEX idx_answers_session_id ON answers(session_id);
CREATE INDEX idx_answers_student_id ON answers(student_id);
CREATE INDEX idx_answers_question_id ON answers(question_id);

CREATE INDEX idx_students_session_id ON students(session_id);

CREATE INDEX idx_sessions_quiz_id ON sessions(quiz_id);
CREATE INDEX idx_sessions_join_code ON sessions(join_code);

