### run test
- go test -v
- go test -v <filename>
- go test -v -run <test>

### test coverage
- go test -v -coverprofile coverage.out
- go tool cover -func=coverage.out
- go tool cover -html=coverage.out
- go tool cover -html=coverage.out -o coverage.html