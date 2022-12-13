# xpress

## How to run:

### Go + Container for Database
- Dependencies : `podman/docker(container)`, `Golang`
- Write `.env` file using `.env.example`
- Run 
``` bash
podman run -itd --rm \
	--name xpress-db \
	--userns keep-id \
	--env-file ./.env \
	--volume ./db:/var/lib/mysql \
	--volume ./init-db.sql:/docker-entrypoint-initdb.d/init-db.sql \
	-p 3306:3306 \
	mariadb
```
- Run `go run .`

---

### Binary + Container for Database
- Dependencies : `podman/docker(container)`
- Download Archive from releases and extract
- Write `.env` file using `.env.example` in the folder
- Run 
``` bash
podman run -itd --rm \
	--name xpress-db \
	--userns keep-id \
	--env-file ./.env \
	--volume ./db:/var/lib/mysql \
	--volume ./init-db.sql:/docker-entrypoint-initdb.d/init-db.sql \
	-p 3306:3306 \
	mariadb
```
- Run `./xpress`

--- 

### Run both website & database in Container
- Soon

---

***Disclaimer : This is personal project. Do not use in production.***

***Beware : The License is AGPLv3***