.PHONY: build

build:
	docker-compose -f deployments/docker-compose.yml build --no-cache
up:
	docker-compose -f deployments/docker-compose.yml up api
up_db:
	docker-compose -f deployments/docker-compose.yml up -d db
run_api:
	docker-compose -f deployments/docker-compose.yml run api ash
run_db:
	docker-compose -f deployments/docker-compose.yml run db ash