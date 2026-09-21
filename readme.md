## VSCODE Atalhos

1. Shift + Alt + F > Auto formatacao

- docker
- postgres in docker
- swagger for docs
- golang migrate for migrations

# Install make on windows
1. Find
   Get-ChildItem "C:\Program Files" -Recurse -Filter make.exe -ErrorAction SilentlyContinue
   Get-ChildItem "C:\Program Files (x86)" -Recurse -Filter make.exe -ErrorAction SilentlyContinue
2. Test it directly
   & "C:\Program Files (x86)\GnuWin32\bin\make.exe" --version
3. Add it to PATH
   $env:Path += ";C:\Program Files (x86)\GnuWin32\bin"
4. Add it permanently
   Start → Environment Variables → Edit the system environment variables → Environment Variables → Path → Edit → New
   C:\Program Files (x86)\GnuWin32\bin


# Creating the project

```go
go mod init github.com/Leo3965/social
```

# Project Structure

### `bin/`

Contains compiled application binaries.

### `cmd/`

Contains the application's executable entry points. Each subdirectory represents a different executable or application within the project.

#### Subdirectories

* **`api/`** — Entry point for the API server. Contains the `main.go` and server-related initialization.
* **`migrate/`** — Entry point for database migrations and migration-related commands.

### `internal/`

Contains internal application packages that should not be imported by external projects.

Packages inside `internal/` are intended to be used only within this project. This helps keep implementation details private and encourages clear boundaries between the application's components.

### `docs/`

Contains API documentation and generated files, such as Swagger/OpenAPI documentation.

### `scripts/`

Contains scripts used to automate development and operational tasks, especially tasks related to setting up and running the application locally.


### `web/`
If it is a mono repo you can put the frontend related files here