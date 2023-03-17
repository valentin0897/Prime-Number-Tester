## Task
Build API to check if a list of given numbers are primes. API should only have a single POST endpoint to accept the request. The request body must be a slice of integers, otherwise, return an error. If the request is valid, return a slice of booleans either the number was prime or not.


## How to launch code
1. `Install Go`: If you haven't already, you'll need to install Go on your system. You can download the latest version of Go from the official website: https://golang.org/dl/

2. `Download the code`: Download the code from the repository

3. `Install dependencies`: The API depends on several third-party packages, including Echo, a popular web framework for Go. You can install these dependencies by running the following command from the root directory of the project:

```
go mod download
```

4. `Launch the server`: You can launch the server by running the following command from the root directory of the project:

```
go run ./cmd/main.go
```

5. `Test the API`: You can test the API by sending a POST request to the / endpoint with a JSON body containing an array of integers. For example:

```
curl -X POST -H "Content-Type: application/json" -d '[2, 3, 5, 7, 10]' http://localhost:5000/
```

## Structure

The application is structured in a layered architecture pattern. The layers are divided into the `API layer` and the `business logic` layer.

1. **api/http/handlers**: This folder contains the `API layer` of your application. The handlers in this folder are responsible for handling incoming requests and sending responses back to clients.

2. **rest**: This folder contains the `server` configuration where the routes and endpoints for the API are defined, along with the methods that are responsible for starting and shutting down the `server`.

3. **model**: This folder contains the `business logic` layer of your application.

## Extensibility

I used the **strategy pattern** for implementing different algorithms, which makes it simpler if the client chooses to switch algorithms or add new endpoints for different algorithms.

## Technical simplifications

1. I didn't create a configuration file to avoid excessive complexity, so I hard-coded port value for starting server
2. The error message only displays the index of the first invalid input, as displaying all the indexes would not be practical when dealing with a large amount of data.