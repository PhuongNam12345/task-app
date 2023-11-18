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

Contains are functions, logics handles in app

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

## Workflow module task

### create repository

-   clone repository from repository to local code

### Dev code module task

-   create taget-branch is staging, then create new branch from staging branch
-   a feature of app is code on new branch this with name branch feat/namefeat
-   start code module task

### Commit code

-   commit code on branch with commit message description
-   push code to branch

### Pull request to target-branch

-   push pull request branch to staging branch
-   review code, check comment, discuss group, fix code, commit code
-   (`edit and  commit on PR this `)

### Merge pull request

-   check pull request success
-   merge pull request in main branch.
