FROM mongo:6.0.5
COPY posts.json /posts.json
COPY users.json /users.json
CMD mongoimport --host mongodb_container --authenticationDatabase admin -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --db testdb --collection posts --type json --file /posts.json --jsonArray && mongoimport --host mongodb_container --authenticationDatabase admin -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --db testdb --collection users --type json --file /users.json --jsonArray
