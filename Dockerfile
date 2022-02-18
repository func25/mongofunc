# this file create a mongo replicaset environment for transaction tests
FROM mongo:4.4.7
RUN echo "rs.initiate();" > /docker-entrypoint-initdb.d/replica-init.js
CMD [ "--replSet", "rs" ]