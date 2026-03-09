# Windows Installation Script for gh-stats

$ErrorActionPreference = "Stop"
$EXECUTABLE_NAME = "gh-stats.exe"
$BINARY_PATH = ".\bin\gh-stats-windows.exe"

# Define the user's local AppData installation path
$INSTALL_DIR = "$env:LOCALAPPDATA\gh-stats\bin"

# Check if binary exists in the current directory
if (-not (Test-Path $BINARY_PATH)) {
    Write-Host "❌ Could not find $BINARY_PATH. Please run this script from the unzipped directory." -ForegroundColor Red
    Exit
}

# Create the installation directory if it doesn't exist
if (-not (Test-Path $INSTALL_DIR)) {
    New-Item -ItemType Directory -Force -Path $INSTALL_DIR | Out-Null
}

# Copy the binary to the installation folder
Write-Host "🚀 Installing $EXECUTABLE_NAME to $INSTALL_DIR..." -ForegroundColor Cyan
Copy-Item -Path $BINARY_PATH -Destination "$INSTALL_DIR\$EXECUTABLE_NAME" -Force

# Add the installation folder to the User PATH if it's not already there
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notmatch [regex]::Escape($INSTALL_DIR)) {
    Write-Host "🔗 Adding $INSTALL_DIR to your User PATH..." -ForegroundColor Cyan
    $newPath = $userPath + ";" + $INSTALL_DIR
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "⚠️ PATH updated! You may need to restart your terminal (or open a new tab) for the command to be recognized." -ForegroundColor Yellow
}

Write-Host "✅ Success! You can now run 'gh-stats' in any git repository." -ForegroundColor Green
