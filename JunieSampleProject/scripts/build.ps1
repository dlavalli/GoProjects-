# PowerShell script to build the application for multiple platforms

# Create the build directory if it doesn't exist
if (-not (Test-Path -Path "build")) {
    New-Item -ItemType Directory -Path "build" | Out-Null
    Write-Host "Created build directory"
}

# Build for Windows (amd64)
Write-Host "Building for Windows (amd64)..."
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o "build\myapp-windows-amd64.exe" .\cmd\myapp
if ($LASTEXITCODE -eq 0) {
    Write-Host "Windows build successful" -ForegroundColor Green
} else {
    Write-Host "Windows build failed" -ForegroundColor Red
    exit 1
}

# Build for Linux (amd64)
Write-Host "Building for Linux (amd64)..."
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "build\myapp-linux-amd64" .\cmd\myapp
if ($LASTEXITCODE -eq 0) {
    Write-Host "Linux build successful" -ForegroundColor Green
} else {
    Write-Host "Linux build failed" -ForegroundColor Red
    exit 1
}

# Build for macOS (amd64)
Write-Host "Building for macOS (amd64)..."
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o "build\myapp-darwin-amd64" .\cmd\myapp
if ($LASTEXITCODE -eq 0) {
    Write-Host "macOS build successful" -ForegroundColor Green
} else {
    Write-Host "macOS build failed" -ForegroundColor Red
    exit 1
}

Write-Host "All builds completed successfully!" -ForegroundColor Green
Write-Host "Binaries are available in the 'build' directory"