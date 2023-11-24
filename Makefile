# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"

# Define the target executable
TARGET = my_app

# Default target: build the executable
all: $(TARGET)

# Rule to build the target executable
$(TARGET): ./app/server.go
	$(GO) build $(GOFLAGS) -o $(TARGET) ./app/

# Clean target: remove the target executable
clean:
	del -f $(TARGET)

# Run target: build and run the target executable
run: $(TARGET)
	./$(TARGET)

# Test target: run Go tests for the project
test:
	$(GO) test ./...
