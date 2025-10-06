/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#include "timer.h"

#include <stdio.h>

void start_timer(struct timespec* ts) {
    clock_gettime(CLOCK_REALTIME, ts);
}

void read_timer(struct timespec* ts_start, struct timespec* ts_end) {
    clock_gettime(CLOCK_REALTIME, ts_end);

    time_t sec_diff = ts_end->tv_sec - ts_start->tv_sec;
    long usec_diff = (ts_end->tv_nsec - ts_start->tv_nsec) / 1000;
    if (usec_diff < 0) {
        sec_diff -= 1;
        usec_diff += 1000000L;
    }

    printf("Elapsed time: %ld.%06ld seconds\n", sec_diff, usec_diff);
}
