# backend
run :
	make build make upd
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

# db
MYSQL_USER := user
MYSQL_PASSWORD := password
MYSQL_DB := db

migrate-up:
	docker-compose exec -T backend sh ../db/migrations/scripts/migrate-up.sh 
mysql:
	docker-compose exec mysql mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) -D$(MYSQL_DB)
seed:
	docker-compose exec mysql mysql -u$(MYSQL_USER) -p$(MYSQL_PASSWORD) -D$(MYSQL_DB) -e "source /seed/seed.sql"
