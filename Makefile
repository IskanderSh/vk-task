run:
	go run cmd/main.go --config=./config/local.yaml

gen-server:
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