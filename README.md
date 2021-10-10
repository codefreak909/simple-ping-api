# simple-ping-api
This is a simple ping API to demonstrate the basics of API.

But this is a sample exercise to demonstrate tests in Go.

## How to run tests

- Run go test, this should run all the test in the current directory.

## How to check test coverage

- Run go test -coverprofile=a.out , This will run the tests and generate a coverprofile.
- Run go tool cover -html=a.out, this will open the test coverage as a html in your browser.

Note: Ensure you run these commands in command prompt or terminal. It doesn't work in powershell.