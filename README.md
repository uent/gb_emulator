# Game Boy Emulator

Emulador de Game Boy clásico escrito en Go.

## Descripción

Este proyecto implementa un emulador del sistema Game Boy original (DMG - Dot Matrix Game), permitiendo ejecutar ROMs de juegos clásicos.

## Requisitos

- Go 1.25 o superior

## Estructura del Proyecto

```
gb-emulator/
├── cmd/
│   └── gb-emulator/      # Punto de entrada de la aplicación
├── pkg/
│   ├── cpu/              # Emulación del CPU (Sharp LR35902)
│   ├── memory/           # Gestión de memoria y mapeo
│   ├── gpu/              # Emulación de la PPU (Picture Processing Unit)
│   ├── input/            # Manejo de controles
│   ├── cartridge/        # Carga y gestión de cartuchos
│   └── debug/            # Herramientas de depuración
├── internal/
│   └── config/           # Configuración interna
└── docs/                 # Documentación adicional
```

## Componentes Principales

### CPU (Sharp LR35902)
- Procesador 8-bit personalizado similar al Z80
- Frecuencia: 4.19 MHz
- Implementación de instrucciones y registros

### Memoria
- 8 KB de RAM de trabajo (WRAM)
- 8 KB de VRAM
- Mapeo de memoria configurable

### GPU/PPU
- Resolución: 160x144 píxeles
- 4 tonos de gris
- Sprites y backgrounds

### Cartridge
- Soporte para diferentes MBC (Memory Bank Controllers)
- Manejo de ROM y RAM del cartucho

## Instalación

### Compilación con Makefile (Recomendado)

```bash
make build
```

### Compilación manual con Go

```bash
go build -o bin/gb-emulator ./cmd/gb-emulator
```

## Uso

### Usando Makefile

Compilar y ejecutar con una ROM:
```bash
make run ROM=ruta/a/juego.gb
```

Ejecutar sin recompilar (usa el binario existente):
```bash
make start ROM=ruta/a/juego.gb
```

### Comandos disponibles del Makefile

| Comando | Descripción |
|---------|-------------|
| `make build` | Compilar el proyecto |
| `make run ROM=<rom>` | Compilar y ejecutar con una ROM |
| `make start ROM=<rom>` | Ejecutar sin compilar |
| `make clean` | Limpiar archivos compilados |
| `make test` | Ejecutar tests |
| `make test-coverage` | Ejecutar tests con reporte de coverage |
| `make fmt` | Formatear código |
| `make vet` | Verificar código con go vet |
| `make deps` | Descargar dependencias |
| `make build-all` | Compilar para múltiples plataformas |
| `make help` | Mostrar ayuda completa |

### Ejemplos

```bash
# Compilar el proyecto
make build

# Ejecutar con una ROM de Tetris
make run ROM=tetris.gb

# Ejecutar con una ROM en carpeta específica
make run ROM=roms/pokemon.gb

# Ver ayuda completa
make help
```

### Uso directo del binario

```bash
./bin/gb-emulator <ruta_al_archivo_rom>
```

## Estado del Proyecto

Este proyecto está en fase inicial de desarrollo.

## Referencias

- [Pan Docs](https://gbdev.io/pandocs/) - Documentación técnica completa del Game Boy
- [Game Boy CPU Manual](http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf)
- [The Ultimate Game Boy Talk](https://www.youtube.com/watch?v=HyzD8pNlpwI)

## Licencia

Por definir
