.PHONY: apply delete

apply:
	@kubectl apply -f redis/storage.yaml -f redis/redis-master.yaml
	@sleep 10
	@kubectl apply -f redis/redis-sentinel-service.yaml
	@kubectl apply -f redis/redis-controller.yaml
	@kubectl apply -f redis/redis-sentinel-controller.yaml
	@sleep 10
	@kubectl scale rc redis --replicas=3
	@kubectl scale rc redis-sentinel --replicas=3
	@sleep 10
	@kubectl delete pods redis-master

delete:
	@kubectl delete \
		-f redis/redis-sentinel-controller.yaml \
		-f redis/redis-controller.yaml \
		-f redis/redis-sentinel-service.yaml

%:
	@: