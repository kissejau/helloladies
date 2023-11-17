run:
	docker compose -f deploy/docker-compose.yml up

docker:
	docker compose -f deploy/docker-compose.yml down
	docker compose -f deploy/docker-compose.yml build --no-cache
	docker compose -f deploy/docker-compose.yml up