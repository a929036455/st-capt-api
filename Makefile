.PHONY: api
# generate api proto
common:
	cd api && kratos proto client ./common/v1/common.proto
auth:
	cd api && kratos proto client ./auth/v1/auth.proto
	cd api && kratos proto client ./auth/v1/auth_error.proto
sms:
	cd api && kratos proto client ./sms/v1/sms.proto
	cd api && kratos proto client ./sms/v1/sms_error.proto
user:
	cd api && kratos proto client ./user/v1/user.proto
wx:
	cd api && kratos proto client ./wx/v1/base.proto
	cd api && kratos proto client ./wx/v1/deal.proto
	cd api && kratos proto client ./wx/v1/deal_error.proto
	cd api && kratos proto client ./wx/v1/message.proto

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api
	#make config
	#make generate

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
