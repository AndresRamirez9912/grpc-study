FROM postgres:16

COPY /src/DB/up.sql /docker-entrypoint-initdb.d/1.sql

EXPOSE 5432

CMD [ "postgres" ]
