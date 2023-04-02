server:
	docker-compose up -d --build --remove-orphans

server_logs:
	docker logs -f ping-server

clean:
	docker compose down
