/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#include <unistd.h>
#include "timer.h"

int main() {
    struct Timer timer;
    reset_timer(&timer);

    start_timer(&timer);

    sleep(5);

    stop_timer(&timer);
    read_timer(&timer);

    return 0;
}
