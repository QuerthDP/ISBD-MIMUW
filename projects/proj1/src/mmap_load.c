/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#include "mmap_load.h"

#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/mman.h>
#include "crc64.h"
#include "timer.h"

void mmap_sequential(int fd, off_t filesize) {
    printf("\nReading file using mmap() system call sequentially\n");

    uint64_t crc = 0;
    ssize_t bytes;
    size_t total_read = 0;
    struct Timer timer;
    reset_timer(&timer);

    start_timer(&timer);
    unsigned char* data = mmap(NULL, filesize, PROT_READ, MAP_PRIVATE, fd, 0);
    if (data == MAP_FAILED) {
        perror("mmap");
        exit(1);
    }

    for (off_t offset = 0; offset < filesize; offset += BLOCK_SIZE) {
        size_t size = (offset + BLOCK_SIZE <= filesize) ? BLOCK_SIZE : filesize - offset;
        stop_timer(&timer);
        crc = crc64_update(crc, data + offset, size);
        start_timer(&timer);
        total_read += size;
    }
    stop_timer(&timer);
    read_timer(&timer);

    printf("Total bytes read: %zu\n", total_read);
    if (total_read != filesize) {
        fprintf(stderr, "Error: Expected filesize %zu but read %zu bytes\n", filesize, total_read);
        munmap(data, filesize);
        close(fd);
        exit(1);
    }
    print_crc64(crc);

    munmap(data, filesize);
}

void mmap_random(int fd, off_t filesize) {
    printf("\nReading file using mmap() system call randomly\n");

    uint64_t crc = 0;
    size_t num_blocks = (filesize + BLOCK_SIZE - 1) / BLOCK_SIZE;
    size_t total_read = 0;
    struct Timer timer;
    reset_timer(&timer);

    start_timer(&timer);
    unsigned char* data = mmap(NULL, filesize, PROT_READ, MAP_PRIVATE, fd, 0);
    if (data == MAP_FAILED) {
        perror("mmap");
        exit(1);
    }

    for (size_t i = 0; i < num_blocks; ++i) {
        size_t block_index = (i % 2 == 0) ? i : (num_blocks - i - 1);
        off_t offset = (off_t)block_index * BLOCK_SIZE;
        if (offset >= filesize)
            continue;
        size_t size = (offset + BLOCK_SIZE <= filesize) ? BLOCK_SIZE : filesize - offset;
        stop_timer(&timer);
        crc = crc64_update(crc, data + offset, size);
        start_timer(&timer);
        total_read += size;
    }
    stop_timer(&timer);
    read_timer(&timer);

    printf("Total bytes read: %zu\n", total_read);
    if (total_read != filesize) {
        fprintf(stderr, "Error: Expected filesize %zu but read %zu bytes\n", filesize, total_read);
        munmap(data, filesize);
        close(fd);
        exit(1);
    }
    print_crc64(crc);

    munmap(data, filesize);
}
