# Tests Directory

This directory contains unit tests for the go-agent codebase.

## Running Tests

To run all tests in this directory:

```bash
go test ./tests -v
```

To run tests with coverage:

```bash
go test ./tests -cover
```

## Test Structure

- `routes_test.go`: Tests for HTTP handlers and structs in the routes package
- `chains_test.go`: Tests for business logic and structs in the chains package

## What's Tested

### Routes Package
- POST /vacation/create endpoint (valid and invalid requests)
- GET /vacation/:id endpoint (found, not found, bad request)
- Struct field validation for:
  - GenerateVacationIdeaRequest
  - GenerateVacationIdeaResponse
  - GetVacationIdeaResponse

### Chains Package
- GetVacationFromDb function (found and not found cases)
- Vacation struct field validation

Note: The GenerateVacationIdeaChange function involves external API calls to Hugging Face and is not unit tested in this suite due to external dependencies. In a production environment, this would be tested with mocks.