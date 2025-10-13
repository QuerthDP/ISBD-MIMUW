/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#pragma once

#ifndef BLOCK_SIZE
#define BLOCK_SIZE (8 * 1024 * 1024) // 8 MB
#endif

#include <sys/types.h>

void read_sequential(int fd, off_t filesize);
void read_random(int fd, off_t filesize);
