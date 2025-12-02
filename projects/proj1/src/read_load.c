/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#include "read_load.h"

#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include "crc64.h"
#include "timer.h"

void read_sequential(int fd, off_t filesize) {
    printf("\nReading file using read() system call sequentially\n");

    unsigned char* buffer = malloc(BLOCK_SIZE);
    if (!buffer) {
        perror("malloc");
        exit(1);
    }

    uint64_t crc = 0;
    ssize_t bytes;
    size_t total_read = 0;
    struct Timer timer;
    reset_timer(&timer);

    start_timer(&timer);
    while ((bytes = read(fd, buffer, BLOCK_SIZE)) > 0) {
        stop_timer(&timer);
        crc = crc64_update(crc, buffer, bytes);
        start_timer(&timer);
        total_read += bytes;
    }
    stop_timer(&timer);
    read_timer(&timer);

    printf("Total bytes read: %zu\n", total_read);
    if (total_read != filesize) {
        fprintf(stderr, "Error: Expected filesize %zu but read %zu bytes\n", filesize, total_read);
        free(buffer);
        close(fd);
        exit(1);
    }
    print_crc64(crc);

    free(buffer);
}

void read_random(int fd, off_t filesize) {
    printf("\nReading file using read() system call randomly\n");

    unsigned char* buffer = malloc(BLOCK_SIZE);
    if (!buffer) {
        perror("malloc");
        exit(1);
    }

    uint64_t crc = 0;
    size_t num_blocks = (filesize + BLOCK_SIZE - 1) / BLOCK_SIZE;
    size_t total_read = 0;
    struct Timer timer;
    reset_timer(&timer);

    start_timer(&timer);
    for (size_t i = 0; i < num_blocks; ++i) {
        size_t block_index = (i % 2 == 0) ? i : (num_blocks - i - 1);
        off_t offset = (off_t)block_index * BLOCK_SIZE;
        if (offset >= filesize)
            continue;
        lseek(fd, offset, SEEK_SET);
        ssize_t bytes = read(fd, buffer, BLOCK_SIZE);
        if (bytes <= 0)
            break;
        stop_timer(&timer);
        crc = crc64_update(crc, buffer, bytes);
        start_timer(&timer);
        total_read += bytes;
    }
    stop_timer(&timer);
    read_timer(&timer);

    printf("Total bytes read: %zu\n", total_read);
    if (total_read != filesize) {
        fprintf(stderr, "Error: Expected filesize %zu but read %zu bytes\n", filesize, total_read);
        free(buffer);
        close(fd);
        exit(1);
    }
    print_crc64(crc);

    free(buffer);
}
