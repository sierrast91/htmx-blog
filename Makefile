run:build
	./bin/htmx-blog-app
build:
	go build -o ./bin/htmx-blog-app ./cmd/htmx-blog-app/main.go
test:
	go test ./... -v

tailwind:
	npx tailwindcss -i global.css -o ./assets/style.css
tw:
	npx tailwindcss -i global.css -o ./assets/style.css --watch

