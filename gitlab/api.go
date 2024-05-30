package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go-obsidian/gitlab/model"
)

type Client struct {
	token  string
	client *http.Client
	url    string
}

func New(token string) *Client {
	return &Client{
		client: &http.Client{},
		url:    "https://gitlab.com",
		token:  token,
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
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (c *Client) FetchStarredProjects() ([]model.Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v4/projects?starred=true", c.url), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("PRIVATE-TOKEN", c.token)

	buff, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var projects []model.Project
	if err = json.Unmarshal(buff, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (c *Client) FetchMergeRequests(projectID int) ([]model.MergeRequest, error) {
	// Fetch open merge requests for the specified project
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v4/projects/%d/merge_requests?state=opened&scope=assigned_to_me", c.url, projectID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("PRIVATE-TOKEN", c.token)

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var mergeRequests []model.MergeRequest
	if err = json.Unmarshal(body, &mergeRequests); err != nil {
		return nil, err
	}

	return mergeRequests, nil
}

func (c *Client) FetchUser() (*model.User, error) {
	// Fetch open merge requests for the specified project
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v4/user/", c.url), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("PRIVATE-TOKEN", c.token)

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *Client) FetchReviewRequests(projectID int, userID int) ([]model.MergeRequest, error) {
	// Fetch open merge requests for the specified project
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v4/projects/%d/merge_requests?state=opened&reviewer_id=%d", c.url, projectID, userID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("PRIVATE-TOKEN", c.token)

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var mergeRequests []model.MergeRequest
	if err = json.Unmarshal(body, &mergeRequests); err != nil {
		return nil, err
	}

	return mergeRequests, nil
}
