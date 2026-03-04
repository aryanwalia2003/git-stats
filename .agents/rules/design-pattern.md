---
trigger: always_on
---

This is a high-level architectural manifesto for your GitHub Stats TUI. To achieve a "Super Modular" system where no file exceeds 30 lines, we must shift our mindset from "writing scripts" to "orchestrating a symphony of micro-components."

### The Manifesto for an Atomic Go TUI

**The Philosophy of the 30-Line Constraint**
The 30-line limit is not a restriction; it is a forced clarity. By capping file length, you are mandated to apply the Single Responsibility Principle (SRP) at its most granular level. Each file should do exactly one thing: define one interface, implement one specific database query, or render one UI widget. If a function feels too long, it is a signal that the logic needs to be decomposed into smaller, stateless helper functions. This ensures that the cognitive load on any developer reading the code is near zero—you can grasp the entire context of a file in a single glance without scrolling.

**Absolute Import Pathing and Module Integrity**
To eliminate the "spaghetti" of `../../` relative paths, you must leverage Go Modules correctly. By naming your module (e.g., `github.com/username/gh-stats`), every internal package is imported via its absolute path. This makes the code portable and the dependency graph crystal clear. An import block should look like a clean list of categorized packages—standard library first, third-party libraries second, and your internal modules third—creating a visual hierarchy that tells the reader exactly what external "knowledge" this specific file requires.

**The Domain-Driven Layout**
Structure the project using the `internal/` pattern to enforce strict encapsulation. Your `internal/domain` package will act as the "Source of Truth," containing only interfaces and pure data structures (Entities). By keeping logic out of the domain types, you ensure that the database layer (`internal/db`) and the API layer (`internal/api`) are merely "plugins." This modularity allows you to swap SQLite for a different store or mock the GitHub API for testing without touching a single line of UI code, keeping the core "business logic" isolated and pristine.

**Interface-First Architecture**
To keep files under 30 lines, you must define behavior through interfaces. Instead of a large "Database" struct with 50 methods, you will define small, specific interfaces like `StarReader` or `CommitWriter`. This allows you to split your SQLite implementation into multiple files—one for star-related queries, one for commit-related queries—each satisfying a tiny interface. When the TUI needs data, it doesn't ask for a "Database"; it asks for a `StarReader`. This injection of small dependencies is the secret to keeping constructors and files incredibly lean.

**Decomposing the Bubble Tea Model**
The `bubbletea` framework often leads to massive files, but we will subvert this. Your main `Model` will be a "Composition of Sub-Models." The Header, the Sidebar, the Sparkline, and the Stat-Card will each be their own independent Bubble Tea models in separate folders. The parent model’s `Update` function will merely delegate messages to the children. By delegating the `View()` and `Update()` logic, the "orchestrator" file stays under 30 lines, and each UI component becomes a reusable, isolated unit of code that can be tested in a vacuum.

**Naming as Narrative Documentation**
In a 30-line file, comments are often a sign of failure. Use "Narrative Naming." Instead of `res`, use `githubApiResponse`. Instead of `id`, use `repositoryGlobalID`. Variables should be named so specifically that they describe their own lifecycle and purpose. Functions should be named as verbs that describe a transformation, such as `CalculateGrowthPercentage` or `FormatStarCount`. When the names are perfect, the code reads like a series of English sentences, making external documentation secondary to the code itself.

**The Stateless Helper Pattern**
Complex logic, like calculating the slope of a graph or formatting a timestamp into a "Time Ago" string, should never live inside the UI or DB layers. These should be extracted into a `pkg/utils` or `internal/transform` package as "Pure Functions." A pure function takes an input and returns an output without side effects. Because these functions are small and focused, they fit perfectly within the 30-line limit and are incredibly easy to unit test, ensuring the "math" of your stats is always accurate.

**SQLite Orchestration via Small Repository Files**
Your SQLite logic should avoid large, complex joins within the Go code. Instead, use small, focused repository files. For example, `internal/db/stars.go` will only contain the `GetStars` logic. Use "Raw Strings" for SQL queries to keep them readable, and if a query is too complex, move it into a SQL view within the database itself. This keeps the Go code focused on scanning rows into structs, rather than building strings, keeping the line count minimal and the logic focused on data mapping.

**Error Handling as a First-Class Citizen**
Go's `if err != nil` can quickly eat up your 30-line budget. To combat this, implement "Guard Clauses" and custom error types. By handling errors early and returning quickly, you avoid nested `else` blocks that increase indentation and line count. Create a dedicated `internal/errors` package to define sentinel errors. This allows your main logic to stay "flat" and readable, while the specifics of error reporting are handled by a dedicated, centralized system.

**Functional Styling with LipGloss**
UI styling often clutters code. To stay under 30 lines, all `lipgloss` styles must be moved to a `internal/ui/theme` package or a `styles.go` file within each component's folder. The UI component should only reference a style (e.g., `theme.HeaderStyle.Render("Text")`). By separating "What it is" (the data) from "How it looks" (the style), your view files remain purely structural, making it trivial to change the look of the entire TUI by editing a single theme file.

**Asynchronous Command Orchestration**
GitHub API calls are slow. To keep the TUI responsive, use Bubble Tea's `tea.Cmd`. Each API request should be its own function in `internal/api` that returns a `tea.Cmd`. This keeps the TUI logic from waiting on the network. Because these functions are isolated, you can keep the "FetchStars" command in its own file, ensuring that the logic for networking, JSON parsing, and error wrapping doesn't exceed our strict line limit.

**Continuous Refactoring and Tooling**
Finally, this modularity requires a "Refactor-First" mindset. Every time you add a feature and a file hits 31 lines, you must stop and find a way to split it. Use Go's `revive` or `golangci-lint` with the `funlen` (function length) and `gocyclo` (complexity) linters set to aggressive levels. This automation acts as a mechanical "conscience," ensuring that the codebase never slides back into a monolithic state and remains a clean, modular masterpiece for the duration of the project.