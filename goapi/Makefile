CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=github.com/mistercloud/test-todo/goapi/cmd/app

all: format build test lint

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

# #####
# PROTOC Section BEGIN
# #####

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(BINDIR)" protoc
PROTO_PATH=api

# Установка всех необходимых зависимостей
.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)
	GOBIN=$(BINDIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(BINDIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(BINDIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
	GOBIN=$(BINDIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
	GOBIN=$(BINDIR) go install github.com/envoyproxy/protoc-gen-validate@latest && \
	GOBIN=$(BINDIR) go install github.com/bufbuild/buf/cmd/buf@v1.21.0

# Вендоринг внешних proto файлов
.PHONY: .vendor-proto
.vendor-proto: .vendor-rm vendor-proto/google/api vendor-proto/google/protobuf vendor-proto/protoc-gen-openapiv2/options vendor-proto/validate

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

# Устанавливаем proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливаем proto описания protoc-gen-openapiv2/options
vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem


# Устанавливаем proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

# Устанавливаем proto описания validate
vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor-proto/validate
		mv vendor-proto/tmp/validate vendor-proto/
		rm -rf vendor-proto/tmp

# Генерация протофайлов с использованием protoc
PHONY: .protoc-generate
.protoc-generate:
	mkdir -p pkg/${PROTO_PATH}

	mkdir -p api/openapiv2
	$(PROTOC) -I ${PROTO_PATH} -I vendor-proto \
 	--plugin=protoc-gen-go=$(BINDIR)/protoc-gen-go --go_out ./pkg/$(PROTO_PATH) --go_opt paths=source_relative \
 	--plugin=protoc-gen-go-grpc=$(BINDIR)/protoc-gen-go-grpc --go-grpc_out pkg/${PROTO_PATH} --go-grpc_opt paths=source_relative \
 	--plugin=protoc-gen-grpc-gateway=$(BINDIR)/protoc-gen-grpc-gateway --grpc-gateway_out pkg/${PROTO_PATH} --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
 	--plugin=protoc-gen-openapiv2=$(BINDIR)/protoc-gen-openapiv2 --openapiv2_out api/openapiv2 --openapiv2_opt logtostderr=true \
 	--plugin=protoc-gen-validate=$(BINDIR)/protoc-gen-validate --validate_out="lang=go,paths=source_relative:pkg/api" \
	api/todo.proto
	go mod tidy

# Генерация протофайлов с использованием protoc
PHONY: proto-generate
proto-generate: .bin-deps .vendor-proto .protoc-generate

.PHONY: proto-fast-generate
proto-fast-generate: .protoc-generate

# #####
# PROTOC Section END
# #####