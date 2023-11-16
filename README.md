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

### **`configsfb/`**

Contains file connecting database

### **`handle/`**

Contains are functions handles in app

### **`middleware/`**

This directory contains file middleware use in app

### **`model/`**

Contains file is define struct data, logic database

### **`security/`**

Contains the security of the app such as jwt,..

### **`routers/`**

Contains routers defines HTTP routes and handles requests and responses from the application's various endpoints.

### `main.go`

File code run app

### `go.mod`

File manages install package automatic when go project

### `go.sum`

File manages version package in app
