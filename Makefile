add-network:
	docker network create go-api
build:
	docker-compose build
up:
	docker-compose up
build:
	docker-compose build 
down:
	docker-compose down
sh:
	docker-compose exec backend-go-api sh
mysql:
	docker compose exec mysql mysql -uroot -ppasswordroot