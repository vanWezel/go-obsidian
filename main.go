package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go-obsidian/gitlab"
	"go-obsidian/jira"
	"go-obsidian/outlook"
)

var home string

func init() {
	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	var path string
	flag.StringVar(&path, "path", "Documents/notes/Work/Eye", "Where live your markdown notes?")

	log.Printf("syncing to folder %s...\n", path)

	outlook.Sync(home, path)
	gitlab.Sync(home, path)
	jira.Sync(home, path)

	log.Println("completed.")
}
