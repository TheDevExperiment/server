server:
	docker-compose up -d --build --remove-orphans --scale mongo-seed=0

server_logs:
	docker logs -f ping-server

prepare_db:
	docker build -t mongo-seed ./mongo-seed
	docker run --network mynet --rm -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=rootpassword mongo-seed

clean:
	docker-compose down
