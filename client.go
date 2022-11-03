package todoist

import (
	"io"
	"net/http"
)
import "fmt"

const DefaultRestUrl string = "https://api.todoist.com/rest/v2"
const DefaultSyncUrl string = "https://api.todoist.com/sync/v9"

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	Host       string
	Base       string
}

func NewClient(apiKey string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
	}
}

func (c *Client) newRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", DefaultRestUrl, path), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) doRequest(req *http.Request) (*http.Response, []byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	response, err := c.HttpClient.Do(req)
	if err != nil {
		return response, nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return response, nil, err
	}
	return response, body, err
}
