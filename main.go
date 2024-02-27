package main

import (
	"flag"
)

// Scan path and its sub folders for .git files?
func scan(path string) {
	println("scanning...")
}

// Generate graph of git contributions
func stats(email string) {
	println("generating stats...", email)
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "myEmail@email.com", "email to scan")

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
