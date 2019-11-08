# ubiquitous-fortnight
Ubiquitous-fortnight is a service boilerplate for different REST APIs written in go.

[![Build Status](https://travis-ci.org/jsburckhardt/ubiquitous-fortnight.svg?branch=master)](https://travis-ci.org/jsburckhardt/ubiquitous-fortnight)

## Getting Started
The application contains 3 basic endpoints:

| Enpoint    | purpose                                                         |
| ---------- | --------------------------------------------------------------- |
| /ping      | Keep Alive: confirm the main service is up and running.         |
| /v1        | Keep Alive: confirm v1 api is up and running.                   |
| /v1/status | Keep Alive: confirm v1 api is up and running with state values. |

### Some details about the application:
- Serves on port 8001
- The version in /v1/status is the last line in **metadata** file
- the lastcommitsha in /v1/status is the environment variable **HASH**. The variable is created during build process.


## Run locally the application
- Clone the repository:
```
git clone https://github.com/jsburckhardt/ubiquitous-fortnight.git 
```
- Install dependencies:
```
go get -d -v ./...
go install -v ./...
```
- Build the application. (for this example we call it **theapp**):
```
go build -o ./bin/theapp -v .
```
- Run the application:
```
go run main.go
```
or after build:
```
./bin/theapp
```
- Try the application by using a explorer and going to http://localhost:8001/v1/ or any of the endpoints.
```
curl localhost:8001/ping
curl localhost:8001/v1
curl localhost:8001/v1/status
```
- Run local unit tests:
```
go test
```

## Build a container (docker) and run it locally
To help with repetitive tasks a **Makefile** was created to help with the process.
- Create theapp container with tag latest
```
make dockerbuild
```
- Run the container
```
make dockerrun
```
- Tests:
In the case of using container. The tests are executed during the container creation.

## CICD
For now the CICD is being manage in Github, Docker-Hub and Travis-Ci.

| Utility    | Details                                                                                                                                                                                                                                                                                                 |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Github     |  CI tool. Follows single trunk development being the trunk **Master** branch.                                                                                                                                                                                                                           |
| Travis-CI  | Orchestrator tool. Triggers builds for any code pushed into the repository (Except for **experiemental** branch). The tool at the moment only pushes containers into Docker-Hub if the build is triggered in **master** branch. Also it creates a git tag with date if the container push is successful.|
| Docker-Hub | repository for container images.                                                                                                                                                                                                                                                                        |

## TODO:
- Create login middleware
- Create a centralise login middleware
- Deployment
- Manage versionsing using husky
- Create changelog using husky
- Expose test results
- Use viper to include configurations
- [semantic-release](https://github.com/semantic-release/semantic-release). Semantic-release coupled with [commitizen](https://github.com/commitizen/cz-cli) and [cz-conventional-changelog](https://github.com/commitizen/cz-conventional-changelog) allow us to get and automated semantic versioning