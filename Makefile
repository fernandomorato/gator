DB_URL := postgres://postgres:postgres@localhost:5432/gator

up:
	cd sql/schema && goose postgres "$(DB_URL)" up

down:
	cd sql/schema && goose postgres "$(DB_URL)" down
