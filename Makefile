add-network:
	docker network create go-network
up:
	docker-compose up
build:
	docker-compose build 
down:
	docker-compose down
sh:
	docker-compose exec backend sh
mysql:
	docker compose exec mysql mysql -uroot -ppasswordroot