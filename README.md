# fibonacci-service-sample

This service calculates fibonacci value for given number and displays result as html. 

![screenshot](./media/screenshot.png?raw=true)

The required number passed as query param to server address: `/?n=123`

The number must be not less than 0 and not greater than 200,000 (for safety reasons). Calculation of 200,000th fibonacci number takes around 100ms on average laptop. If needed this number can be increased in `internal/handler/handler.go`.

## Requirements 

In order to run service the following is required:
* Golang (tested with 1.16)
* Make (in order to use Makefile commands)

## Commands

* `make run` to run service locally on 0.0.0.0:1123
* `make test` to run tests
* `make lint` to run linters (golangci-lint)
* `make check` to run both tests and linters
* `make build` to build service binary