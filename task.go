package todoist

import (
	"net/http"
	"strconv"
)
import "encoding/json"
import "fmt"
import "bytes"

const TasksUrl = DefaultRestUrl + "/tasks"

type Task struct {
	Id           string `json:"id"`
	ProjectId    string `json:"project_id"`
	Content      string `json:"content"`
	Completed    bool   `json:"completed"`
	LabelIds     []int  `json:"label_ids"`
	Order        int    `json:"order"`
	Indent       int    `json:"indent"`
	Priority     int    `json:"priority"`
	Due          Due    `json:"due"`
	Url          string `json:"url"`
	CommentCount int    `json:"comment_count"`
}

type jTask struct {
	Id           string `json:"id"`
	ProjectId    string `json:"project_id"`
	Content      string `json:"content"`
	Completed    bool   `json:"completed"`
	LabelIds     []int  `json:"label_ids"`
	Order        int    `json:"order"`
	Indent       int    `json:"indent"`
	Priority     int    `json:"priority"`
	Due          Due    `json:"due"`
	Url          string `json:"url"`
	CommentCount int    `json:"comment_count"`
}

type NewTask struct {
	Content     string `json:"content"` // required
	ProjectId   int    `json:"project_id,omitempty"`
	Order       int    `json:"order,omitempty"`
	LabelIds    []int  `json:"label_ids,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	DueString   string `json:"due_string,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	DueDatetime string `json:"due_datetime,omitempty"`
	DueLang     string `json:"due_lang,omitempty"`
}

type Due struct {
	String   string `json:"string"` // required
	Date     string `json:"date"`   // required
	Datetime string `json:"datetime"`
	Timezone string `json:"timezone"`
}

type CompletedTask struct {
	Content       string `json:"content"`
	MetaData      string `json:"meta_data"`
	UserId        string `json:"user_id"`
	TaskId        string `json:"task_id"`
	ProjectId     string `json:"project_id"`
	CompletedDate string `json:"completed_date"`
	Id            string `json:"id"`
}

type jCompletedTask struct {
	Content       string `json:"content"`
	MetaData      string `json:"meta_data"`
	UserId        int    `json:"user_id"`
	TaskId        string `json:"task_id"`
	ProjectId     string `json:"project_id"`
	CompletedDate string `json:"completed_date"`
	Id            string `json:"id"`
}

type CompletedTaskResponse struct {
	Items []CompletedTask `json:"items"`
}

func (t *Task) UnmarshalJSON(b []byte) error {
	var jt jTask

	err := json.Unmarshal(b, &jt)
	if err != nil {
		return err
	}

	// Parse the Id as a string
	t.Id = jt.Id

	t.ProjectId = jt.ProjectId
	t.Content = jt.Content
	t.Completed = jt.Completed
	t.LabelIds = jt.LabelIds
	t.Order = jt.Order
	t.Indent = jt.Indent
	t.Priority = jt.Priority
	t.Due = jt.Due
	t.Url = jt.Url
	t.CommentCount = jt.CommentCount

	return nil
}

func (ct *CompletedTask) UnmarshalJSON(b []byte) error {
	var jct jCompletedTask

	err := json.Unmarshal(b, &jct)
	if err != nil {
		return err
	}

	// Parse the Id as a string
	ct.Content = jct.Content
	ct.MetaData = jct.MetaData
	ct.UserId = strconv.Itoa(jct.UserId)
	ct.TaskId = jct.TaskId
	ct.ProjectId = jct.ProjectId
	ct.CompletedDate = jct.CompletedDate
	ct.Id = jct.Id

	return nil
}

func (c *Client) GetTasks() (*http.Response, []Task, error) {
	req, err := http.NewRequest("GET", TasksUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	response, result, err := c.doRequest(req)
	if err != nil {
		return response, nil, err
	}

	var tasks []Task

	err = json.Unmarshal(result, &tasks)
	if err != nil {
		return response, nil, err
	}

	return response, tasks, nil
}

func (c *Client) GetTask(id string) (*http.Response, *Task, error) {
	var task Task

	req, err := http.NewRequest("GET", fmt.Sprintf(TasksUrl+"/%s", id), nil)
	if err != nil {
		return nil, nil, err
	}

	response, result, err := c.doRequest(req)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(result, &task)
	if err != nil {
		return response, nil, err
	}

	return response, &task, nil
}

func (c *Client) CreateTask(t *NewTask) (*http.Response, *Task, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", TasksUrl, bytes.NewBuffer(j))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	response, result, err := c.doRequest(req)
	if err != nil {
		return response, nil, err
	}

	var task Task
	err = json.Unmarshal(result, &task)
	if err != nil {
		return nil, nil, err
	}

	return response, &task, nil
}

func (c *Client) CloseTask(id string) (*http.Response, error) {
	req, err := http.NewRequest("POST", TasksUrl+"/"+id+"/close", nil)
	if err != nil {
		return nil, err
	}

	response, _, err := c.doRequest(req)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) ReopenTask(id string) (*http.Response, error) {
	req, err := http.NewRequest("POST", TasksUrl+"/"+id+"/reopen", nil)
	if err != nil {
		return nil, err
	}

	response, _, err := c.doRequest(req)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) DeleteTask(id string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", TasksUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	response, _, err := c.doRequest(req)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) UpdateTask(t *Task) (*http.Response, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", TasksUrl+"/"+t.Id, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	response, _, err := c.doRequest(req)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) GetCompletedTasks() (*http.Response, []CompletedTask, error) {
	req, err := http.NewRequest("GET", DefaultSyncUrl+"/completed/get_all", nil)
	if err != nil {
		return nil, nil, err
	}

	response, result, err := c.doRequest(req)
	if err != nil {
		return nil, nil, err
	}

	var ctr CompletedTaskResponse
	var completedTasks []CompletedTask

	err = json.Unmarshal(result, &ctr)
	if err != nil {
		return response, nil, err
	}

	completedTasks = ctr.Items
	return response, completedTasks, nil
}

func (c *Client) GetCompletedTask(id string) (*http.Response, *CompletedTask, error) {
	reponse, completedTasks, err := c.GetCompletedTasks()
	if err != nil {
		return reponse, nil, err
	}

	for _, ct := range completedTasks {
		if ct.TaskId == id {
			return reponse, &ct, nil
		}
	}

	return reponse, nil, fmt.Errorf("completed task: %s not found", id)
}

func (c *Client) ConvertCompletedTasksToTasks(completedTasks []CompletedTask) []Task {
	var tasks []Task

	for _, completedTask := range completedTasks {
		tasks = append(tasks, Task{
			Id:        completedTask.TaskId,
			ProjectId: completedTask.ProjectId,
			Content:   completedTask.Content,
			Completed: true,
		})
	}

	return tasks
}
