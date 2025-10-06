/*
 * Copyright (c) 2025-present Dawid Pawlik
 *
 * For educational use only by employees and students of MIMUW.
 * See LICENSE file for details.
 */

#pragma once

#include <time.h>

void start_timer(struct timespec* ts);
void read_timer(struct timespec* ts_start, struct timespec* ts_end);
