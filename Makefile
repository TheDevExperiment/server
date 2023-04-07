server:
	docker-compose up -d --build --remove-orphans --scale mongo-seed=0

server_logs:
	docker logs -f ping-server

prepare_db:
	docker-compose up -d --build --remove-orphans mongo-seed

clean:
	docker compose down
