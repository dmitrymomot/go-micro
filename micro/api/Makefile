.PHONY: apply delete

apply:
	@kubectl apply \
		-f micro/api/deploy.yml \
		-f micro/api/service.yml

delete:
	@kubectl delete \
		-f micro/api/deploy.yml \
		-f micro/api/service.yml

%:
	@: