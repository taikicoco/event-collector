add-network:
	docker network create glc_network
build:
	docker-compose build
build-nc:
	docker-compose build --no-cache
up:
	docker-compose up
down:
	docker-compose down
sh:
	docker-compose exec go-log-collector sh
mysql:
	docker compose exec mysql mysql -uroot -ppasswordroot -Ddb
migrate-up:
	docker-compose exec -T go-log-collector sh ./scripts/migrate-up.sh
migrate-down:
	docker-compose exec -T go-log-collector sh ./scripts/migrate-down.sh
