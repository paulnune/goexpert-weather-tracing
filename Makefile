.PHONY: lab play

# Variables
CEP=01214000
CAST_FILE_RHEL=.assets/weather-tracing-rhel.cast
CAST_FILE_UBUNTU=.assets/weather-tracing-ubuntu.cast

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
	@sleep 10s ;
	curl -X POST -d '{"cep": "$(CEP)"}' http://localhost:8080
	@echo '\n' ;

svc-b:	
	@sleep 10s ;
	curl http://localhost:8081/weather?cep=$(CEP)
	@sleep 10s ;
	@echo '\n' ;

all: up svc-a svc-b

apps: svc-a svc-b

play_rhel:
	@if ! command -v asciinema >/dev/null 2>&1; then \
		echo "Installing asciinema..."; \
		if [ -f /etc/os-release ]; then \
			source /etc/os-release; \
			case $$ID in \
				ubuntu|debian) \
					sudo apt update && sudo apt install -y asciinema; \
					;; \
				fedora|rhel|centos|rocky|almalinux) \
					sudo dnf install -y asciinema; \
					;; \
				*) \
					echo "Unsupported OS: $$ID"; \
					exit 1; \
					;; \
			esac; \
		else \
			echo "Cannot detect OS. Please install asciinema manually."; \
			exit 1; \
		fi \
	else \
		echo "asciinema is already installed."; \
	fi; \
	if [ -f "$(CAST_FILE_RHEL)" ]; then \
		echo "Playing asciinema cast file: $(CAST_FILE_RHEL)"; \
		asciinema play $(CAST_FILE_RHEL); \
	else \
		echo "Cast file not found: $(CAST_FILE_RHEL)"; \
		exit 1; \
	fi

play_ubuntu:
	@if ! command -v asciinema >/dev/null 2>&1; then \
		echo "Installing asciinema..."; \
		if [ -f /etc/os-release ]; then \
			source /etc/os-release; \
			case $$ID in \
				ubuntu|debian) \
					sudo apt update && sudo apt install -y asciinema; \
					;; \
				fedora|rhel|centos|rocky|almalinux) \
					sudo dnf install -y asciinema; \
					;; \
				*) \
					echo "Unsupported OS: $$ID"; \
					exit 1; \
					;; \
			esac; \
		else \
			echo "Cannot detect OS. Please install asciinema manually."; \
			exit 1; \
		fi \
	else \
		echo "asciinema is already installed."; \
	fi; \
	if [ -f "$(CAST_FILE_UBUNTU)" ]; then \
		echo "Playing asciinema cast file: $(CAST_FILE_UBUNTU)"; \
		asciinema play $(CAST_FILE_UBUNTU); \
	else \
		echo "Cast file not found: $(CAST_FILE_UBUNTU)"; \
		exit 1; \
	fi
