package scaleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetAllProjectTasks(project string) []Task {
	url := "https://api.scale.com/v1/tasks"

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	q := req.URL.Query()
	q.Add("project", project)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(c.apiKey, "")

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	var t TasksResponse
	err = Unmarshal(resBody, &t)
	if err != nil {
		fmt.Printf("Error unmarshaling data: %v", err)
	}

	return t.Tasks
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
