#!/bin/bash

echo "Setting up the environment..."

# Go installation
if ! [ -x "$(command -v go)" ]; then
    echo "Installing Go..."
    wget https://go.dev/dl/go1.19.5.linux-amd64.tar.gz -O go.tar.gz
    sudo tar -C /usr/local -xzf go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo "Go installed successfully."
else
    echo "Go is already installed."
fi

# MongoDB installation
if ! [ -x "$(command -v mongod)" ]; then
    echo "Installing MongoDB..."
    if [ "$(uname)" == "Darwin" ]; then
        brew tap mongodb/brew
        brew install mongodb-community@6.0
    elif [ "$(uname)" == "Linux" ]; then
        sudo apt-get install -y mongodb
    fi
    echo "MongoDB installed successfully."
else
    echo "MongoDB is already installed."
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

echo "Setup complete. Ready to build and run the project!"