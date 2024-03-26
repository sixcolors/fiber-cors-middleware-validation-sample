#!/bin/bash

# Check if the Go server is already running
if [ -f go_server.pid ] && kill -0 $(cat go_server.pid) 2>/dev/null; then
    echo "Go server is already running"
else
    # Install Go dependencies
    go mod download

    # Build the Go binary
    go build -o server.out main.go

    # Start the Go server and store its PID
    ./server.out &
    echo $! > go_server.pid
fi

# Navigate to the simple-dev-server directory
cd simple-dev-server

# Check if the Node.js server is already running
if [ -f node_server.pid ] && kill -0 $(cat node_server.pid) 2>/dev/null; then
    echo "Node.js server is already running"
else
    # Install dependencies
    npm install

    # Start the Node.js server and store its PID
    npm start &
    echo $! > node_server.pid
fi

# Navigate back to the base directory
cd ..

# Wait for the servers to start
sleep 5

# Open http://localhost:3000 in the default browser
open http://localhost:3000