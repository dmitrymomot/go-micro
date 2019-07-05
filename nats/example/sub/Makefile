.PHONY: apply delete build

apply: build
	@kubectl apply \
		-f nats/example/sub/deploy.yml \
		-f nats/example/sub/service.yml

delete:
	@kubectl delete \
		-f nats/example/sub/service.yml \
		-f nats/example/sub/deploy.yml

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o nats/example/sub/.build/app nats/example/sub/main.go
	docker build -f nats/example/sub/Dockerfile -t nats-sub .

%:
	@: