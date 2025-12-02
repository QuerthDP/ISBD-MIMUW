# Minimum CMake version required
cmake_minimum_required(VERSION 3.10)

# Project name and language
project(file_reader C)

# Set the C standard
set(CMAKE_C_STANDARD 17)
set(CMAKE_C_STANDARD_REQUIRED True)

# Gather all .c files from the src directory
file(GLOB SOURCES "src/*.c")

# Create the executable
add_executable(${PROJECT_NAME} ${SOURCES})

# Include the src directory for header files
target_include_directories(${PROJECT_NAME} PRIVATE src)
