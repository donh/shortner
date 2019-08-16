# shortner [![CircleCI Build Status](https://circleci.com/gh/donh/shortner.svg?style=shield)](https://circleci.com/gh/donh/shortner) [![Go Report Card](https://goreportcard.com/badge/github.com/donh/shortner)](https://goreportcard.com/report/github.com/donh/shortner) [![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/CircleCI-Public/circleci-demo-go/master/LICENSE.md)

## Installation
- [Go 1.12](https://golang.org/dl/)
  ```bash
  $ wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz
  $ sudo tar -C /usr/local -xzf go1.12.9.linux-amd64.tar.gz
  $ mkdir $HOME/go
  $ export PATH=$PATH:/usr/local/go/bin
  $ export GOPATH=$HOME/go
  $ go get github.com/go-sql-driver/mysql
  $ go get github.com/jmoiron/sqlx
  $ go get github.com/rs/cors
  $ go get gopkg.in/yaml.v2
  ```
- Git
  ```bash
  $ sudo apt-get install git
  $ mkdir $HOME/code
  $ cd $HOME/code
  $ git clone https://github.com/donh/shortner.git
  $ cd $HOME/code/shortner
  ```
- MySQL
  ```bash
  $ sudo apt-get install mysql-server
  $ mysql -h 127.0.0.1 -P 3306 -u root -p < ./scripts/tree.sql
  ```

## Lint
- Go
  ```bash
  $ go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
  ```
  ```bash
  $ golangci-lint run --fix
  ```

## Build
```bash
$ go build -i -v ./...
```

## Run
```bash
$ go run ./main.go
```

## API
- API list
  - /api/v1/add
  - /{key}

### API Example
- **/api/v1/add**
  - Add a shorten URL
  - method
    - POST
  - Request
  ```bash
  POST http://localhost:8000/api/v1/add
  ```
  ```bash
  {
    "url": "https://github.com/donh/shortner"
  }
  ```
  - Response
  ```bash
  {
    "Result": "http://localhost:8000/1g3e",
    "Status": 200,
    "Error": "",
    "Time": "2019-08-16 16:36:24"
  }
  ```
- **/{key}**
  - Redirect to original URL
  - method
    - GET
  - Request
  ```bash
  GET http://localhost:8000/1g3e
  ```
  - Response
    - The content of "https://github.com/donh/shortner"
