.PHONY: protoc
protoc:
	protoc -I $(SRC_DIR)/pb $(SRC_DIR)/pb/*.proto --go_out=plugins=grpc:./$(SRC_DIR)/pb

.PHONY: install-proto
install-proto:
	brew install protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go