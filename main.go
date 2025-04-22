package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	auth *AuthenticatedClient
}

func filter[T any](slice []T, predice func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predice(v) {
			result = append(result, v)
		}
	}
	return result
}

func formatJsonResponse[T any](t T) (string, error) {
	jsonBytes, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	return string(jsonBytes), nil
}

func formatDisplayResponse(t Task) string {
	response := fmt.Sprintf("Task: %s\nPriority: %d\nDescription: %s\nTags:%s\n", t.Content, t.Priority, t.Description, t.Labels)
	return response
}

type Config struct {
	EnvPath string `json:"env_path"`
}

func main() {


	content, err := os.ReadFile("/home/joe/dev/todoist_picker/config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

    godotenv.Load(config.EnvPath)

	Todoist := &App{
		auth: NewAuthenticatedClient(os.Getenv("TODOAPI"), os.Getenv("TODOURL")),
	}

	// log.Println("Getting today and overdue tasks:")
	data, err := Todoist.getTasks("(today|overdue)")

	if len(data.Result) == 0 {
		// log.Println("No today or overdue tasks found, proceeding to get all tasks from AUPP Admin:")
		data, err = Todoist.getTasks(url.QueryEscape(os.Getenv("SECONDARYTASK")))
	}

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if len(data.Result) == 0 {
		log.Println("No tasks found, exiting. you can relax now.")
		return
	}

	randomPickTask := rand.Intn(len(data.Result))

	fmt.Println("Tasks: ", formatDisplayResponse(data.Result[randomPickTask]))
}
