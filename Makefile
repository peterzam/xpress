up:
	db-up server-up

down:
	db-down server-down

db-up:
	podman run -it --rm \
	--name xpress-db \
	--userns keep-id \
	--env-file .env \
	--volume ./db:/var/lib/mysql \
	--volume ./init-db.sql:/docker-entrypoint-initdb.d/mysql-init.sql \
	-p 3306:3306 \
	mariadb

db-down:
	podman stop xpress-db

server-up:

server-down: