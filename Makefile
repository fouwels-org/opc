# SPDX-FileCopyrightText: 2021 Kaelan Thijs Fouwels <kaelan.thijs@fouwels.com>
#
# SPDX-License-Identifier: MIT

COMPOSE=docker compose
BUILDFILE=compose.yml

build: Dockerfile
	$(COMPOSE) -f $(BUILDFILE) build
push:
	$(COMPOSE) -f $(BUILDFILE) push
up:
	$(COMPOSE) -f $(BUILDFILE) up
up-d:
	$(COMPOSE) -f $(BUILDFILE) up-d
down:
	$(COMPOSE) -f $(BUILDFILE) down