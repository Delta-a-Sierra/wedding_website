templ-watch:
	templ generate --watch --proxy="http://localhost:3000"

tailwind-watch:
	tailwindcss -i ./internal/adapters/presentation/htmx/static/css/input.css -o ./internal/adapters/presentation/htmx/static/css/style.css --watch

dev:
	air
