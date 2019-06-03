SHELL += -eu

BLUE  := \033[0;34m
GREEN := \033[0;32m
RED   := \033[0;31m
NC    := \033[0m

GO111MODULE := on
SRC_DIR := ./src

# App env vars
BIGTABLE_EMULATOR_HOST ?= 127.0.0.1:8086
BIGTABLE_PROJECT ?= my-project
BIGTABLE_INSTANCE ?= my-instance
NUCLEUS_LOG_WITH_FIELDS ?= false
LOG_FORMAT ?= text
LOG_LEVEL ?= 4

.PHONY: all
all:
	$(MAKE) build
	$(MAKE) run

.PHONY: build
build:
	cd $(SRC_DIR) && \
	  env GO111MODULE=${GO111MODULE} go build -o ../bin/app .

.PHONY: run
run:
	env \
	  GO111MODULE=${GO111MODULE} \
	  BIGTABLE_EMULATOR_HOST=${BIGTABLE_EMULATOR_HOST} \
	  NUCLEUS_LOG_WITH_FIELDS=${NUCLEUS_LOG_WITH_FIELDS} \
	  LOG_FORMAT=${LOG_FORMAT} \
	  LOG_LEVEL=${LOG_LEVEL} \
	  ./bin/app

include makefiles/*.mk