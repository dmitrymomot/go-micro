
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/jwt/jwt.proto

.PHONY: build
build: proto
	go clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o jwtapi main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t jwtapi:latest

.PHONY: apply delete

apply:
	@kubectl apply \
		-f auth/jwtapi/deploy.yml \
		-f auth/jwtapi/service.yml

delete:
	@kubectl delete \
		-f auth/jwtapi/deploy.yml \
		-f auth/jwtapi/service.yml

%:
	@: