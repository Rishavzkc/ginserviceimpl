#!/usr/bin/env make

.PHONY: docker-up
docker-up:
	docker run -d \
        --name mysql-service \
        -e MYSQL_ROOT_PASSWORD=Quest1234 \
        -e MYSQL_PASSWORD=Password123 \
        -e MYSQL_DATABASE=company \
        -e MYSQL_USER=admin \
        -p 3306:3306 \
        mysql:latest
	@echo "Wait for 20 seconds for initialization"; sleep 20

.PHONY: docker-down
docker-down:
	docker kill mysql-service
	docker rm -f mysql-service

.PHONY: run
run:
	go run main.go

.PHONY: sample
sample:
	@echo "- Should insert record {\"1\" \"Apple Inc.\" \"California\"}:"
	@curl localhost:8080/company/ \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"id": "1","name": "Apple Inc.","location": "California"}'

	@echo "\n\n- Should insert record {\"2\" \"Tesla\" \"California}:"
	@ curl localhost:8080/company/ \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"id": "2","name": "Tesla","location": "California"}'

	@echo "\n\n- Should fail insertion due to duplication of ID in {\"2\" \"Tesla\" \"California}:"
	@ curl localhost:8080/company/ \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"id": "2","name": "Tesla","location": "California"}'

	@echo "\n\n- Should retrieve all records:"
	@ curl localhost:8080/company/

	@echo "\n\n- Should retrieve record for ID = 1:"
	@ curl localhost:8080/company/1

	@echo "\n\n- Should retrieve record for ID = 2:"
	@ curl localhost:8080/company/2

	@echo "\n\n- Should fail retrieval of record with ID = 3:"
	@ curl localhost:8080/company/3

	@echo "\n\n- Should update record with ID = 2 as {\"2\" \"Tesla\" \"Texas}:"
	@ curl localhost:8080/company/ \
        --header "Content-Type: application/json" \
        --request "PUT" \
        --data '{"id": "2","name": "Tesla","location": "Texas"}'

	@echo "\n\n- Should delete record with ID = 1:"
	@ curl localhost:8080/company/1 \
		--request "DELETE"

	@echo "\n\n- Should delete record with ID = 2:"
	@ curl localhost:8080/company/2 \
		--request "DELETE"

	@echo "\n\n- Should fail to delete record with ID = 3 as not found:"
	@ curl localhost:8080/company/3 \
		--request "DELETE"