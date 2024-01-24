hello:
	@echo "Hello"

build:
	@go build -o bin/main cmd/main.go

templ_gen:
	templ generate

style_gen:
	npm run tailwind

full_build: templ_gen build style_gen

run:
	air
