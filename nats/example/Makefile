.PHONY: apply delete build

apply: build
	@kubectl apply \
		-f test/deploy.yml \
		-f test/service.yml

delete:
	@kubectl delete \
		-f test/service.yml \
		-f test/deploy.yml

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o test/.build/app test/main.go
	docker build -f test/Dockerfile -t nats-test .

%:
	@: