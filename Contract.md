# Realtime Quiz Platform – Backend ↔ Frontend Contract

This document defines the **exact API and WebSocket contracts** between backend (Go) and frontend (HTML/React).

Principles:
- HTTP = commands (create, join, submit)
- WebSocket = events (realtime updates)
- Backend is the source of truth
- Frontend reacts, never decides quiz state

---

## Authentication (Teacher only)

### POST /auth/login
Authenticate teacher and return JWT.

**Request**
```json
{
  "email": "teacher@test.com",
  "password": "plaintext"
}
````

**Response**

```json
{
  "access_token": "jwt",
  "teacher_id": "uuid"
}
```

Frontend stores token and sends it in `Authorization` header.

---

## Quiz Management (Teacher)

### POST /quizzes

Create a new quiz.

**Headers**

```
Authorization: Bearer <jwt>
```

**Request**

```json
{
  "title": "DBMS Quiz"
}
```

**Response**

```json
{
  "quiz_id": "uuid",
  "title": "DBMS Quiz"
}
```

---

### POST /quizzes/{quizId}/questions

Add a question to a quiz.

**Request**

```json
{
  "text": "What is ACID?",
  "options": ["A", "B", "C", "D"],
  "correct_option": 1,
  "time_limit_seconds": 20,
  "order_index": 1
}
```

**Response**

```json
{
  "question_id": "uuid"
}
```

---

### GET /quizzes/{quizId}/questions

Fetch quiz questions (teacher preview).

**Response**

```json
[
  {
    "id": "uuid",
    "text": "What is ACID?",
    "options": ["A", "B", "C", "D"],
    "time_limit_seconds": 20,
    "order_index": 1
  }
]
```

---

## Session Lifecycle

### POST /sessions

Create a live quiz session.

**Headers**

```
Authorization: Bearer <jwt>
```

**Request**

```json
{
  "quiz_id": "uuid"
}
```

**Response**

```json
{
  "session_id": "uuid",
  "join_code": "ABCD12",
  "status": "WAITING"
}
```

---

### POST /sessions/join

Student joins a session using join code.

**Request**

```json
{
  "join_code": "ABCD12",
  "name": "StudentName"
}
```

**Response**

```json
{
  "student_id": "uuid",
  "session_id": "uuid"
}
```

No authentication for students.

---

### POST /sessions/{sessionId}/start

Teacher starts the quiz.

**Headers**

```
Authorization: Bearer <jwt>
```

**Response**

```json
{
  "status": "ACTIVE",
  "current_question_index": 0
}
```

Triggers WebSocket broadcast.

---

## Questions During Session

### GET /sessions/{sessionId}/current-question

Fetch current question (used on refresh / reconnect).

**Response**

```json
{
  "question_id": "uuid",
  "text": "What is ACID?",
  "options": ["A", "B", "C", "D"],
  "time_limit_seconds": 20,
  "started_at": "2026-02-02T10:00:00Z"
}
```

Frontend computes remaining time as:

```
remaining = time_limit_seconds - (now - started_at)
```

---

## Answers (Student)

### POST /answers

Submit or update an answer.

**Request**

```json
{
  "session_id": "uuid",
  "student_id": "uuid",
  "question_id": "uuid",
  "selected_option": 2
}
```

**Response**

```json
{
  "status": "saved"
}
```

Answers are idempotent until timer ends.

---

## WebSocket (Realtime Events)

### GET /ws/sessions/{sessionId}

Query params:

* Teacher: `?role=teacher`
* Student: `?role=student&student_id=uuid`

This is the **only WebSocket endpoint**.

Frontend does NOT send quiz logic over WS.

---

## WebSocket Event Types (Server → Client)

### SESSION_STARTED

```json
{
  "type": "SESSION_STARTED"
}
```

Frontend:

* Fetch `/current-question`
* Initialize UI

---

### QUESTION_STARTED

```json
{
  "type": "QUESTION_STARTED",
  "question_index": 1,
  "started_at": "2026-02-02T10:01:00Z"
}
```

Frontend:

* Fetch current question
* Reset timer

---

### QUESTION_ENDED

```json
{
  "type": "QUESTION_ENDED"
}
```

Frontend:

* Disable answering
* Show waiting state

---

### SESSION_ENDED

```json
{
  "type": "SESSION_ENDED"
}
```

Frontend:

* Navigate to results / summary page

---

## Timer & State Rules (Critical)

* Backend controls question timing
* Backend stores `question_started_at`
* Backend advances questions via goroutines/tickers
* Frontend NEVER decides when a question ends
* Frontend uses server timestamps only

---

## Summary

* HTTP = actions
* WebSocket = notifications
* DB = source of truth
* Frontend = reactive UI

Following this contract guarantees:

* No desync
* No cheating
* Clean separation of concerns
* Scalable realtime behavior

