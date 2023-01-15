# Zwitter Backend

This is the backend for the Zwitter project, a Twitter clone. The backend is built using Go and provides a RESTful API for the Zwitter platform using the [Zwitter API specification](https://github.com/zwitter/zwitter-api).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go version 1.15+
- PostgreSQL 12+
- Zwitter API specification

### Installing

- Clone the repository
- Run `go mod download` to download the dependencies
- Create a `.env` file in the root of the project and set the following environment variables:
  - `DB_URL`: The URL of the PostgreSQL database
  - `API_HOST`: IP and port of the API server
  - `SECRET`: A secret key for JWT authentication
- Run `build/build.sh` to build the project
- Run `./zwitter-server` to start the server

### Running the tests

- Run `go test ./...` to run the tests

### API documentation

API documentation can be found in the zwitter-api specification.

### Deployment

- Build the binary with `build/build.sh`
- Deploy the binary on a server
- Set the environment variables

### Built With

- [Go](https://golang.org/) - The programming language used
- [PostgreSQL](https://www.postgresql.org/) - The database used
- [Zwitter API specification](https://github.com/zwitter/zwitter-api) - The API specification used

### Contributing

Please read [CONTRIBUTING.md](https://github.com/zwitter/zwitter/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.
