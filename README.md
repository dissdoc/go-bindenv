# BindEnv
Minimalistic extension for any configuration readers

## Overview

This is a simple extension tool for any configuration readers. It keeps careful for your code and just does the following:
* bind to your selected struct
* find environment variables
* set values from env variables to your configuration struct

## Installation

To install the package run
```sh
go get github.com/dissdoc/go-bindenv
```

## Usage

The extension is oriented to be simple in use and explicitness

The main idea is to use your exist code and decouple logic the code and configuration file. You should edit only config files without recompile your code

There are just several actions you can do with this tool:
* read environment variables
* read certain struct and find special values
* concat special values with environment variables

## Read configuration

Set special values to the config file
```yaml
properties:
    host: ${SERVER_HOST}
    port: ${SERVER_PORT}
```
Define a structure
```go
type ApiServer struct {
    Host string
    Port string
}
```
Bind the extension
```go
import (
    bindenv "github.com/dissdoc/go-bindenv"
)

// ... some code for initialize instances
// ... read config yaml with any modules
config := &ApiServer{}
// ... some more code

bindenv.BindEnv(config)

// ... then let's use it
fmt.Println(config.Host, config.Port)
```

## Blog Posts

[Why we need to use it](https://habr.com/ru/articles/776284/)