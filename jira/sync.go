package jira

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"

	"go-obsidian/jira/model"
	"go-obsidian/note"
)

func Sync() {
	domain := os.Getenv("JIRA_DOMAIN")
	token := os.Getenv("JIRA_TOKEN")
	username := os.Getenv("JIRA_USERNAME")

	output := new(bytes.Buffer)
	client := New(domain, username, token)

	issues, err := client.FetchIssues()
	if err != nil {
		log.Println("got error", err)
		return
	}

	ordered := map[string][]model.Issue{}
	for _, issue := range issues {
		status := issue.Fields.Status.Name
		list, ok := ordered[status]
		if !ok {
			list = []model.Issue{}
			ordered[status] = list
		}

		list = append(list, issue)
		ordered[status] = list
	}

	for _, orderedIssues := range ordered {
		for _, issue := range orderedIssues {
			writeIssue(output, domain, issue)
		}
	}

	note.Save(output, path.Join(note.GetNotesPath(note.Jira), "Jira.md"))
	log.Println("jira completed.")
}

func writeIssue(output *bytes.Buffer, domain string, issue model.Issue) {
	/*
		> [! task] TDR-1643 `In Progress`
		> [BE] set-up S3 trigger for awareness metrics with Risk
	*/
	link := fmt.Sprintf("> [View](https://%s.atlassian.net/browse/%s)", domain, issue.Key)

	var variant string
	switch issue.Fields.Status.Name {
	case "Blocked":
		variant = "warning"
	case "In Progress":
		variant = "task"
	default:
		variant = "faq"
	}

	fmt.Fprintf(output, ">[! %s] %s \n %s - %s \n %s \n\n\n", variant, issue.Fields.Status.Name, issue.Key, issue.Fields.Summary, link)
	fmt.Fprintln(output)
}
