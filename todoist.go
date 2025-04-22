package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Due struct {
	Date        string `json:"date"`
	TimeZone    string `json:"timezone"`
	IsRecurring bool   `json:"is_recurring"`
}
type Task struct {
	UserId      string    `json:"user_id"`
	Id          string    `json:"id"`
	ProjectId   string    `json:"project_id"`
	SectionId   string    `json:"section_id"`
	ParentId    string    `json:"parent_id"`
	AddedByUid  string    `json:"added_by_uid"`
	Labels      []string  `json:"labels"`
	Deadline    time.Time `json:"deadline"`
	Checked     bool      `json:"checked"`
	IsDeleted   bool      `json:"is_deleted"`
	AddedAt     time.Time `json:"added_at"`
	CompletedAt time.Time `json:"completed_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Due         Due       `json:"due"`
	Priority    int       `json:"priority"`
	ChildOrder  int       `json:"child_order"`
	Content     string    `json:"content"`
}

type TodoistResponse struct {
	Result []Task `json:"results"`
}

func (app *App) getTasks(filter string) (*TodoistResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := "/tasks"
	if filter != "" {
		url = fmt.Sprintf("/tasks/filter?query=%s", filter)
	}

	task, err := app.auth.Get(ctx, url, nil)
	if err != nil {
		fmt.Println("Error getting tasks:", err)
	}

	var data *TodoistResponse

	err = json.NewDecoder(task.Body).Decode(&data)

	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	return data, nil
}
