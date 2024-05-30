package model

import "time"

type MergeRequest struct {
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	UserNotesCount      int      `json:"user_notes_count"`
	SourceBranch        string   `json:"source_branch"`
	State               string   `json:"state"`
	WebURL              string   `json:"web_url"`
	DetailedMergeStatus string   `json:"detailed_merge_status"`
	ProjectId           int      `json:"project_id"`
	HeadPipeline        Pipeline `json:"head_pipeline"`
}

type Pipeline struct {
	Id         int         `json:"id"`
	Iid        int         `json:"iid"`
	ProjectId  int         `json:"project_id"`
	Sha        string      `json:"sha"`
	Ref        string      `json:"ref"`
	Status     string      `json:"status"`
	Source     string      `json:"source"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	WebUrl     string      `json:"web_url"`
	BeforeSha  string      `json:"before_sha"`
	Tag        bool        `json:"tag"`
	YamlErrors interface{} `json:"yaml_errors"`
	User       User        `json:"user"`
}
