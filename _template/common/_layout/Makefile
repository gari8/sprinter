init:
	sh scripts/init-local.sh && make build && make up_db && make up
build:
	docker-compose build --no-cache
up:
	docker-compose up api
up_db:
	docker-compose up -d db
run_api:
	docker-compose run api ash
run_db:
	docker-compose run db ash
exec_api:
	docker-compose exec api ash
exec_db:
	docker-compose exec db ash
down:
	docker-compose down
stop:
	docker-compose stop