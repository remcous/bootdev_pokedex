# Name of the executable
BINARY_NAME=bootdev_pokedex

# Go build flags
LDFLAGS=-s -w

# Default target
all: build

# Build the executable
build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) main.go

# Clean up the executable
clean:
	rm -f $(BINARY_NAME)

# Run the application
run: 
	go run main.go

.PHONY: all build clean run