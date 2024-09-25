# Lectures

This repository contains lecture slides for Course Go.

## Content

0. Course [[slides](https://lectures.course-go.dev/00-course.slide)]
    - Introduction
    - Requirements
    - Overview
1. Introduction [[slides](https://lectures.course-go.dev/01-introduction.slide) | [exercise](https://github.com/course-go/exercises/blob/master/01-workspace-and-project-basics/README.md)]
    - Introduction to Go
    - IDEs & Editors
    - Installing Go
    - Running Go
    - Project basics
    - Resources
2. Fundamentals #1 [[slides](https://lectures.course-go.dev/02-fundamentals.slide) | [exercise](https://github.com/course-go/exercises/blob/master/02-katas/README.md)]
    - Built-in types
    - Variables
    - Control flow
    - Functions
    - Custom types
    - Pointers
3. Fundamentals #2 [[slides](https://lectures.course-go.dev/03-fundamentals.slide) | [exercise](https://github.com/course-go/exercises/blob/master/03-katas/README.md)]
    - Interfaces
    - Errors
    - Arrays
    - Slices
    - Maps
    - Range
4. Concurrency & parallelism [[slides](https://lectures.course-go.dev/04-concurrency-and-parallelism.slide) | [exercise](https://github.com/course-go/exercises/blob/master/04-concurrency/README.md)]
    - Goroutines
    - Runtime
    - Channels
    - Select
    - Related packages
5. Advanced #1 [[slides](https://lectures.course-go.dev/05-advanced.slide) | [exercise](https://github.com/course-go/exercises/blob/master/05-generics-and-testing/README.md)]
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
8. Containers [[slides](https://lectures.course-go.dev/08-containers.slide) | [exercise](https://github.com/course-go/exercises/blob/master/08-docker/README.md)]
    - Containerization
    - Docker
    - Kubernetes
9. Databases [[slides](https://lectures.course-go.dev/09-databases.slide) | [exercise](https://github.com/course-go/exercises/blob/master/09-databases/README.md)]
    - SQL
    - RDBMSs
    - Migrations
    - sql
    - sqlx
    - sqlc
    - GORM
10. Infrastructure [[slides](https://lectures.course-go.dev/10-infrastructure.slide) | [exercise](https://github.com/course-go/exercises/blob/master/10-infrastructure/README.md)]
    - CI/CD
    - Infrastructure
    - GCP
11. Observability [[slides](https://lectures.course-go.dev/11-observability.slide) | [exercise](https://github.com/course-go/exercises/blob/master/11-observability/README.md)]
    - Health
    - Metrics
    - Logs
    - Traces
    - OpenTelemetry

## Running the slides

The slides are made using the [go present syntax](https://pkg.go.dev/golang.org/x/tools/present) which has its own tooling.

The project maintains an up-to-date deployed version of the slides on [lectures.course-go.dev](lectures.course-go.dev).

The only disadvantage is that this version does not allow you to run the Go code presented on the slides for security reasons. 
To run the Go code, you must run the slides locally.

### Running locally

#### Using Docker

You can use an already pre-built container image with lecture slides.

```
docker run -p "3999:3999" \
    ghcr.io/course-go/lectures:latest \
    present -http=:3999
```

#### Using present directly

Or you can run the slides locally using the Go present CLI.

##### Clone the repository

```
git clone git@github.com:course-go/lectures.git
```

##### Installation

[Install Go](https://go.dev/doc/install) if you do not have it yet.

Install the CLI present tool using Go install:

```
go install golang.org/x/tools/cmd/present@latest
```

This installs the **present** executable into your 
**$GOPATH/bin** directory if **GOPATH** is set or the 
**$HOME/go/bin** directory otherwise.

##### Usage

Make sure that the directory with the **present** executable is in your $PATH. 

After that, running the present tool is simple:

```
present
```

Alternatively, you can specify on which address the server should listen:

```
present http=:8080
```

A webserver is run on the specified address. You can now preview the slides using your favourite web browser.

## Attribution

Some of the lecture slides are based on the [GoCourse](https://github.com/RedHatOfficial/GoCourse) 
project developed under [Red Hat](https://github.com/RedHatOfficial). The GoCourse is licensed under 
[CC BY-SA 4.0 DEED license](https://creativecommons.org/licenses/by-sa/4.0/deed.en), which this 
repository respects and therefore shares.

