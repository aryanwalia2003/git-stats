# Project Folder Structure Guide (GitHub Stats TUI)

Yeh guide humare 30-line constraint manifesto par based hai. Harr ek folder ki ek specific responsibility (Single Responsibility Principle) hai. Yahan detail mein diya gaya hai ki kis folder mein kya jayega.

## Folder Breakdown

### 1. `cmd/`
**Kya jayega:** Application ka main entrypoint.
**Example file:** `cmd/gh-stats/main.go`
**Rule:** Isme sirff basic setup, configuration load aur Bubble Tea program initialization hona chahiye. Koi business logic ya UI drawing yahan nahi hogi.

### 2. `internal/domain/`
**Kya jayega:** **Source of Truth**. Yahan sirff Interfaces aur pure Data Structures (Entities) aayenge. Is package mein koi external dependency nahi honi chahiye (No DB, no UI).
**Example file:** `internal/domain/user.go`
**Dikhne mein kaisa hoga:**
```go
package domain

// Entity
type User struct {
    ID    string
    Login string
}

// Interface
type UserRepository interface {
    GetUser(id string) (*User, error)
}
```

### 3. `internal/db/`
**Kya jayega:** SQLite queries aur database interactions jo `domain` package mein bani interfaces ko satisfy karti hain.
**Example file:** `internal/db/stars.go`
**Rule:** Complex ORM avoid karo, raw SQL strings use karo. Agar join bahut bada ho raha hai, toh SQL Views bana lo DB mein thake Go files clean rahein. File length 30 lines ke andar rakhne ke liye chhote repository files banao.

### 4. `internal/api/`
**Kya jayega:** External HTTP requests aur GitHub API se baat karne wala code. Ye functions data fetch karke `tea.Cmd` return karenge taaki network requests completely asynchronous rahen aur UI freeze na ho.
**Example file:** `internal/api/fetch_commits.go`

### 5. `internal/ui/`
**Kya jayega:** TUI Components aur Bubble Tea model layer. UI ek large monolith model banne ki jagah "Compositions of Sub-Models" hona chahiye. 
**Sub-folders:**
- `internal/ui/header/`: Header UI ka apna model, update, view structure.
- `internal/ui/sidebar/`: Sidebar logic ke liye independent files.
**Example file:** `internal/ui/sidebar/view.go`
**Rule:** Main application loop sirf in chhote sub-models ko messages delegate (pass) karega.
- `internal/ui/app/`: Root application shell jo saare sub-models (Header, Sidebar, Views) ko coordinate aur compose karta hai.
**Example file:** `internal/ui/app/model.go`
**Rule:** Isme business logic nahi hogi, sirf sub-models ko messages pass karna aur layout build karna iska kaam hai.```

### 6. `internal/ui/theme/`
**Kya jayega:** Sirf Design, Layout aur Colors! Saari LipGloss styling yahan hogi taaki `internal/ui` elements clean rahen aur HTML ki tarah CSS styles alag rahen.
**Example file:** `internal/ui/theme/buttons.go` aur `internal/ui/theme/colors.go`

### 7. `internal/transform/` ya `pkg/utils/`
**Kya jayega:** "Stateless Helper Pattern" i.e., Pure functions jinme koi side effect nahi hota.
Time ago strings banane ka logic, Data formatting, numbers ko format karna aadi idhar rahenge kyuki inme error ya logic test karna aasaan rehta hai.
**Example file:** `pkg/utils/timeline.go` 
**Dikhne mein kaisa hoga:** `func CalculateGrowthPercentage(oldVal, newVal int) float64`

### 8. `internal/errors/`
**Kya jayega:** Custom error types aur sentinel errors taaki app code mein har jagah flat error-handling ho (Guard clauses).
**Example file:** `internal/errors/sentinel.go`

---

### Golden Rules (File likhte waqt):
1. **File Limit (30 Lines):** Try your best to keep files under 30 lines. Agar code 30 lines cross kar raha hai, toh ek naya helper function nikalo doosri file mein.
2. **Narrative Naming:** Variables ka naam clearly poora rakhna hai (e.g., `githubApiResponse` not `res`, `repositoryGlobalID` not `id`), taaki code hi documentation ban jaaye.
3. **Guard Clauses:** `if err != nil` se error jaldi return karao aur nested `else` statements avoid karo.
4. **Absolute Imports:** Import ke liye hamesha apne package ka absolute path use karo (e.g. `github.com/aryanwalia/gh-stats/internal/domain`), relative `../../` paths allow nahi karne hai.
