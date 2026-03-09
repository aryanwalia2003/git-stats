# 📂 Project Folder Structure Guide

Yeh ek **Lazygit-style TUI** hai jo kisi bhi git repo ke andar chala kar uski stats dikhata hai.
Humare architecture ka rule hai: **Har file 30 lines se kam**, har folder ki **ek specific responsibility**.

---

## File Naming Convention (Sabse Zaroori!)

Har file ka naam **3 cheezein** batata hai: **Kis struct ka hai** → **Kya topic hai** → **Kis type ka code hai**

| Suffix | Matlab | Example |
|---|---|---|
| `_struct.go` | Struct definition + constructor | `reader_struct.go` |
| `_method.go` | Kisi struct ka method | `reader_repo_method.go` |
| `_iface.go` | Interface definition | `git_reader_iface.go` |
| `_entity.go` | Pure data struct (no logic) | `repo_entity.go` |
| `_helper.go` | Shared utility function | `run_helper.go` |
| `_init.go` | Initialization / setup code | `db_init.go` |
| `_schema.go` | SQL table definitions | `migrations_schema.go` |
| `_style.go` | LipGloss styling constants | `colors_style.go` |
| `parse_*.go` | Stateless parser function | `parse_origin_url.go` |

---

## Folder Breakdown

### 1. `cmd/gh-stats/`
> App ka main entrypoint. Sirf wiring: DB connect karo, Reader banao, Bubble Tea start karo. **Zero business logic.**

```
cmd/gh-stats/
└── main.go
```

### 2. `internal/domain/`
> **Source of Truth.** Sirf Interfaces aur Data Structures. Koi external dependency nahi (no DB, no UI, no API).
> Baaki saare packages (db, git, api) in interfaces ko "satisfy" karte hain.

```
internal/domain/
├── repo_entity.go            ← Repo struct (LocalPath, Owner, Name, Branch...)
├── stat_entity.go            ← Stat struct (Label, Value, Date...)
├── git_reader_iface.go       ← LocalGitReader interface (GetCurrentRepo, GetRecentCommits...)
└── github_reader_iface.go    ← GitHubReader interface (GetTraffic, GetStargazers...)
```

**Example** — `repo_entity.go`:
```go
package domain

type Repo struct {
    LocalPath    string
    RemoteOrigin string
    Name         string
    Owner        string
    Branch       string
}
```

### 3. `internal/git/`
> Local `.git` folder se data parse karta hai. `domain.LocalGitReader` interface implement karta hai.
> `reader_*` files = Reader struct ke methods. `parse_*` files = standalone parser functions.

```
internal/git/
├── reader_struct.go                ← Reader struct + NewReader() constructor
├── reader_repo_method.go           ← GetCurrentRepo()  — remote origin & branch padhta hai
├── reader_commits_method.go        ← GetRecentCommits() — git log parse karta hai
├── reader_contributors_method.go   ← GetLocalContributors() — git shortlog se top authors
├── reader_churn_method.go          ← GetCodeChurn() — lines added/deleted per commit
├── run_helper.go                   ← runGit() — shared helper jo git commands execute karta hai
├── parse_origin_url.go             ← remote URL se owner/name nikalta hai
├── parse_contributors.go           ← shortlog output ko Stat[] mein convert karta hai
├── parse_commits.go                ← git log output ko Stat[] mein convert karta hai
└── parse_churn.go                  ← shortstat output se insertions/deletions nikalta hai
```

### 4. `internal/db/`
> SQLite database layer. GitHub API ka cached data store karta hai (Stars, Forks snapshots) aur AI insights cache karta hai.
> Local git data yahan NAHI aata — woh seedha `.git` se on-the-fly parse hota hai.

```
internal/db/
├── db_init.go              ← SQLite connection setup + migrations call
└── migrations_schema.go    ← SQL CREATE TABLE statements (repo_snapshots, ai_insights)
```

**Tables:**
- `repo_snapshots` — GitHub API se aaye Stars/Forks/Issues ka daily snapshot
- `ai_insights` — LLM-generated timeline events ka cache ("The Major Refactor")

### 5. `internal/api/`
> GitHub API se remote stats fetch karta hai (Stars, Traffic, Issues). `domain.GitHubReader` implement karega.
> Har request `tea.Cmd` return karegi taaki UI freeze na ho.

```
internal/api/
├── client_struct.go        ← (future) HTTP client + auth setup
├── client_traffic_method.go ← (future) GetTraffic()
└── client_stars_method.go  ← (future) GetStargazers()
```

### 6. `internal/ui/`
> Bubble Tea TUI components. Har UI section apna independent sub-model hai.

```
internal/ui/
├── app/                    ← Root orchestrator (delegates to sub-models)
│   ├── model.go
│   ├── update.go
│   └── view.go
├── header/                 ← (future) Header bar component
├── sidebar/                ← (future) Navigation panel
├── stats/                  ← (future) Main stats dashboard view
└── theme/
    └── colors_style.go     ← LipGloss colors, fonts, layouts (CSS jaisa)
```

**Rule:** `app/` mein koi business logic nahi — sirf sub-models ko messages forward karna aur layout banana.

---

## ⚡ Golden Rules

1. **30 Lines Max:** File 30 lines cross kare → naya helper function alag file mein nikalo
2. **Struct naam file mein:** `reader_repo_method.go` = Reader struct ka Repo method
3. **Narrative Naming:** `githubApiResponse` likho, `res` mat likho
4. **Guard Clauses:** `if err != nil { return }` jaldi likho, nested else avoid karo
5. **Absolute Imports:** `github.com/aryanwalia2003/git-stats/internal/domain` use karo, `../../` kabhi nahi
