# Tip
Run the following to create a virtual environment and install the dependencies:
- `export GOPATH=$PWD` in the project root
- `cd src/devops_coding_sessions`
- `go mod init devops_coding_sessions` 
- `go mod tidy`


# Structure
- `pkg/`: golang packages (dependencies)
- `src/`: our code
  - `devops_coding_sessions/`: module name
    - `domain/`: have the entites responsible to the most important business rules
    - `application/`: have the application business rules and will use entities to fulfill the use cases
    - `adapters/`: have the adapters to transform the external data to a format that works to the internal layers (business rule layers)
    - `infrastructure/`: the external tools needed to make the platform work (frameworks, databases, etc)

# Flow
Web (infrastructure) -> Controller (adapters) -> Use Case (application) -> Entity (domain)

# Test
Run the project with `go run .` and the command below in terminal or postman:
```
curl --location --request POST 'http://localhost:8080/input' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Martin"
}'
```