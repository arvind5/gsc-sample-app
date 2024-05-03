# Copyright (C) 2024 Intel Corporation
# All rights reserved.
# SPDX-License-Identifier: BSD-3-Clause

FROM golang:1.21.6-bullseye AS builder
WORKDIR /app
COPY . .
RUN go build -o gscexample

FROM ubuntu:22.04 AS final
WORKDIR /
COPY --from=builder /app/gscexample ./
CMD ["/gscexample"]
