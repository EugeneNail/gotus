.PHONY: shell up production development
.SILENT:

target := $(firstword $(MAKECMDGOALS))

ifeq ($(filter up down reup, $(target)),$(target))
	arg := $(or $(word 2, $(MAKECMDGOALS)), development)
else ifeq ($(filter shell build, $(target)),$(target))
	arg := $(word 2, $(MAKECMDGOALS))
endif



up:
	docker compose -f ./deploy/$(arg)/docker-compose.yml up -d

down:
	docker compose -f ./deploy/$(arg)/docker-compose.yml down

reup: down up

shell:
	docker exec -it gotus-$(arg) /bin/sh

build:
ifeq ($(arg),ui)
	cd ui && BUILD_PATH=../web npm run build && touch ../web/.gitkeep
else ifeq ($(arg),api)
	@echo "not implemented"
else ifeq ($(arg),)
	make build ui
	@echo "not implemented"
endif
