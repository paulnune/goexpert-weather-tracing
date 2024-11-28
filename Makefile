.PHONY: lab


# Variables
CEP=01214000

up:
	docker compose up -d;

down:
	docker compose down;

restart:
	docker compose restart;

clean:
	@if [ "$(shell docker ps -a -q)" != "" ]; then \
		sudo docker rm -f $(shell docker ps -a -q); \
	else \
		echo "No containers to remove."; \
	fi
	@if [ "$(shell docker images -q)" != "" ]; then \
		sudo docker rmi -f $(shell docker images -q); \
	else \
		echo "No images to remove."; \
	fi
	@if [ "$(shell docker volume ls -q)" != "" ]; then \
		sudo docker volume prune -f; \
	else \
		echo "No volumes to remove."; \
	fi
	sudo docker system prune -af

svc-a:
	@sleep 3s ;
	curl -X POST -d '{"cep": "$(CEP)"}' http://localhost:8080
	@echo '\n' ;

svc-b:	
	@sleep 3s ;
	curl http://localhost:8081/weather?cep=$(CEP)
	@sleep 3s ;
	@echo '\n' ;

all: up svc-a svc-b

apps: svc-a svc-b
