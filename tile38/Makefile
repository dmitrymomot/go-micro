.PHONY: apply delete

apply:
	@kubectl apply \
		-f tile38/storage.yml \
		-f tile38/deploy.yml \
		-f tile38/service.yml

delete:
	@kubectl delete \
		-f tile38/service.yml \
		-f tile38/deploy.yml \
		-f tile38/storage.yml

%:
	@: