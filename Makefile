.PHONY: shell up production development
.SILENT:

target := $(firstword $(MAKECMDGOALS))

ifeq ($(filter up down reup, $(target)),$(target))
	arg := $(or $(word 2, $(MAKECMDGOALS)), development)
else ifeq ($(target),shell)
	arg := $(word 2, $(MAKECMDGOALS))
endif



up:
	docker compose -f ./deploy/$(arg)/docker-compose.yml up -d

down:
	docker compose -f ./deploy/$(arg)/docker-compose.yml down

reup: down up

shell:
	docker exec -it gotus-$(arg) /bin/sh