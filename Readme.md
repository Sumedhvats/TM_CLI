# Task CLI

A simple and efficient command-line task manager built with Go, using Cobra CLI framework and BoltDB for persistent storage.

## Features

- ✅ Add new tasks to your list
- 📋 List all pending tasks
- ✔️ Mark tasks as complete (removes them from the list)
- 💾 Persistent storage using BoltDB
- 🏠 Tasks stored in your home directory

## Installation

### Prerequisites
- Go 1.16 or higher

### Build from source
```bash
git clone https://github.com/Sumedhvats/task
cd task
go build -o task
```

### Install globally (optional)
```bash
go install
```

## Usage

### Add a new task
```bash
./task add "Buy groceries"
./task add "Complete the project documentation"
./task add "Call dentist for appointment"
```

### List all tasks
```bash
./task list
```

Output example:
```
1) Buy groceries
2) Complete the project documentation
3) Call dentist for appointment
```

### Complete tasks
Mark single task as complete:
```bash
./task do 1
```

Mark multiple tasks as complete:
```bash
./task do 1 3
```

Output example:
```
Completed task 1
Completed task 3
```

### Help
```bash
./task --help
./task add --help
./task list --help
./task do --help
```

## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `add` | Add a new task to your list | `task add "New task"` |
| `list` | Display all pending tasks | `task list` |
| `do` | Mark one or more tasks as complete | `task do 1 2 3` |

## Technical Details

### Architecture
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) for command-line interface
- **Database**: [BoltDB](https://github.com/boltdb/bolt) for embedded key-value storage
- **Storage Location**: `~/tasks.db` (in user's home directory)

### Project Structure
```
.
├── cmd/
│   ├── add.go      # Add command implementation
│   ├── do.go       # Complete task command implementation
│   ├── list.go     # List tasks command implementation
│   └── root.go     # Root command and CLI setup
├── db/
│   └── tasksDB.go  # Database operations and task management
├── main.go         # Application entry point
└── go.mod          # Go module dependencies
```

### Database Schema
Tasks are stored as key-value pairs where:
- **Key**: Auto-incrementing integer (converted to bytes)
- **Value**: Task description as string
- **Bucket**: `"tasks"`

## Dependencies

```go
require (
    github.com/boltdb/bolt v1.3.1
    github.com/spf13/cobra v1.8.0
)
```

## Error Handling

The application handles various error scenarios:
- Invalid task numbers when completing tasks
- Database connection issues
- Parsing errors for task IDs
- Missing or corrupted database files

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Future Enhancements

Potential improvements for future versions:
- [ ] Task priorities (high, medium, low)
- [ ] Due dates and reminders
- [ ] Task categories/tags
- [ ] Search functionality
- [ ] Task editing capabilities
- [ ] Export tasks to different formats
- [ ] Task completion history
- [ ] Colorized output

## License

This project is open source. Please check the repository for license details.

## Author

[Sumedhvats](https://github.com/Sumedhvats)

---

*Built with ❤️ using Go*