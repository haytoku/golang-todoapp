include .env
export

export PROJECT_ROOT := $(shell pwd)

env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down

env-cleanup:
	@read -p "Are you sure you want to remove the environment? [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down && \
		rm -rf $(PROJECT_ROOT)/out/pgdata && \
		echo "Environment removed successfully."; \
	else \
		echo "Environment removal canceled."; \
	fi

env-port-forward:
	docker compose up -d port-forward

env-port-forward-stop:
	docker compose down port-forward

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Error: Please provide a sequence number using the 'seq' variable."; \
		exit 1; \
	fi
	docker compose run --rm todoapp-postgres-migrate create \
		-ext sql \
		-dir /migrations \
		-seq \
		"$(seq)"

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down

migrate-action:
	docker compose run --rm todoapp-postgres-migrate  \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable"	\
		"${action}"
	

