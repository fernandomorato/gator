DB_URL = ./.app.db

up:
	cd sql/schema && goose sqlite3 "$(DB_URL)" up

down:
	cd sql/schema && goose sqlite "$(DB_URL)" down

reset:
	make down && make up

gen:
	sqlc generate

format:
	gofmt -l -w .
