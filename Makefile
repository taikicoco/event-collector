build:
	docker-compose build
build-nc:
	docker-compose build --no-cache
up:
	docker-compose up
upd:
	docker-compose up -d
down:
	docker-compose down
rm-volumes:
	docker-compose down --volumes
log:
	docker-compose logs -f
sh:
	docker-compose exec backend sh
mysql:
	docker-compose exec mysql mysql -uroot -ppasswordroot -Ddb

# migrate
migrate-up:
	docker-compose exec -T backend sh ./scripts/migrate-up.sh
migrate-down:
	docker-compose exec -T backend sh ./scripts/migrate-down.sh
