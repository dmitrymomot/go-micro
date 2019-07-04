.PHONY: apply delete

apply:
	@kubectl apply \
		-f pg/local/config.yml \
		-f pg/local/storage.yml \
		-f pg/local/deploy.yml \
		-f pg/internal/service.yml

delete:
	@kubectl delete \
		-f pg/internal/service.yml \
		-f pg/local/deploy.yml \
		-f pg/local/storage.yml \
		-f pg/local/config.yml

%:
	@: