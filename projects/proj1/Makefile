# Root Makefile to drive CMake build

BUILD_DIR := build

MAKEFLAGS += --no-print-directory

.PHONY: all clean run

all:
	@mkdir -p $(BUILD_DIR)
	@cmake -S . -B $(BUILD_DIR)
	@$(MAKE) -C $(BUILD_DIR)

run: all
	@./$(BUILD_DIR)/file_reader

clean:
	@$(MAKE) -C $(BUILD_DIR) clean || true
	@rm -rf $(BUILD_DIR)
