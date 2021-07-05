# SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
#
# SPDX-License-Identifier: MIT

FROM alpine:3.14.0 as build

# Install tools required for project
RUN apk add go cmake make build-base

ENV VERSION_OPEN62541=1.2.2

WORKDIR /build
RUN wget -q -O open62541.tar.gz https://github.com/open62541/open62541/archive/refs/tags/v$VERSION_OPEN62541.tar.gz && tar -xzf open62541.tar.gz
RUN mkdir -p open62541-${VERSION_OPEN62541}/out

RUN apk add python3
RUN export CC=/usr/bin/clang && export CXX=/usr/bin/clang++
RUN cd open62541-${VERSION_OPEN62541}/out && cmake -DUA_ENABLE_AMALGAMATION=ON -DUA_LOGLEVEL=300 \
    ..
RUN cd open62541-${VERSION_OPEN62541}/out && make && make install

ENV GOBIN=/build/out
COPY . ./go
RUN cd go && go install .

FROM alpine:3.14.0
COPY --from=build /build/out /build/out

CMD ["/build/out/opc"]
