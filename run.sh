#!/bin/bash

echo "Starting Apollo Project setup..."
# Step 1: Setup dependencies
./setup.sh

# Step 2: Populate the database
make populate

# Step 3: Build the project
make build

# Step 4: Run the project
make run
