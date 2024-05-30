package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go-obsidian/jira/model"
)

type Client struct {
	client   *http.Client
	url      string
	username string
	token    string
}

func New(domain string, username string, token string) *Client {
	return &Client{
		client:   &http.Client{},
		url:      fmt.Sprintf("https://%s.atlassian.net", domain),
		username: username,
		token:    token,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return buff, nil
}

type Payload struct {
	Expand       []string `json:"expand"`
	Fields       []string `json:"fields"`
	FieldsByKeys bool     `json:"fieldsByKeys"`
	Jql          string   `json:"jql"`
}

func (c *Client) FetchIssues() ([]model.Issue, error) {
	body, err := json.Marshal(Payload{
		Jql:    "status in (\"Blocked\", \"In Progress\", \"Testing & Acceptance\") AND assignee in (currentUser())",
		Expand: []string{},
		Fields: []string{
			"summary",
			"status",
			"description",
		},
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/rest/api/2/search", c.url), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.token)

	buff, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response model.Response
	if err = json.Unmarshal(buff, &response); err != nil {
		return nil, err
	}

	return response.Issues, nil
}
