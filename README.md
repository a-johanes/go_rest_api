# Go REST API

Simple REST API project using Golang, Fiber `2.7.1`, and GORM `1.21.6`

## Setup

Make sure you have Go installed ([download](https://golang.org/dl/)) and properly setted up. Version `1.14` or higher is required.

1. Download dependencies

```bash
$ go mod download
```

2. Setup the database (Postgres) and change the `dns` value in `models/setup.go` to match the database config

3. Run, the default port is `4321`

```bash
$ go run main.go
# or
$ go build && ./go_rest_api
```
