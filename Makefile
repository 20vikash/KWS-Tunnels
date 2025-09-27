include .env

up:
	docker compose up -d
	docker compose logs -f

down:
	docker compose down

stop:
	docker compose stop

start:
	docker compose start
	docker compose logs -f

.PHONY: up down stop start
