# Mock Server
[![Go](https://github.com/teootoledo/go-mock-server/actions/workflows/go.yml/badge.svg)](https://github.com/teootoledo/go-mock-server/actions/workflows/go.yml)

## Description
This is a mock server that can be used to mock responses from a real server.
Sometimes the real service and credentials are not available yet, so to test certain flows this is a good alternative.

## How to use
### Docker
The easiest way to use this mock server is by using docker.
To run the mock server, run the following command:
```bash
docker run -p 8080:8080 -d --name mock-server mock-server:latest
```
Or with docker-compose:
```bash
docker-compose up -d mock-server
```

### Local
To run the mock server locally, you need to have go installed.
Then run the following command:
```bash
go run main.go
```

## Considerations
This is a very simple mock server, it does not persist any data.
So if the server is restarted, all the data will be lost.

## Documentation (Open API Specification)
The documentation of the API can be found [here](http://localhost:8080/v1/docs/index.html) if you run it locally or with docker.

## Contributing
If you want to contribute to this project, create a pull request with your changes.

## Contact
If you have any questions, feel free to contact me vía [email](mailto:teootoledo@gmail.com) or [LinkedIn](https://www.linkedin.com/in/teootoledo/).
