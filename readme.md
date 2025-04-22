# Introduction
One of the reason task keep mounted is mostly due to the one of the well known decision paralysis, where you have too much freedom to choose and endpoint choosing nothing.
This project will pick only one task for you to begin with, it will start with `today and overdue` task, if there's none, then it will proceed to your selected `project`.

# Setup
Make sure to copy `env.example` to `.env` file. 

```bash
cp env.example .env
```

Get your todoist api key from the todoist application by `profile -> settings -> integrations -> developer` and api token.
Your project name must be the exact name of your project on todoist application.

# Build
Run the following command to build the application
```bash
go mod tiny # install dependency
go build -o todopicker
```
