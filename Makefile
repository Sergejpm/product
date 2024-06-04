LOCAL_BIN := $(CURDIR)/bin
BUF_VENDOR := $(CURDIR)/buf/vendor
BUF_GOOGLEAPI_PB := $(BUF_VENDOR)/google/api
BUF_OPENAPIV2_PB := $(BUF_VENDOR)/protoc-gen-openapiv2/options

LDFLAGS:=-X 'github.com/sergejpm/product/app.Name=product'
SERVICE_NAME:=product

GO_VERSION:=$(shell go version)
GO_VERSION_SHORT:=$(shell echo $(GO_VERSION) | sed -E 's/.* go(.*) .*/\1/g')

ifneq ("1.22","$(shell printf "$(GO_VERSION_SHORT)\n1.22" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.22. Found: $(GO_VERSION_SHORT))
endif

generate: .dir-deps .bin-deps .pb-deps 
	$(info generating api stubs...)
	buf generate
	$(info tidying up modules...)
	go mod tidy

.dir-deps:
	$(info creating buf directories...)
	mkdir -p $(BUF_VENDOR)
	mkdir -p $(LOCAL_BIN)
	mkdir -p $(BUF_GOOGLEAPI_PB)
	mkdir -p $(BUF_OPENAPIV2_PB)

.bin-deps:
	$(info installing binary dependecies...)
	#go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	#go get github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_BIN) go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.3.0

.pb-deps:
	$(info installing pb dependencies...)
	curl -o $(BUF_GOOGLEAPI_PB)/annotations.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
	curl -o $(BUF_GOOGLEAPI_PB)/http.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
	curl -o $(BUF_OPENAPIV2_PB)/annotations.proto https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/annotations.proto
	curl -o $(BUF_OPENAPIV2_PB)/openapiv2.proto https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/openapiv2.proto

.PHONY: build-api
build-api:
	$(info Building service...)
	go build -ldflags "$(LDFLAGS)" -o $(LOCAL_BIN) ./cmd/api

.PHONY: run-api
run-api:
	$(info Running...)
	$(BUILD_ENVPARMS) go run -ldflags "$(LDFLAGS)" ./cmd/api
