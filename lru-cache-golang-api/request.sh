#!/bin/bash

# Add a key-value pair
curl -X POST http://localhost:8080/cache -H "Content-Type: application/json" -d '{"key":"example_key", "value":"example_value"}'

# Retrieve a value by key
curl http://localhost:8080/cache/example_key