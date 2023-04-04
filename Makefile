server:
	docker-compose up -d --build --remove-orphans ping-server mongodb_container mongo-express

server_logs:
	docker logs -f ping-server

prepare_db:
	docker-compose up -d --build --remove-orphans mongo-seed

clean:
	docker compose down
