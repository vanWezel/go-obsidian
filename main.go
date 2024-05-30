package main

import (
	"log"

	"github.com/joho/godotenv"

	"go-obsidian/gitlab"
	"go-obsidian/jira"
	"go-obsidian/outlook"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("starting...")
	outlook.Sync()
	gitlab.Sync()
	jira.Sync()
	log.Println("completed.")
}
