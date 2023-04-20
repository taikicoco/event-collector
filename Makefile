add-network:
	docker network create rest-go-network
up:
	docker-compose up
down:
	docker-compose down
sh:
	docker-compose exec backend sh
mysql:
	docker compose exec mysql mysql -uroot -ppasswordroot -Ddb

migrate-up:
	docker-compose exec -T backend sh ./scripts/migrate-up.sh

