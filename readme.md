# Go Project Setup

## Install Go  
Download and install Go from [go.dev](https://go.dev/).  

## Check Go Version  
Verify your installation by running:  
```sh
go version
```
Example output:  
```sh
go version go1.24.1 windows/amd64
```

## Initialize Go Module  
Run one of the following commands to initialize a Go module:  
```sh
go mod init projectName
```
or, using a GitHub repository:  
```sh
go mod init github.com/username/reponame
```

## Run the Go Project  
Use the following command to run your Go application:  
```sh
go run restapi/main.go
```
