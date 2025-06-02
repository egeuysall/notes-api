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
- [ ] Write `main.go` with `chi.NewRouter`
- [ ] Create `GET /ping` route
- [ ] Respond with JSON manually: `json.NewEncoder(w).Encode(...)`
- [ ] Set correct headers/status codes

### 🕒 Hour 2: Build Note API
- [ ] Add `POST /note`
  - [ ] Decode JSON body into struct
  - [ ] Generate random ID
  - [ ] Store in map
  - [ ] Respond with ID in JSON
- [ ] Add `GET /note/{id}`
  - [ ] Get param with `chi.URLParam`
  - [ ] Lookup in map, return value or 404

### 🕒 Hour 3: Review & Refactor
- [ ] Clean up code structure (handlers, error messages)
- [ ] Write a reflection note in Obsidian:
  - What did I remember easily?
  - What did I forget?
  - Any confusing areas?
- [ ] Save code to `~/go-snippets/notes-api`
