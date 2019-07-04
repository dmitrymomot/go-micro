.PHONY: help up down info

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up: ## Deployment of a service. Usage example: make up micro/api
	@make -f $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))/Makefile apply

down: ## Shut down a service. Usage example: make down micro/api
	@make -f $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))/Makefile delete

info: ## Get kubernetes cluster information
	@kubectl get all

%:
	@: