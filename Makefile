down:
	@sudo docker-compose down

up:
	@sudo docker-compose up --force-recreate --no-deps -d

start: down up

run:
	@cd runner && go run main.go
