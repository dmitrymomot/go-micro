.PHONY: apply delete

apply:
	@kubectl apply \
		-f nats/operator/prereqs.yml \
		-f nats/operator/deploy.yml

delete:
	@kubectl delete \
		-f nats/operator/deploy.yml \
		-f nats/operator/prereqs.yml

%:
	@: