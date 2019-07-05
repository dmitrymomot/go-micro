.PHONY: apply delete build

apply: build
	@kubectl apply \
		-f nats/example/pub/deploy.yml \
		-f nats/example/pub/service.yml

delete:
	@kubectl delete \
		-f nats/example/pub/service.yml \
		-f nats/example/pub/deploy.yml

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o nats/example/pub/.build/app nats/example/pub/main.go
	docker build -f nats/example/pub/Dockerfile -t nats-pub .

%:
	@: