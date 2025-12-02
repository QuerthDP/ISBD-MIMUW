/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#include "timer.h"

#include <stdio.h>

void start_timer(struct Timer* timer) {
    if (!timer->running) {
        clock_gettime(CLOCK_MONOTONIC, &timer->start);
        timer->running = true;
    }
}

void stop_timer(struct Timer* timer) {
    if (timer->running) {
        clock_gettime(CLOCK_MONOTONIC, &timer->end);

        time_t sec_diff = timer->end.tv_sec - timer->start.tv_sec;
        long usec_diff = (timer->end.tv_nsec - timer->start.tv_nsec) / 1000;
        if (usec_diff < 0) {
            sec_diff -= 1;
            usec_diff += 1000000L;
        }
        timer->total_sec += sec_diff;
        timer->total_usec += usec_diff;
        // Normalize microseconds
        if (timer->total_usec >= 1000000L) {
            timer->total_sec += timer->total_usec / 1000000L;
            timer->total_usec %= 1000000L;
        }
        timer->running = false;
    }
}

void reset_timer(struct Timer* timer) {
    timer->total_sec = 0;
    timer->total_usec = 0;
    timer->running = false;
}

void read_timer(const struct Timer* timer) {
    printf("Elapsed time: %ld.%06ld seconds\n", timer->total_sec, timer->total_usec);
}
