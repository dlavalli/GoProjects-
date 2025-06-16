# Project Guidelines
    
This is a placeholder of the project guidelines for Junie.
Replace this text with any project-level instructions for Junie, e.g.:

* What is the project structure

This is the recommended project structure. For more information, you can also check https://github.com/golang-standards/project-layout
myapp/
|___ cmd/
|    |___myapp/                 # Main application entry point
|        |____ main.go
|___ internal/                  # Private application and library code
|    |___ auth/                 # Example: internal authentication logic
|    |___ db/                   # Example: internal DB access code
|___ pkg/                       # Public library code (can be used by other projects)
|    |___ utils/                # Example: reusable utility functions
|___ api/                       # OpenAPI/Swagger specs, Protobuf, gRPC definitions
|___ configs/                   # Configuration files (JSON, YAML, TOML, etc.)
|___ deployments/               # Docker, Kubernetes, Helm, etc.
|___ scripts/                   # Bash, CLI tools for setup, migration, CI, etc.
|___ web/                       # Web frontend (React, static files, templates, etc.)
|___ test/                      # Additional test data and utilities
|___ go.mod
|___ go.sum
|___ README.md

* Whether Junie should run tests to check the correctness of the proposed solution

Yes, it should include unit tests to verify the functionality end-to-end. Additionally, if the solution involves performance or latency considerations, make sure to include benchmark tests in Go as well.  

* How does Junie run tests (once it requires any non-standard approach)
* Whether Junie should build the project before submitting the result

Yes, the project should be built, and the build artifacts should be placed in a build folder at the project root. This folder should include builds for Windows, Linux, and macOS platforms.

* Any code-style related instructions

As an option you can ask Junie to create these guidelines for you.

Junie, after creating the project, ensure all dependencies are listed in the go.mod file. Then, run go mod download to fetch them. Before starting the application, run go mod tidy to clean up any unused dependencies.

- If you are building CRUD application, the use postgres as the default database choice
- Make sure to always include a Dockerfile under deployments.
- 