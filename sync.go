package todoist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Args struct {
	Id        string   `json:"id"`
	Content   string   `json:"content"`
	ProjectId string   `json:"project_id"`
	Labels    []string `json:"labels"`
}

type RequestParameter struct {
	Type   string `json:"type,omitempty"`
	TempId string `json:"temp_id,omitempty"`
	Uuid   string `json:"uuid,omitempty"`
	Args   Args   `json:"args,omitempty"` // required
}

type TaskResponse struct {
	FullSync      bool                   `json:"full_sync"`
	SyncStatus    map[string]interface{} `json:"sync_status"`
	SyncToken     string                 `json:"sync_token"`
	TempIdMapping map[string]interface{} `json:"temp_id_mapping"`
}

type Command struct {
	Commands []RequestParameter `json:"commands"`
}

func (c *Client) CreateTaskSync(param []RequestParameter) (*http.Response, *TaskResponse, error) {
	j, err := json.Marshal(param)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Println(command)
	//var data = strings.NewReader(`commands=`)
	var data = fmt.Sprintf("commands=%s", string(j))
	fmt.Println(data)
	req, err := http.NewRequest("POST", DefaultSyncUrl+"/sync", strings.NewReader(data))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, result, err := c.doRequest(req)
	if err != nil {
		return response, nil, err
	}

	var taskResponse TaskResponse
	err = json.Unmarshal(result, &taskResponse)
	if err != nil {
		return response, nil, err
	}
	return response, &taskResponse, nil
}

func (c *Client) CompleteTaskSync(param []RequestParameter) (*http.Response, *TaskResponse, error) {
	j, err := json.Marshal(param)
	if err != nil {
		return nil, nil, err
	}

	var data = fmt.Sprintf("commands=%s", string(j))
	fmt.Println(data)

	req, err := http.NewRequest("POST", DefaultSyncUrl+"/sync", strings.NewReader(data))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, result, err := c.doRequest(req)
	if err != nil {
		return nil, nil, err
	}

	var taskResponse TaskResponse
	err = json.Unmarshal(result, &taskResponse)
	if err != nil {
		return response, nil, err
	}
	return response, &taskResponse, nil
}

func (c *Client) UncompleteTaskSync(param []RequestParameter) (*http.Response, *TaskResponse, error) {
	j, err := json.Marshal(param)
	if err != nil {
		return nil, nil, err
	}

	var data = fmt.Sprintf("commands=%s", string(j))
	fmt.Println(data)

	req, err := http.NewRequest("POST", DefaultSyncUrl+"/sync", strings.NewReader(data))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, result, err := c.doRequest(req)
	if err != nil {
		return response, nil, err
	}

	var taskResponse TaskResponse
	err = json.Unmarshal(result, &taskResponse)
	if err != nil {
		return response, nil, err
	}
	return response, &taskResponse, nil
}

func (c *Client) DeleteTaskSync(param []RequestParameter) (*http.Response, *TaskResponse, error) {
	j, err := json.Marshal(param)
	if err != nil {
		return nil, nil, err
	}

	var data = fmt.Sprintf("commands=%s", string(j))
	fmt.Println(data)

	req, err := http.NewRequest("POST", DefaultSyncUrl+"/sync", strings.NewReader(data))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, result, err := c.doRequest(req)
	if err != nil {
		return response, nil, err
	}

	var taskResponse TaskResponse
	err = json.Unmarshal(result, &taskResponse)
	if err != nil {
		return response, nil, err
	}
	return response, &taskResponse, nil
}
