package gitlab

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"main/gitlab/model"
)

func Sync(home string, path string) {
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

	save(outputMerge, home, path, "Merge Requests.md")
	save(outputReview, home, path, "Review Requests.md")
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

func save(output *bytes.Buffer, home string, path string, name string) {
	fmt.Fprintf(output, "_last updated @ %s_ \n\n", time.Now().Format("2006-01-02 15:04"))

	file, err := os.OpenFile(fmt.Sprintf("%s/%s/%s", home, path, name), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println("got error", err)
		return
	}

	if _, err := io.Copy(file, output); err != nil {
		log.Println("got error", err)
		return
	}

	log.Println("gitlab sync completed.")
}
