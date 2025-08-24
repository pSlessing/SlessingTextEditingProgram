This is a project im using, both to learn GO, but also to implement a larger type of application, in this case, being a text editor :D

### Current Features

- Main Loop using commands on the status bar
- Write loop using the command write
- Open command used to open a file
- Save command for saving current file
- SaveAs command to specify the filename
- Row and column count indicating current cursor position

### Upcoming features
- Syntax highlighting for different programming languages
- Color customizability, saved between sessions
- Fuzzy searching for opening files in the current directory
- Shorthand aliasing of commands for faster workflow
- Cursor movement in Function mode
- Cursor position being transient between loops (Except for when opening a new file)


### Setup
# 1. Initialize the Go module
go mod init ste-text-editor

# 2. Download and organize dependencies
go mod tidy

# 3. Download all dependencies to local cache
go mod download

# 4. Verify the dependencies haven't been tampered with
go mod verify

### Running the editor

## Linux

# Run directly using GO
go run .

# Or if GO is mad about unused variables due to a test build
go run -gcflags="-l" .
