# SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
#
# SPDX-License-Identifier: MIT

FROM alpine:3.14.0 as build

# Install tools required for project
RUN apk add go

FROM alpine:3.14.0
COPY --from=build /build/out /build/out
