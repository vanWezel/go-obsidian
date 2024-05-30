package note

import (
	"fmt"
	"os"
)

type noteType string

const Outlook noteType = "OUTLOOK"
const Jira noteType = "JIRA"
const Gitlab noteType = "GITLAB"

func GetNotesPath(part noteType) string {
	base := "NOTES_PATH"

	if path := os.Getenv(fmt.Sprintf("%s_%s", base, part)); path != "" {
		return path
	}

	return os.Getenv(base)
}
