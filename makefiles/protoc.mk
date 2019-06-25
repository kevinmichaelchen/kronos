.PHONY: protoc
protoc:
	protoc -I internal/pb internal/pb/*.proto --go_out=plugins=grpc:./internal/pb

.PHONY: install-proto
install-proto:
	brew install protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go