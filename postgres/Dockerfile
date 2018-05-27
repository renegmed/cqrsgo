FROM postgres:10.3 
# SQL files inside /docker-entrypoint-initdb.d directory will be executed in alphabetical order.
COPY up.sql /docker-entrypoint-initdb.d/1.sql
CMD ["postgres"]
