# 🚀 GitHub Stats TUI: Lazygit-Style Repo Analyzer 📊

This tool acts like `lazygit`. You `cd` into a directory, type `gh-stats`, and instantly get a hacker-style dashboard for **that specific repository**.

We achieve this by combining **Local Git History** (fast, deep analysis of your local `.git` folder) with **Remote GitHub API Data** (Stars, Traffic, Issues).

Here are the cool, informative, and playful stats we can show for a singular repo:

## 📈 The Dashboard (Remote + Local Info)
*   **Repo Vitality:** Current Stargazers, Forks, and total Open Issues (Fetched from GitHub API).
*   **The Traffic Sparkline:** A graph showing clones or views over the last 14 days (Fetched from GitHub API).
*   **The Top Contributors:** A leaderboard of who actually wrote the most code in *this* repo (Calculated from local `git log`).
*   **Issue/PR Triage Pane:** A list of the most recent open Pull Requests or Issues so you know what needs attention *right now*.

## 🦉 Local Developer Vibes (Behavioral Stats for THIS Repo)
*   **Night Owl Index 🌙:** Percentage of commits made between *Midnight and 6 AM* by you in this specific codebase.
*   **Busiest Day & Time 🗓️:** "The team is most productive on **Tuesdays at 2:00 PM**."
*   **The Code Churner 🌪️:** Which specific *files* or *folders* in this repo are being modified the most? (Helps find complex/buggy areas).
*   **The Talker 🗯️:** Average length of commit messages in this project. Is your team writing essays or just `fixed bug`?
*   **Emoji King/Queen 👑:** Most frequently used emoji in the commit messages of this repo.

## 🎭 Repo-Specific Trophies (Fun Awards)
*   **The Monolith File 🏛️:** The absolute largest file in the current repository by line count or size.
*   **The Ancient Relic 🏺:** The file that hasn't been touched in the longest time.
*   **Caffeine Driven Development ☕:** Spotting sudden, massive spikes of commits from a single author within a 1-hour window.
*   **The Bug Squasher 🪲:** The contributor who closes or mentions the most issues in their commits.

## 🎨 Visualization Ideas for the TUI
*   **Activity Heatmap:** A terminal-based git contribution graph purely for this repository (similar to the GitHub profile graph, but specific to this project).
*   **Codebase Breakdown:** A colorful `lipgloss` progress bar showing the percentage of languages used (e.g., 60% Go, 40% Markdown).
*   **Traffic History Chart:** An ASCII line chart using a library like `asciigraph` to show repository views over the last two weeks.

---
*Architecture Note: We use `LocalGitReader` to parse local history for the "Vibes" and `GitHubReader` to fetch external stats for the "Dashboard", blending them into a single fast TUI.*
