# Name of the executable
BINARY_NAME=bootdev_pokedex

# Go build flags
LDFLAGS=-s -w

# Main file location
MAIN=cmd/main.go

# Default target
all: build run

# Build the executable
build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) $(MAIN)

# Clean up the executable
clean:
	rm -f $(BINARY_NAME)

# Run the application
run: 
	./$(BINARY_NAME)

.PHONY: all build clean run