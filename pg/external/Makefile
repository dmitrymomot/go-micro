.PHONY: apply delete

apply:
	@kubectl apply \
		-f pg/local/config.yml \
		-f pg/local/storage.yml \
		-f pg/local/deploy.yml \
		-f pg/external/service.yml

delete:
	@kubectl delete \
		-f pg/external/service.yml \
		-f pg/local/deploy.yml \
		-f pg/local/storage.yml \
		-f pg/local/config.yml

%:
	@: