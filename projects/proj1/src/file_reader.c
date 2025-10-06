/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#include <unistd.h>
#include "timer.h"

int main() {
    struct timespec start_ts, ts;

    start_timer(&start_ts);

    sleep(5);

    read_timer(&start_ts, &ts);

    return 0;
}