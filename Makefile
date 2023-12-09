run:
	docker compose -f deploy/docker-compose.yml up

docker:
	docker-compose -f deploy/docker-compose.yml down
	docker-compose -f deploy/docker-compose.yml build --no-cache
	docker-compose -f deploy/docker-compose.yml up

clear:
	docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

swagger:
	swag init -g ./cmd/main.go
	swag fmt
