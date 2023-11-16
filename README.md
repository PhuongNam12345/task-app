# Gin-gonic framework

## Getting started

Prerequisites
Go: any one of the three latest major releases (we test it with these).

### Getting Gin

-   With Go module support, simply add the following import

import "github.com/gin-gonic/gin"

to your code, and then go [build|run|test] will automatically fetch the necessary dependencies.
Otherwise, run the following Go command to install the gin package:

-   go get -u github.com/gin-gonic/gin

### Build Project

-   go run build

### Run Project

-   go run main.go

### Project Directory Structure

### 1 **`configsfb/`**

Contains file connecting database

### 1 **`handle/`**

Contains are functions handles in app

### 1. **`middleware/`**

This directory contains file middleware use in app

### **`model/`**

Contains file query database

### **`security/`**

Contains the security of the app such as jwt,..

### `main.go`

File code run app

### `go.mod`

File manages install package automatic when go project

### `go.sum`

File manages version package in app
