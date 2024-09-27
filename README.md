## Requirements

- [Go 1.19+](https://golang.org/dl/)
- [GoFiber v2](https://gofiber.io/)
- [Air](https://github.com/cosmtrek/air) (for live reloading)

## Getting Started

### Clone the repository
```
git clone <repo=url>
cd <project-name>
```

### Install Dependencies
```
go mod tidy
```

### Install Air for Live Reloading
Air is a development tool that automatically reloads your Go server when file changes are detected.

To install `air` globally, run:
```
go install github.com/cosmtrek/air@latest
```

Configuration for Air

Creates an .air.toml configuration file in the root directory:
```
air init
```

## Run the Application
To start the application with live reloading, simply run:
```
air
```
Air will watch your files for changes and automatically reload the server. The API will be available at http://localhost:3000.

Alternatively, you can run the app without Air:

````
go run main.go
````

### API Endpoints
- **GET /classes**: Lists all classes.
- **POST /classes**: Creates a new class.
- **GET /booking**: Lists all bookings.
- **POST /booking**: Creates a booking.

## Run tests
To run tests, simply run:
```
go test -v ./tests/...
```