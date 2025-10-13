# Root Makefile to drive CMake build

BUILD_DIR := build

MAKEFLAGS += --silent --no-print-directory

CFLAGS := -Wall -Wextra -Werror -Wpedantic -Wshadow -Wformat=2 -Wfloat-equal -Wconversion -Wsign-conversion

.PHONY: all clean run

all:
	@mkdir -p $(BUILD_DIR)
	@cmake -S . -B $(BUILD_DIR)
	@$(MAKE) -C $(BUILD_DIR)

run: all generate-files
	@for file in $(filter-out $@,$(MAKECMDGOALS)); do \
			echo "Running ./$(BUILD_DIR)/file_reader $$file"; \
			./$(BUILD_DIR)/file_reader $$file; \
		done

generate-files:
	@mkdir -p test_files
	@python3 generate_file.py test_files/file_small 1024    # 1KB
	@python3 generate_file.py test_files/file_medium 102400 # 100KB
	@python3 generate_file.py test_files/file_large 1048576 # 1MB

clean:
	@$(MAKE) -C $(BUILD_DIR) clean || true
	@rm -rf $(BUILD_DIR)
	@rm -rf test_files
