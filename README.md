🧩 Task Tracker CLI (Go)

A fast, lightweight Command-Line Task Tracker built with Golang, designed to demonstrate clean CLI architecture, file persistence, and idiomatic Go patterns.

This project supports full CRUD operations, task status management, and filtering — all backed by a local JSON store.

🔗 Project Repository: https://github.com/AboloreDev/Go-Task-Cli-Tracker

🚀 Features

✅ Add new tasks

✏️ Update existing tasks

❌ Delete tasks

🔄 Mark tasks as todo, in-progress, or done

📋 List tasks with filters

💾 Persistent storage using JSON

🧠 Clean separation of concerns (handlers, storage, helpers)

🛠️ Tech Stack

Language: Go (Golang)

Storage: JSON file persistence

Architecture: Modular packages

CLI Parsing: os.Args

Time Handling: time.Time

Error Handling: Idiomatic Go patterns

📁 Project Structure
task-tracker/
│
├── main.go                 # Entry point & command router
├── go.mod                  # Go module definition
├── tasks.json              # Task persistence (auto-generated)
│
├── helpers/               # CLI command handlers
│   ├── helpers.go
│ 
│
├── storage/                # File I/O & persistence
│   └── storage.go
│
└── progress-trackers/                # Utility helpers

⚙️ Installation & Setup
Prerequisites

Go 1.20+

macOS / Linux / Windows

Clone the repository
git clone https://github.com/your-username/task-tracker.git
cd task-tracker

Run the application
go run main.go <command>

📌 Usage
Add a task
task-tracker add "Finish Go project"

Update a task
task-tracker update 1 "Finish Go CLI project"

Delete a task
task-tracker delete 1

Mark task as in progress
task-tracker mark-in-progress 2

Mark task as done
task-tracker mark-done 2

List tasks
task-tracker list
task-tracker list done
task-tracker list todo
task-tracker list in-progress

🧠 Design Decisions

Switch-based command routing for clarity and extensibility

Index-based slice mutation to avoid common Go pitfalls

Defensive error handling for CLI safety

Human-readable JSON using json.MarshalIndent

No external dependencies — pure Go standard library

⚠️ Notes on File Permissions

The application automatically creates tasks.json on first use.

⚠️ Do not run the application with sudo
This can cause file ownership issues.

🧪 Future Improvements

🔒 File locking for concurrency safety

📦 Package as a standalone binary

🧪 Unit tests

🎨 Colored CLI output

🏠 Store tasks in $HOME/.task-tracker/

👤 Author

Alabi Fathiu
Software Developer
Go

⭐️ Support

If you found this useful, feel free to ⭐️ the repository.

