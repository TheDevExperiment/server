server:
	docker-compose up --build --remove-orphans

clean:
	docker rm docker-ping || true
