# Lectures

This repository contains lecture slides for Course Go.

> [!NOTE]
> The course is currently a WIP! \
> All of the things stated here are still being determined and are due to change.

## Content

0. Course [[slides](https://lectures.course-go.dev/00-course.slide)]
    - Introduction & requirements
    - Course overview
1. Introduction [[slides](https://lectures.course-go.dev/01-introduction.slide) | [exercise](https://github.com/course-go/exercises/blob/master/01-workspace-setup-and-project-basics/README.md)]
    - Origins of Go
    - Key characteristics
    - Use cases & motivation
    - Editors & IDEs
    - Environment setup
    - Hello world!
2. Fundamentals #1 [[slides](https://lectures.course-go.dev/02-fundamentals.slide) | [exercise](https://github.com/course-go/exercises/blob/master/02-simple-katas/README.md)]
    - Packages & Visibility
    - Variables
    - Keywords
    - Data types
    - Control flow
    - Functions
    - Pointers
    - Structures
3. Fundamentals #2 [[slides](https://lectures.course-go.dev/03-fundamentals.slide) | [exercise](https://github.com/course-go/exercises/blob/master/03-data-structures-katas/README.md)]
    - Interfaces
    - Errors
    - Arrays
    - Slices
    - Maps
    - Range
4. Concurrency & parallelism [[slides](https://lectures.course-go.dev/04-concurrency-and-parallelism.slide) | [exercise](https://github.com/course-go/exercises/blob/master/04-concurrency-and-parallelism/README.md)]
    - Goroutines
    - Runtime
    - Channels
    - Select
    - Related packages
5. Advanced #1 ([slides](https://lectures.course-go.dev/05-advanced.slide) | [exercise](https://github.com/course-go/exercises/blob/master/05-generics-and-testing/README.md))
    - Generics
    - Packages
    - Testing
6. Advanced #2 [[slides](https://lectures.course-go.dev/06-advanced.slide) | [exercise](https://github.com/course-go/exercises/blob/master/06-pprof/README.md)]
    - Benchmarks
    - Optimizations
    - CGo
    - Unsafe & Reflect
7. REST APIs [[slides](https://lectures.course-go.dev/07-rest-api.slide) | [exercise](https://github.com/course-go/exercises/blob/master/07-rest-api/README.md)]
    - JSON
    - HTTP
    - REST API
    - HTTP package
    - Routers & Web frameworks
    - OpenAPI
    - Templating
8. Docker [[slides](https://lectures.course-go.dev/08-docker.slide) | [exercise](https://github.com/course-go/exercises/blob/master/08-docker/README.md)]
    - Virtualization 
    - Docker engine & basics
    - Dockerfile
    - YAML
    - Docker compose
    - Kubernetes
9. Databases [[slides](https://lectures.course-go.dev/09-databases.slide) | [exercise](https://github.com/course-go/exercises/blob/master/09-databases/README.md)]
    - SQL
    - RDBMSs
    - Migrations
    - sql
    - sqlx
    - sqlc
    - GORM
10. CI/CD [[slides](https://lectures.course-go.dev/10-cicd.slide) | [exercise](https://github.com/course-go/exercises/blob/master/10-cicd/README.md)]
    - Project setup
    - Git, GitHub & GitLab
    - GitHub Actions
    - CI/CD
    - GCP
11. Observability [[slides](https://lectures.course-go.dev/11-observability.slide) | [exercise](https://github.com/course-go/exercises/blob/master/11-prometheus/README.md)]
    - Health checks
    - Logging
    - Monitoring
    - Tracing
    - Proxying

## Tooling

The slides are made using the [go present syntax](https://pkg.go.dev/golang.org/x/tools/present).
You can run the slides locally using the go present CLI command.

### Installation

You can install the CLI command via go install like so:

```
go install golang.org/x/tools/cmd/present@latest
```

This installs the **present** executable into your 
**$GOPATH/bin** directory if **GOPATH** is set or the 
**$HOME/go/bin** directory otherwise.

### Usage

Make sure that the directory with the **present** binary is in your $PATH. After that, running the present tool is simple:

```
present
```

A webserver is run on localhost where you can preview the slides using your favourite web browser.

## Attribution

Some of the lecture slides are based on the [GoCourse](https://github.com/RedHatOfficial/GoCourse) 
project developed under [Red Hat](https://github.com/RedHatOfficial). The GoCourse is licensed under 
[CC BY-SA 4.0 DEED license](https://creativecommons.org/licenses/by-sa/4.0/deed.en), which this 
repository respects and therefore shares.

