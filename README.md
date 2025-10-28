# Game Boy Emulator

Emulador de Game Boy clásico escrito en Go.

## Descripción

Este proyecto implementa un emulador del sistema Game Boy original (DMG - Dot Matrix Game), permitiendo ejecutar ROMs de juegos clásicos.

## Requisitos

- Go 1.25 o superior
- Dependencias (se instalan automáticamente con `go mod download`):
  - [Ebiten v2](https://github.com/hajimehoshi/ebiten) - Biblioteca de desarrollo de juegos para rendering y ventanas

## Estructura del Proyecto

```
gb-emulator/
├── internal/
│   ├── cpu/              # Emulación del CPU (Sharp LR35902)
│   │   ├── cpu.go               # Estructura y registros del CPU
│   │   ├── instruction_execution.go  # Ejecución de instrucciones
│   │   ├── instruction_functions.go  # Implementación de instrucciones
│   │   └── instruction_map.go        # Mapeo de opcodes
│   ├── memory/           # Gestión de memoria y mapeo
│   │   ├── memory.go            # Sistema de memoria principal
│   │   └── memory_view.go       # Vistas y utilidades de memoria
│   ├── gb/               # Lógica principal del emulador
│   │   ├── gb.go                # Estructura principal del Game Boy
│   │   ├── game.go              # Loop principal del juego
│   │   └── rom.go               # Carga y gestión de ROMs
│   └── config/           # Configuración interna (en desarrollo)
├── gbctr.pdf             # Documentación técnica de referencia
├── go.mod                # Dependencias del proyecto
└── go.sum                # Checksums de dependencias
```

## Componentes Principales

### CPU (Sharp LR35902)
- Procesador 8-bit personalizado similar al Z80
- Frecuencia: 4.19 MHz
- **Estado actual**: ✅ Implementado
  - Registros: A (Acumulador), B, C, D, E, H, L
  - Registros de 16-bit: PC (Program Counter), SP (Stack Pointer)
  - Flags: Z (Zero), N (Subtraction), H (Half Carry), C (Carry)
  - Sistema de ejecución de instrucciones por ciclos
  - Mapeo de opcodes y funciones de instrucción

### Memoria
- Sistema de direccionamiento de 16-bit (0x0000 - 0xFFFF)
- **Estado actual**: ⚠️ En desarrollo
  - Lectura y escritura de memoria implementada
  - Sistema de mapeo de direcciones
  - Soporte para mirrors y bancos de memoria
  - ⚠️ Nota: Actualmente usa estructura de memoria tipo NES, necesita adaptación a Game Boy

### GPU/PPU (Picture Processing Unit)
- Resolución: 160x144 píxeles
- 4 tonos de gris
- Sprites y backgrounds
- **Estado actual**: ❌ Pendiente de implementación

### Rendering y Ventana
- **Estado actual**: ✅ Dependencias instaladas
  - Ebiten v2 para rendering 2D
  - Gestión de ventana y entrada de usuario

### Cartridge / ROM
- **Estado actual**: ⚠️ En desarrollo
  - Sistema básico de carga de ROMs implementado
  - Pendiente: Soporte para diferentes MBC (Memory Bank Controllers)
  - Pendiente: Manejo de RAM del cartucho

## Instalación

### Instalar dependencias

Primero, descarga todas las dependencias necesarias:

```bash
go mod download
```

O simplemente:

```bash
make deps
```

### Compilación con Makefile (Recomendado)

```bash
make build
```

### Compilación manual con Go

```bash
go build -o bin/gb-emulator ./cmd/gb-emulator
```

**Nota**: Si no existe el directorio `cmd/gb-emulator`, la compilación fallará. El punto de entrada de la aplicación aún está en desarrollo.

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

Este proyecto está en fase inicial de desarrollo. Componentes actuales:

### ✅ Completado
- Estructura base del proyecto
- Sistema de CPU con registros y flags
- Sistema de ejecución de instrucciones
- Dependencias de rendering (Ebiten v2)

### ⚠️ En Desarrollo
- Sistema de memoria (requiere adaptación de NES a Game Boy)
- Carga y gestión de ROMs
- Sistema de Game Boy principal (estructuras base implementadas)

### ❌ Pendiente
- Implementación completa del set de instrucciones del CPU
- PPU/GPU para rendering de gráficos
- Sistema de entrada (controles)
- Audio (APU)
- Debugging tools
- Tests unitarios y de integración

### 📝 Notas Técnicas
- Algunos componentes contienen código/comentarios de NES que necesitan ser adaptados a Game Boy
- La arquitectura de memoria necesita ajustarse al mapa de memoria del Game Boy
- Se recomienda revisar el archivo `gbctr.pdf` para especificaciones técnicas del hardware

## Referencias

- [Pan Docs](https://gbdev.io/pandocs/) - Documentación técnica completa del Game Boy
- [Game Boy CPU Manual](http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf)
- [The Ultimate Game Boy Talk](https://www.youtube.com/watch?v=HyzD8pNlpwI)

## Licencia

Por definir
