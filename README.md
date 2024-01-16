# Lectures

This repository contains lecture slides for Course Go.

## Tooling

The slides are made using the [go present syntax](https://pkg.go.dev/golang.org/x/tools/present).
You can run the slides locally using the go present CLI command.


### Installation

You can install the CLI command via go install like so:

```
go install golang.org/x/tools/cmd/present
```

This installs the **present** executable into your **$GOPATH/bin** directory.

### Usage

Running the present tool is simple:

```
present
```

A webserver is run on localhost where you can preview the slides using your favourite web browser.