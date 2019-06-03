.PHONY: code-quality
code-quality:
	@$(MAKE) vet
	@$(MAKE) tidy
	@$(MAKE) imports
	@$(MAKE) fmt
	@$(MAKE) test
	@$(MAKE) lint

.PHONY: test
test:
	@echo "${GREEN}✓ Running tests${NC}\n"
	@cd ${SRC_DIR} ; \
	env GO111MODULE=${GO111MODULE} \
	  BIGTABLE_EMULATOR_HOST=${BIGTABLE_EMULATOR_HOST} \
	  NUCLEUS_LOG_WITH_FIELDS=${NUCLEUS_LOG_WITH_FIELDS} \
	  LOG_FORMAT=${LOG_FORMAT} \
	  LOG_LEVEL=${LOG_LEVEL} \
	  go test -count=1 ./... && \
	  cd _tests && \
	env GO111MODULE=${GO111MODULE} \
	  BIGTABLE_EMULATOR_HOST=${BIGTABLE_EMULATOR_HOST} \
	  NUCLEUS_LOG_WITH_FIELDS=${NUCLEUS_LOG_WITH_FIELDS} \
	  LOG_FORMAT=${LOG_FORMAT} \
	  LOG_LEVEL=${LOG_LEVEL} \
	  go test -count=1 ./...

.PHONY: testv
testv:
	@echo "${GREEN}✓ Running tests${NC}\n"
	@cd ${SRC_DIR} ; env GO111MODULE=${GO111MODULE} \
	  BIGTABLE_EMULATOR_HOST=${BIGTABLE_EMULATOR_HOST} \
	  NUCLEUS_LOG_WITH_FIELDS=${NUCLEUS_LOG_WITH_FIELDS} \
	  LOG_FORMAT=${LOG_FORMAT} \
	  LOG_LEVEL=${LOG_LEVEL} \
	  go test -v -count=1 ./... && \
	  cd _tests && \
	env GO111MODULE=${GO111MODULE} \
	  BIGTABLE_EMULATOR_HOST=${BIGTABLE_EMULATOR_HOST} \
	  NUCLEUS_LOG_WITH_FIELDS=${NUCLEUS_LOG_WITH_FIELDS} \
	  LOG_FORMAT=${LOG_FORMAT} \
	  LOG_LEVEL=${LOG_LEVEL} \
	  go test -v -count=1 ./...

.PHONY: tidy
tidy:
	@echo "${GREEN}✓ Pruning dependencies${NC}\n"
	@cd ${SRC_DIR} ; env GO111MODULE=${GO111MODULE} go mod tidy

.PHONY: imports
imports:
	@echo "${GREEN}✓ Cleaning up imports${NC}\n"
	@echo "${BLUE}✓ This may take a few seconds...${NC}\n"
	@cd ${SRC_DIR} ; env GO111MODULE=${GO111MODULE} goimports -w .

.PHONY: importsv
importsv:
	@echo "${GREEN}✓ Cleaning up imports${NC}\n"
	@echo "${BLUE}✓ This may take a few seconds...${NC}\n"
	@cd ${SRC_DIR} ; env GO111MODULE=${GO111MODULE} goimports -v -w .

.PHONY: fmt
fmt:
	@echo "${GREEN}✓ Formatting code${NC}\n"
	@cd ${SRC_DIR} ; env GO111MODULE=${GO111MODULE} gofmt -s -w .

.PHONY: vet
vet:
	@echo "${GREEN}✓ Checking code for correctness${NC}\n"
	@cd ${SRC_DIR} ; env GO111MODULE=${GO111MODULE} go vet ./...

.PHONY: lint
lint:
	@echo "${GREEN}✓ Checking code style${NC}\n"
	@cd ${SRC_DIR} ; env GO111MODULE=${GO111MODULE} golint ./...
