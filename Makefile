hello:
	@echo "Hello"

build:
	@go build -o bin/main cmd/main.go

templ_gen:
	templ generate
style_gen:
	npx tailwind -i 'tailwind.css' -o 'public/styles.css' --minify

run:
	air
