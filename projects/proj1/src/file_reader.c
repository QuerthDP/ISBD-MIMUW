/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#include <fcntl.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/stat.h>
#include "mmap_load.h"
#include "read_load.h"
#include "crc64.h"

void run_read_load(int fd, off_t filesize) {
    read_sequential(fd, filesize);
    read_random(fd, filesize);
}

void run_mmap_load(int fd, off_t filesize) {
    mmap_sequential(fd, filesize);
    mmap_random(fd, filesize);
}

int main(int argc, char* argv[]) {
    if (argc != 2) {
        perror("Usage: ./file_reader <file_path>\n");
        return 1;
    }
    const char* file_path = argv[1];

    int fd = open(file_path, O_RDONLY);
    if (fd < 0) {
        perror("open");
        return 1;
    }

    struct stat st;
    if (fstat(fd, &st) < 0) {
        perror("fstat");
        close(fd);
        return 1;
    }
    if (st.st_size == 0) {
        printf("File is empty\n");
        close(fd);
        return 0;
    }
    off_t filesize = st.st_size;

    crc64_init();

    run_read_load(fd, filesize);
    run_mmap_load(fd, filesize);

    close(fd);
    return 0;
}
