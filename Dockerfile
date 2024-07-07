FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=10203150112Zaa!

COPY ./ /docker-entrypoint-initdb.d/