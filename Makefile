# Host Makefile.

include Makefile.include
-include documentation/Makefile
-include build/Makefile.clickhouse

ifeq ($(PROFILES),)
PROFILES := 'pmm'
endif

env-up: 							## Start devcontainer
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose -f ./docker-compose.dev.yml up -d --wait --wait-timeout 100

env-up-rebuild: env-update-image	## Rebuild and start devcontainer. Useful for custom $PMM_SERVER_IMAGE
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose -f ./docker-compose.dev.yml up --build -d

env-update-image:					## Pull latest dev image
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose -f ./docker-compose.dev.yml pull

env-compose-up: env-update-image
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose up --detach --renew-anon-volumes --remove-orphans --wait --wait-timeout 100

env-devcontainer:
	docker exec -it --workdir=/root/go/src/github.com/percona/pmm --user root pmm-server python .devcontainer/setup.py

env-down:							## Stop devcontainer
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose -f ./docker-compose.dev.yml down --remove-orphans

env-remove:
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose -f ./docker-compose.dev.yml down --volumes --remove-orphans

TARGET ?= _bash
DEVCONTAINER_NAME ?= $(or $(shell docker ps --filter 'name=^/devcontainer-cosmos-dev-1$$' --format '{{.Names}}'),pmm-server)
DEVCONTAINER_WORKDIR ?= $(if $(filter devcontainer-cosmos-dev-1,$(DEVCONTAINER_NAME)),/workspace/cosmos-main/apps/pmm,/root/go/src/github.com/percona/pmm)
DEVCONTAINER_USER_FLAG ?= $(if $(filter devcontainer-cosmos-dev-1,$(DEVCONTAINER_NAME)),--user codespace,)

env:								## Run `make TARGET` in devcontainer (`make env TARGET=help`); TARGET defaults to bash
	COMPOSE_PROFILES=$(PROFILES) \
	docker exec -it $(DEVCONTAINER_USER_FLAG) --workdir=$(DEVCONTAINER_WORKDIR) $(DEVCONTAINER_NAME) make $(TARGET)

env-root:								## Run `make TARGET` in devcontainer (`make env-root TARGET=help`); TARGET defaults to bash
	COMPOSE_PROFILES=$(PROFILES) \
	docker exec -it --workdir=$(DEVCONTAINER_WORKDIR) --user root $(DEVCONTAINER_NAME) make $(TARGET)

rotate-encryption: 							## Rotate encryption key
	go run ./encryption-rotation/main.go

include Makefile.datacosmos
