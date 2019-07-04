.PHONY: apply delete

apply:
	@kubectl apply \
		-f micro/web/deploy.yml \
		-f micro/web/service.yml

delete:
	@kubectl delete \
		-f micro/web/deploy.yml \
		-f micro/web/service.yml

%:
	@: