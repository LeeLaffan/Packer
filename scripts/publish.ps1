param(
    [Parameter(Mandatory=$true)]
    [Alias("p")]
    [ValidateSet("win", "linux")]
    [string]$platform
)

if (Test-Path bin) { 
    Remove-Item bin -Recurse -Force 
} mkdir bin

if ($platform -eq "win") {
    $env:GOOS = "windows"
    go build -o bin/pack.exe .
} else {
    $env:GOOS = "linux"
    go build -o bin/pack .
}
