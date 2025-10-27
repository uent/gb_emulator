.PHONY: build run clean test help

# Variables
BINARY_NAME=gb-emulator
MAIN_PATH=./cmd/gb-emulator
BUILD_DIR=./bin

# Target por defecto
all: build

# Compilar el proyecto
build:
	@echo "Compilando $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Compilación completada: $(BUILD_DIR)/$(BINARY_NAME)"

# Ejecutar el programa (requiere especificar ROM como argumento)
# Uso: make run ROM=path/to/rom.gb
run: build
	@if [ -z "$(ROM)" ]; then \
		echo "Error: Especifica una ROM con ROM=path/to/rom.gb"; \
		echo "Ejemplo: make run ROM=tetris.gb"; \
		exit 1; \
	fi
	@echo "Ejecutando $(BINARY_NAME) con ROM: $(ROM)"
	$(BUILD_DIR)/$(BINARY_NAME) $(ROM)

# Ejecutar sin compilar (usar binario existente)
start:
	@if [ -z "$(ROM)" ]; then \
		echo "Error: Especifica una ROM con ROM=path/to/rom.gb"; \
		echo "Ejemplo: make start ROM=tetris.gb"; \
		exit 1; \
	fi
	$(BUILD_DIR)/$(BINARY_NAME) $(ROM)

# Limpiar archivos compilados
clean:
	@echo "Limpiando archivos compilados..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(BINARY_NAME)
	@echo "Limpieza completada"

# Ejecutar tests
test:
	@echo "Ejecutando tests..."
	go test -v ./...

# Ejecutar tests con coverage
test-coverage:
	@echo "Ejecutando tests con coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Reporte de coverage generado: coverage.html"

# Formatear código
fmt:
	@echo "Formateando código..."
	go fmt ./...

# Verificar código con go vet
vet:
	@echo "Verificando código..."
	go vet ./...

# Instalar dependencias
deps:
	@echo "Descargando dependencias..."
	go mod download
	go mod tidy

# Compilación para múltiples plataformas
build-all:
	@echo "Compilando para múltiples plataformas..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	@echo "Compilación completada para todas las plataformas"

# Ayuda
help:
	@echo "Makefile para Game Boy Emulator"
	@echo ""
	@echo "Uso:"
	@echo "  make build              - Compilar el proyecto"
	@echo "  make run ROM=<rom>      - Compilar y ejecutar con una ROM"
	@echo "  make start ROM=<rom>    - Ejecutar sin compilar (usa binario existente)"
	@echo "  make clean              - Limpiar archivos compilados"
	@echo "  make test               - Ejecutar tests"
	@echo "  make test-coverage      - Ejecutar tests con reporte de coverage"
	@echo "  make fmt                - Formatear código"
	@echo "  make vet                - Verificar código con go vet"
	@echo "  make deps               - Descargar dependencias"
	@echo "  make build-all          - Compilar para múltiples plataformas"
	@echo "  make help               - Mostrar esta ayuda"
	@echo ""
	@echo "Ejemplos:"
	@echo "  make build"
	@echo "  make run ROM=tetris.gb"
	@echo "  make start ROM=roms/pokemon.gb"
