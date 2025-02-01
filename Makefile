new-migration:
	migrate create -ext sql -dir pkg/sql/migrations -seq ${name}

mysql-docker:
	docker run --name mysql-container -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_USER=admin -e MYSQL_PASSWORD=secret -e MYSQL_DATABASE=dating_app -d mysql:latest

redis-docker:
	docker run --name redis-container -p 6379:6379 -d redis:latest