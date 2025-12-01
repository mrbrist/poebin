CYAN := \033[36m
DIM := \033[2m
RESET := \033[0m
CHECK := ✓
ARROW := →

GO_BIN := $(shell which go)
BUN_BIN := $(shell which bun)
MAIN_GO := cmd/server/main.go

PROJECT_NAME := $(shell basename $(CURDIR))
STATIC_DIR := web/static
gohtml_FILES := $(shell find . -type f -name "*.gohtml")

ALPINE_URL := https://cdn.jsdelivr.net/npm/alpinejs@latest/dist/cdn.min.js

TAILWIND_CONFIG := 'module.exports = {\n\
  content: ["./web/**/*.gohtml"],\n\
  theme: {\n\
    extend: {}\n\
  },\n\
  plugins: []\n\
}'

# ----------------------
# Build / Run Targets
# ----------------------

run:
	@$(GO_BIN) run $(MAIN_GO)

run-compiled:
	@./build/main

compile:
	@$(GO_BIN) build -o build/main $(MAIN_GO)

# ----------------------
# Assets
# ----------------------

download-alpine:
	@printf "\n$(CYAN)Downloading Alpine.js$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@mkdir -p $(STATIC_DIR)/js
	@curl -s $(ALPINE_URL) -o $(STATIC_DIR)/js/alpine.min.js
	@printf "$(CHECK) Alpine.js ready\n"

setup-tailwind:
	@printf "\n$(CYAN)Setting up Tailwind CSS$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(BUN_BIN) install tailwindcss @tailwindcss/cli >/dev/null 2>&1
	@echo $(TAILWIND_CONFIG) > tailwind.config.js
	@echo '@tailwind base;\n@tailwind components;\n@tailwind utilities;' > $(STATIC_DIR)/css/input.css
	@printf "$(CHECK) Tailwind CSS ready\n"

css:
	@printf "\n$(CYAN)Generating CSS$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(BUN_BIN) tailwindcss -i $(STATIC_DIR)/css/input.css -o $(STATIC_DIR)/css/styles.css --minify
	@printf "$(CHECK) CSS generated\n"

# ----------------------
# Serve
# ----------------------

serve:
	@printf "\n$(CYAN)Starting server$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@$(GO_BIN) run $(MAIN_GO) & 
	@$(BUN_BIN) tailwindcss -i $(STATIC_DIR)/css/input.css -o $(STATIC_DIR)/css/styles.css --minify

# ----------------------
# Watch with Hot Reload
# ----------------------

watch:
	@printf "\n$(CYAN)Watching for changes$(RESET)\n"
	@printf "$(DIM)────────────────────────────────────$(RESET)\n"
	@printf "\n$(CYAN)Rebuilding...$(RESET)\n"
	@find web/templates web/static/css -type f \( -name "*.gohtml" -o -name "input.css" \) | entr -r make serve


