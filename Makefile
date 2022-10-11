up:
	db-up server-up

down:
	db-down server-down

db-up:
	podman run -itd --rm \
	--name xpress-db \
	--userns keep-id \
	--env-file ./website/.env \
	--volume ./db:/var/lib/mysql \
	--volume ./init-db.sql:/docker-entrypoint-initdb.d/init-db.sql \
	-p 3306:3306 \
	mariadb

db-down:
	podman stop xpress-db

db-clean:
	rm -rf db/*

server-up:

server-down:

clean:
	rm -rf db/*
	clear
