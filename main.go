package main

import (
	"flag"
	"log"
	"os"

	"main/gitlab"
	"main/outlook"
)

var home string

func init() {
	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
}

func main() {
	var path string
	flag.StringVar(&path, "path", "Documents/notes/Work/Eye", "Where live your markdown notes?")

	log.Printf("syncing to folder %s...\n", path)

	outlook.Sync(home, path)
	gitlab.Sync(home, path)

	log.Println("completed.")
}
