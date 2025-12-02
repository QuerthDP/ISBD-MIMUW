/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#pragma once

#include <time.h>
#include <stdbool.h>

struct Timer {
    struct timespec start;
    struct timespec end;
    time_t total_sec;
    long total_usec;
    bool running;
};

void start_timer(struct Timer* timer);
void stop_timer(struct Timer* timer);
void reset_timer(struct Timer* timer);
void read_timer(const struct Timer* timer);
