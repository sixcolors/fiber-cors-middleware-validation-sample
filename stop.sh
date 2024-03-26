#!/bin/bash

# Check if the Go server is running
if [ -f go_server.pid ] && kill -0 $(cat go_server.pid) 2>/dev/null; then
    # Stop the Go server
    kill $(cat go_server.pid)
    rm go_server.pid
    rm server.out
else
    echo "Go server is not running"
fi

# Navigate to the simple-dev-server directory
cd simple-dev-server

# Check if the Node.js server is running
if [ -f node_server.pid ] && kill -0 $(cat node_server.pid) 2>/dev/null; then
    # Stop the Node.js server
    kill $(cat node_server.pid)
    rm node_server.pid
else
    echo "Node.js server is not running"
fi

# Navigate back to the base directory
cd ..