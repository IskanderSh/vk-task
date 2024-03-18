run:
	go run cmd/main.go --config=./config/local.yaml

gen-server:
	mkdir -p internal/generated
	swagger generate server \
	-f ./swagger-api/swagger.yml \
	-t ./internal/generated \
	-C ./swagger-templates/default-server.yml \
	--template-dir ./swagger-templates/templates \
	--name server

gen-server-windows:
	docker run --rm -it quay.io/goswagger/swagger generate server \
	-f ./swagger-api/swagger.yml \
    -t ./internal/generated \
	-C ./swagger-templates/default-server.yml \
	--template-dir ./swagger-templates/templates \
 	--name server

migrate:
	goose -dir "./migrations" postgres "host=postgres port=5432 user=postgres password=password" up

docker-up:
	docker-compose up -d