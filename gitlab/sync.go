package gitlab

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"go-obsidian/gitlab/model"
	"go-obsidian/note"
)

func Sync() {
	client := New(os.Getenv("GITLAB_TOKEN"))
	log.Println("fetching starred projects on gitlab...")

	projects, err := client.FetchStarredProjects()
	if err != nil {
		log.Println("got error", err)
		return
	}

	outputMerge := new(bytes.Buffer)
	outputReview := new(bytes.Buffer)

	user, err := client.FetchUser()
	if err != nil {
		log.Println("got error", err)
		return
	}

	for _, project := range projects {
		log.Printf("fetching merge requests for project '%s'...\n", project.Name)

		mergeRequests, err := client.FetchMergeRequests(project.ID)
		if err != nil {
			continue
		}

		writeProject(outputMerge, project, mergeRequests)

		reviewRequests, err := client.FetchReviewRequests(project.ID, user.Id)
		if err != nil {
			continue
		}

		writeProject(outputReview, project, reviewRequests)
	}

	note.Save(outputMerge, path.Join(note.GetNotesPath(note.Jira), "Merge Requests.md"))
	note.Save(outputReview, path.Join(note.GetNotesPath(note.Jira), "Review Requests.md"))
	log.Println("gitlab completed.")
}

func writeProject(output *bytes.Buffer, project model.Project, mergeRequests []model.MergeRequest) {
	if len(mergeRequests) == 0 {
		return
	}

	fmt.Fprintf(output, "#### %s\n", project.Name)

	for _, mr := range mergeRequests {
		footer := []string{fmt.Sprintf("> [View](%s)", mr.WebURL)}
		if mr.HeadPipeline.Status != "" {
			footer = append(footer, fmt.Sprintf("[🧑‍🚒 %s](%s)", mr.HeadPipeline.Status, mr.HeadPipeline.WebUrl))
		}

		footer = append(footer, fmt.Sprintf("`%s`", mr.SourceBranch))
		footer = append(footer, fmt.Sprintf("💭 %d", mr.UserNotesCount))
		footer = append(footer, fmt.Sprintf("(%s)", mr.DetailedMergeStatus))

		fmt.Fprintf(output, "> %s \n%s \n\n\n", mr.Title, strings.Join(footer, " | "))
	}

	fmt.Fprintln(output)
}
