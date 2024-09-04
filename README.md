### Targets:

```bash
make tailwind-watch
```

This target watches the ./static/css/input.css file and automatically rebuilds the Tailwind CSS styles whenever changes are detected.

```
make tailwind-build
```

This target minifies the Tailwind CSS styles by running the tailwindcss command.

```
make templ-watch
```

This target watches for changes to \*.templ files and automatically generates them.

```
make templ-generate
```

This target generates templates using the templ command.

```
make dev
```

This target runs the development server using Air, which helps in hot-reloading your Go application during development.

```
make build
```

This target orchestrates the building process by executing the tailwind-build, templ-generate, and go build commands sequentially. It creates the binary output in the ./bin/ directory.
