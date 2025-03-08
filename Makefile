# Host Makefile.

include Makefile.include
-include documentation/Makefile

ifeq ($(PROFILES),)
PROFILES := 'pmm'
endif

env-up: 							## Start devcontainer
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose up -d --wait --wait-timeout 100

env-up-rebuild: env-update-image	## Rebuild and start devcontainer. Useful for custom $PMM_SERVER_IMAGE
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose up --build -d

env-update-image:					## Pull latest dev image
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose pull

env-compose-up: env-update-image
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose up --detach --renew-anon-volumes --remove-orphans --wait --wait-timeout 100

env-devcontainer:
	docker exec -it --workdir=/root/go/src/github.com/percona/pmm pmm-server .devcontainer/setup.py

env-down:							## Stop devcontainer
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose down --remove-orphans

env-remove:
	COMPOSE_PROFILES=$(PROFILES) \
	docker compose down --volumes --remove-orphans

TARGET ?= _bash

env:								## Run `make TARGET` in devcontainer (`make env TARGET=help`); TARGET defaults to bash
	COMPOSE_PROFILES=$(PROFILES) \
	docker exec -it --workdir=/root/go/src/github.com/percona/pmm pmm-server make $(TARGET)

rotate-encryption: 							## Rotate encryption key
	go run ./encryption-rotation/main.go

include build/Makefile.clickhouse

###############################################################################
# Packaging Targets: Build completo (tarball, RPM, DEB e release)
###############################################################################

# Nome do pacote
PKG_NAME         := pmm
# Obtém a versão sanitizada a partir do Git (usa somente a tag principal, sem sufixos)
VERSION          := $(shell git describe --tags --abbrev=0 | sed 's/^v//')
# Distribuição – ex.: "el7" para CentOS/OL7 (ajuste conforme necessário)
DIST             := el7
# Arquitetura – geralmente "noarch" para binários Go
ARCH             := noarch
# Diretório base do rpmbuild (normalmente ~/rpmbuild)
RPM_TOPDIR       := $(HOME)/rpmbuild
# Diretório do submódulo que contém os arquivos de empacotamento (usaremos branch v3)
SUBMODULE_DIR    := pmm-submodules
# Define a branch do submódulo que deve ser usada
SUBMODULE_BRANCH := v3
# Caminho do arquivo spec para empacotamento (conforme a estrutura atual do pmm-submodules)
SPEC_FILE        := $(SUBMODULE_DIR)/sources/pmm/src/github.com/percona/pmm/build/packages/rpm/server/SPECS/pmm-managed.spec

.PHONY: update-submodules build-tarball build-rpm build-deb pkg-release pkg-all

## Atualiza os submódulos; se o diretório não existir, adiciona-o usando a branch definida
update-submodules:
	@echo "Verificando submódulos..."
	if [ ! -d "$(SUBMODULE_DIR)" ]; then \
	  echo "Submódulo '$(SUBMODULE_DIR)' não encontrado, adicionando na branch $(SUBMODULE_BRANCH)..."; \
	  git submodule add -b $(SUBMODULE_BRANCH) https://github.com/Percona-Lab/pmm-submodules.git $(SUBMODULE_DIR); \
	fi
	git submodule update --init --recursive

## Cria o tarball do código-fonte usando git archive (que ignora arquivos não rastreados e o .gitignore)
build-tarball: update-submodules
	@echo "Criando tarball do código-fonte..."
	@mkdir -p $(RPM_TOPDIR)/SOURCES
	# Gera o tarball com prefixo "pmm-<versão>/" (apenas arquivos rastreados)
	git archive --format=tar --prefix=$(PKG_NAME)-$(VERSION)/ HEAD > $(PKG_NAME)-$(VERSION).tar
	@mv $(PKG_NAME)-$(VERSION).tar $(RPM_TOPDIR)/SOURCES/
	@echo "Tarball criado em $(RPM_TOPDIR)/SOURCES/$(PKG_NAME)-$(VERSION).tar"

## Configura o ambiente de rpmbuild, copia o arquivo .spec e executa o rpmbuild
build-rpm: build-tarball
	@echo "Configurando ambiente de rpmbuild..."
	@mkdir -p $(RPM_TOPDIR)/{BUILD,RPMS,SPECS,SRPMS}
	@if [ ! -f "$(SPEC_FILE)" ]; then \
	  echo "Erro: Arquivo spec não encontrado em $(SPEC_FILE)"; \
	  exit 1; \
	else \
	  echo "Arquivo spec encontrado: $(SPEC_FILE)"; \
	fi
	@echo "Copiando arquivo .spec..."
	@cp $(SPEC_FILE) $(RPM_TOPDIR)/SPECS/
	@echo "Executando rpmbuild..."
	rpmbuild -ba $(RPM_TOPDIR)/SPECS/$(notdir $(SPEC_FILE)) \
	  --define "version $(VERSION)" \
	  --define "dist $(DIST)" \
	  --define "_topdir $(RPM_TOPDIR)"
	@echo "Pacote RPM gerado em $(RPM_TOPDIR)/RPMS/$(ARCH)/"

## Converte o pacote RPM gerado para um pacote DEB utilizando alien
build-deb: build-rpm
	@echo "Instalando alien (se necessário)..."
	sudo apt-get update && sudo apt-get install -y alien
	@echo "Convertendo RPM para DEB..."
	sudo alien -k --verbose --to-deb $(RPM_TOPDIR)/RPMS/$(ARCH)/*.rpm
	@echo "Pacote DEB gerado."

## Consolida os artefatos (tarball, RPM e DEB) em um diretório "release"
pkg-release: build-deb
	@echo "Organizando artefatos para release..."
	@mkdir -p release
	@cp $(RPM_TOPDIR)/SOURCES/$(PKG_NAME)-$(VERSION).tar release/
	@cp $(RPM_TOPDIR)/RPMS/$(ARCH)/*.rpm release/
	@cp $(PKG_NAME)*.deb release/
	@echo "Artefatos consolidados no diretório 'release'."

## Executa todas as etapas de empacotamento
pkg-all: update-submodules build-tarball build-rpm build-deb pkg-release
	@echo "Build completo de pacotes finalizado."
