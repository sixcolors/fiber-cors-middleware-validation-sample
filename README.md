# Fiber CORS Middleware Validation Sample

This sample is used to validate the Fiber CORS middleware functionality with real clients, such as a web browser.

## Why this sample?

The Fiber CORS middleware is a powerful tool to manage the Cross-Origin Resource Sharing (CORS) policy for your Fiber server. However, it can be challenging to validate the middleware's functionality without real clients, such as a web browser.

This sample provides a simple Fiber server that serves a static web page and includes route handlers for the CORS middleware validation. The web page will make requests to the Fiber server to validate the CORS middleware is correctly handling cross-origin requests.

## Prerequisites

- [Go](https://golang.org/dl/)
- [Node.js](https://nodejs.org/en/download/)
- [NPM](https://www.npmjs.com/get-npm)

## Installation

**Install the Fiber CORS Middleware Validation Sample**:

```bash
git clone https://github.com/sixcolors/fiber-cors-middleware-validation-sample.git
cd fiber-cors-middleware-validation-sample
```

## Known Issues

Fiber v2.52.3 has a [bug](https://github.com/gofiber/fiber/issues/2936) that prevents the CORS middleware from working correctly.

This repository can be used to validate the bug and test the fix.

Replace the Fiber v2.52.3 module with a local copy of Fiber with the fix [#2937](https://github.com/gofiber/fiber/pull/2937).

**Checkout the fix**:

```bash
cd ..
git clone https://github.com/sixcolors/fiber.git
cd fiber
git checkout 2936-cors-error
```

To replace the Fiber module with a local copy, update the go.mod file to point to the local Fiber repository.

```go.mod
replace github.com/gofiber/fiber/v2 v2.52.3 => ../Documents/GitHub/fiber
```

## How to run

**Start the Servers**:

```bash
sh start.sh
```

The start.sh script will start two servers:

1. The first server is a simple Fiber server that will serve a static web page at `http://localhost:3000/`. It also includes route handlers for the CORS middleware validation.
    - `http://localhost:3000/` [GET]: The main page.
        - This page will make requests to the Fiber server to validate the CORS middleware is not interfering with the same-origin requests.
    - `http://localhost:3000/hello` [GET]: A route that will return a simple String response.
    - `http://localhost:3000/api/:id` [GET]: A route that will return the request body as a response.
    - `http://localhost:3000/api/:id` [POST]: A route that will return the request body as a response.
    - `http://localhost:3000/*` [OPTIONS]: A route that will return 204 No Content response, and will set a `Obvious-Header: Obvious-Value` header. Used to verify the CORS preflight request is not blocking same-origins OPTIONS requests.

2. The second server is a simple express.js HTTP server that will serve a static web page at `http://localhost:8080/`.
    - `http://localhost:8080/` [GET]: The main page.
        - This page will make requests to the Fiber server to validate the CORS middleware is correctly handling cross-origin requests.

**Test the Servers**:

1. Open the browser and navigate to `http://localhost:3000/`.
    - The page should load correctly and make requests to the Fiber server.
    - Press the buttons to make requests to the Fiber server and verify the responses.
2. Open the browser and navigate to `http://localhost:8080/`.
    - The page should load correctly and make requests to the Fiber server.
    - Press the buttons to make requests to the Fiber server and verify the responses.

**Stop the Servers**:
    
```bash
sh stop.sh
```