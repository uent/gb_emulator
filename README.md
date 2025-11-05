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
│   │   ├── memory.go            # Sistema de memoria Game Boy completo
│   │   └── memory_view.go       # Vistas y utilidades de memoria
│   ├── gb/               # Lógica principal del emulador
│   │   ├── gb.go                # Estructura principal del Game Boy
│   │   ├── game.go              # Loop principal del juego (Ebiten)
│   │   └── rom.go               # Carga y gestión de ROMs/Boot ROM
│   └── config/           # Configuración interna (en desarrollo)
├── roms/                 # Directorio para archivos ROM (.gb, .gbc)
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
  - PC inicializado correctamente en 0x0000
  - Sistema de ejecución de instrucciones por ciclos
  - Mapeo de opcodes y funciones de instrucción
  - Instrucciones implementadas:
    - 0x00: NOP (No Operation)
    - 0x06: LD (Load Immediate)
    - 0x41: LD (Load Register to Register)

### Memoria
- Sistema de direccionamiento de 16-bit (0x0000 - 0xFFFF)
- **Estado actual**: ✅ Implementado
  - Mapa de memoria completo del Game Boy:
    - 0x0000-0x3FFF: ROM Bank #0 (16KB) / Boot ROM
    - 0x4000-0x7FFF: ROM Bank #1 switchable (16KB)
    - 0x8000-0x9FFF: Video RAM (8KB)
    - 0xA000-0xBFFF: External RAM switchable (8KB)
    - 0xC000-0xDFFF: Work RAM (8KB)
    - 0xE000-0xFDFF: Echo RAM
    - 0xFE00-0xFE9F: OAM (Sprite Attribute Memory)
    - 0xFF00-0xFF4B: I/O Ports
    - 0xFF80-0xFFFE: High RAM (HRAM)
    - 0xFFFF: Interrupt Enable Register
  - Lectura de memoria implementada con soporte para Boot ROM
  - Sistema de bancos de memoria preparado

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
- **Estado actual**: ✅ Implementado (básico)
  - Función `LoadROM()` para cargar ROMs en memoria
  - Función `LoadBootROM()` para cargar Boot ROM
  - Utilidad `ReadFileBytes()` para lectura de archivos
  - Soporte para ROMs en carpeta `roms/`
  - Pendiente: Soporte para diferentes MBC (Memory Bank Controllers)
  - Pendiente: Validación completa de headers de cartuchos
  - Pendiente: Manejo de RAM del cartucho con persistencia

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
- Sistema de CPU con registros, flags y PC inicializado
- Sistema de ejecución de instrucciones
- Mapa de memoria completo del Game Boy (adaptado correctamente desde NES)
- Carga de ROMs y Boot ROM en memoria
- Dependencias de rendering (Ebiten v2)
- Instrucciones básicas del CPU (NOP, LD)

### ⚠️ En Desarrollo
- Sistema de Game Boy principal (estructuras base implementadas)
- Escritura de memoria (función Write pendiente de completar)
- Sistema de bancos de memoria conmutables (MBC)

### ❌ Pendiente
- Implementación completa del set de instrucciones del CPU (restantes ~500 instrucciones)
- PPU/GPU para rendering de gráficos
- Sistema de entrada (controles/joypad)
- Audio (APU)
- Interrupciones
- Timers
- Debugging tools
- Tests unitarios y de integración
- Loop principal del emulador

### 📝 Notas Técnicas
- ✅ El mapa de memoria ya está correctamente adaptado al Game Boy (no más referencias a NES)
- La función `Write()` en memory.go necesita implementación completa
- Se recomienda revisar el archivo `gbctr.pdf` para especificaciones técnicas del hardware
- El sistema soporta Boot ROM para emular el inicio real del Game Boy

## Referencias

- [Pan Docs](https://gbdev.io/pandocs/) - Documentación técnica completa del Game Boy
- [Game Boy CPU Manual](http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf)
- [The Ultimate Game Boy Talk](https://www.youtube.com/watch?v=HyzD8pNlpwI)

## Licencia

Por definir
