docker-build:
	docker-compose build
docker-up:
	docker-compose up -d
run:
	(cd golang && go run main.go)
psql:
	docker-compose exec db psql -U clean
run_dev:
	docker-compose exec golang_dev go run main.go
# 以下定義しているけど動かない...
migrate:
	docker-compose exec db_dev /bin/sh
migrate-create:
	docker-compose exec db_dev migrate create -ext sql -dir migrations -seq name
migrate-one-up:
	docker-compose exec db_dev migrate -path migrations -database ${POSTGRESQL_URL} up 1
migrate-one-down:
	docker-compose exec db_dev migrate -path migrations -database ${POSTGRESQL_URL} down 1
migrate-all-up:
	docker-compose exec db_dev migrate -path migrations -database ${POSTGRESQL_URL} up
migrate-all-down:
	docker-compose exec db_dev migrate -path migrations -database ${POSTGRESQL_URL} down -all
prisma:
	docker-compose exec prisma /bin/sh
prisma-migrate:
	docker-compose exec prisma npm run prisma migrate dev --name init
prisma-studio:
	docker-compose exec prisma npm run prisma studio
# proto
cp-ptoro:
	cp ./grpc/googleapis/google/api/annotations.proto ./grpc/google/api/annotations.proto && \
	cp ./grpc/googleapis/google/api/http.proto ./grpc/google/api/http.proto
