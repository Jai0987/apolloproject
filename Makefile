# Variables
APP_NAME = ApolloProject
BUILD_DIR = build
SRC_DIR = .
MAIN_FILE = main.go

# Commands
GO_CMD = go
MKDIR_CMD = mkdir -p
RM_CMD = rm -rf

# Targets
.PHONY: build
build:
	@echo "[APOLLO] Building $(APP_NAME)..."
	$(MKDIR_CMD) $(BUILD_DIR)
	$(GO_CMD) build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)/$(MAIN_FILE)
	@echo "[APOLLO] Build complete. Binary is located in $(BUILD_DIR)/$(APP_NAME)"

.PHONY: run
run:
	@echo "[APOLLO] Running $(APP_NAME)..."
	$(GO_CMD) run $(MAIN_FILE)

.PHONY: clean
clean:
	@echo "[APOLLO] Cleaning up build artifacts..."
	$(RM_CMD) $(BUILD_DIR)
	@echo "[APOLLO] Clean complete."

.PHONY: setup
setup:
	@echo "[APOLLO] Setting up dependencies..."
	./setup.sh

.PHONY: populate
populate:
	@echo "[APOLLO] Populating database with sample data..."
	$(GO_CMD) run populate_db.go
