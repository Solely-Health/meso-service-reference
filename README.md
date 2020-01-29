Based docker image build off of:
https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

This guide follows a technique called multistage building to minimize container size

What is going on here?

`go.mod && go.sum`
These files are a result of the `go init mod` command. 
In Go, there is a feature called modules, which allows our developers to create modules outside of the $GOPATH. Go Modules also help us manage our package management. 

`main.go`
This is the main Go file that is compiled and exicuted.

`/database`
In this directory exists an internal package called 'database', this is a rough mock up of how we will structure our repositories.

