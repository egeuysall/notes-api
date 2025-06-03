# 🧠 Go Backend Practice – Day 1

**Goal:** Build a minimal REST API with chi to lock in routing, JSON decoding, and error handling — all from memory.

---

## ✅ Mini Project: Notes API (In-Memory)

| Method | Route          | Description                              |
|--------|----------------|------------------------------------------|
| GET    | `/ping`        | Returns `{ "message": "pong" }`          |
| POST   | `/note`        | Accepts `{ "text": "..." }`, returns ID  |
| GET    | `/note/{id}`   | Returns note by ID or 404                |

- Use: `chi`, `json.NewDecoder`, `json.NewEncoder`, `http.Error`, `chi.URLParam`
- Store data in: `map[string]string`
- Generate random ID (e.g. `6-char` string)

---

## ⏱️ Schedule (3 Hours)

### 🕒 Hour 1: Recall Drill – No Copying
- [x] Write `main.go` with `chi.NewRouter`
- [x] Create `GET /ping` route
- [x] Respond with JSON manually: `json.NewEncoder(w).Encode(...)`
- [x] Set correct headers/status codes

### 🕒 Hour 2: Build Note API
- [x] Add `POST /note`
  - [x] Decode JSON body into struct
  - [x] Generate random ID
  - [x] Store in map
  - [x] Respond with ID in JSON
- [x] Add `GET /note/{id}`
  - [x] Get param with `chi.URLParam`
  - [x] Lookup in map, return value or 404
- [x] Add `DELETE /note/{id}`
  - [x] Get param with `chi.URLParam`
  - [x] Check if note exists
  - [x] Delete from map
  - [x] Return 204 No Content

### 🕒 Hour 3: Review & Refactor
- [ ] Build the Notes CLI
  - [ ] Flags package
  - [ ] error handling
- [ ] Publish via Goreleaser
  - [ ] Configure Goreleaser

### 🕒 Hour 3: Review & Refactor
- [x] Clean up code structure (handlers, error messages)
- [x] Write a reflection note in Obsidian:
  - What did I remember easily?
  - What did I forget?
  - Any confusing areas?
- [x] Save code to `~/go-snippets/notes-api`
