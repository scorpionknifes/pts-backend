# Pass the Story Backend

[![Build Status](https://dev.azure.com/chengzhenyang/pts-backend/_apis/build/status/scorpionknifes.pts-backend?branchName=master)](https://dev.azure.com/chengzhenyang/pts-backend/_build/latest?definitionId=3&branchName=master) [![Go Report Card](https://goreportcard.com/badge/github.com/scorpionknifes/pts-backend)](https://goreportcard.com/report/github.com/scorpionknifes/pts-backend)
[![Docker Status](https://img.shields.io/docker/cloud/build/zhenknz/pts-backend)](https://hub.docker.com/repository/docker/zhenknz/pts-backend/)

[Hosted on Azure](https://pts-backend.azurewebsites.net/) and [Self Hosted](https://zhenk.ml/)

This project is an assignment for Microsoft Student Accelerator NZ 2020 Phase 2.

Pass the Story Backend for [frontend](https://github.com/scorpionknifes/pts-frontend).

Using [Docker](https://hub.docker.com/repository/docker/zhenknz/pts-backend) image to run on azure.



The following project uses:
- Golang
- [gqlgen](https://gqlgen.com/)
- [GORM](https://gorm.io/index.html)
- SQL Server
- Docker

[Microsoft Student Accelerator](http://aka.ms/nzmsawebsite)
[2020-Phase-2](https://github.com/NZMSA/2020-Phase-2)

## Setup

1. Install Golang
2. git clone project
3. setup .env file using template (connect to azure SQL server)
4. Read ```How to use```

## How to use

Golang Commands - go mod will auto install dependency
```bash
go run server.go
go build server.go
go test ./...
```

Docker Commands
```bash
docker build -t pts-backend .
docker run --publish <PORT>:<PORT> --name test --rm pts-backend
docker stop test
```

gqlgen Commans
```go
gqlgen generate // run when change schema.graphql
```

## Example Queries
<details>
  <summary>Click to expand!</summary>
  ```go
# Write your query or mutation here
query stories {
  stories{
    id
    name
    count
    people
    tags
    createdAt
    updatedAt
  }
}

mutation createStory {
  createStory(input: {
    name: "Anonymous's 2000",
    tags: "example, cool, love"
  }){
    id
    name
    tags
  }
}

subscription subscriptionStory {
  stories{
    id
    name
    count
    people
    tags
    createdAt
    updatedAt
  }
}

query Story {
  story(id: 1) {
    id
    name
    turns{
      id
      value
      user{
        id
      }
    }
    count
    people
    tags
    createdAt
    updatedAt
  }
}
  ```

</details>