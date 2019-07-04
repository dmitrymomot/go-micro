.PHONY: help up down info

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up: ## Deployment of a service. Usage example: make up micro/api
	@kubectl apply \
		-f $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))/deploy.yml \
		-f $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))/service.yml

down: ## Shut down a service. Usage example: make down micro/api
	@kubectl delete \
		-f $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))/service.yml \
		-f $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))/deploy.yml

info: ## Get kubernetes cluster information
	@kubectl get all

%:
	@: