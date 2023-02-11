
# SPlit Management API

The Split Management API is a RESTful API that allows users to sign up, login, and manage groups. The API implements various endpoints for users to perform actions such as creating a group, adding people to the group, splitting expenses, and viewing payment status.

## Prerequisites


- Go programming language
- Go libraries such as gin and gorm
- A database system, such as PostgreSQL



## Installation

- clone git repo
```bash
 git clone https://github.com/shuhaib-kv/Split-Payment.git
```
    
- initialize missing modules necessary to build the current module's packages and dependencies
```bash
go mod tidy
```
- make an env file and input the databasedetails and port
```bash
touch .env
```


- run the main fiel

```bash
go run main.go
```

## Documentation

[POSTMAN](https://documenter.getpostman.com/view/23876360/2s935uFKw5)

