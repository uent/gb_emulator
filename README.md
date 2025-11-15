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
│   │   ├── cpu.go                    # Estructura y registros del CPU con flags
│   │   ├── instruction_execution.go  # Ejecución de instrucciones (1 y 2 bytes)
│   │   ├── instruction_functions.go  # Implementación de instrucciones básicas
│   │   ├── instruction_map.go        # Mapeo de opcodes y tabla CB
│   │   ├── advances_functions.go     # Instrucciones avanzadas (prefijo CB)
│   │   ├── stack.go                  # Operaciones de stack (push/pop)
│   │   └── utils.go                  # Utilidades para manipulación de bytes
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
- **Estado actual**: ✅ Implementado (en expansión)
  - Registros de 8-bit: A (High de AF), B, C, D, E, H, L
  - Registros de 16-bit: PC (Program Counter), SP (Stack Pointer)
  - Flags implementados como booleanos con documentación mejorada:
    - ZFlag (Zero flag, bit 7 de AF) - Se activa cuando el resultado es 0
    - NFlag (Subtraction flag / BCD, bit 6 de AF) - Indica operación de resta
    - HFlag (Half Carry flag / BCD, bit 5 de AF) - Acarreo de los 4 bits bajos
    - CFlag (Carry flag, bit 4 de AF) - Acarreo/préstamo general
  - PC inicializado correctamente en 0x0000
  - **SP inicializado en 0xFFFE** (valor correcto del Game Boy)
  - **Stack completamente funcional** con operaciones push/pop de 8 y 16 bits
  - Sistema de ejecución de instrucciones por ciclos
  - Mapeo de opcodes y funciones de instrucción
  - Soporte para instrucciones de 2 bytes (prefijo 0xCB)
  - **Instrucciones implementadas** (27 instrucciones base + 2 avanzadas):
    - 0x00: NOP - No Operation (mueve PC, retorna 1 ciclo)
    - 0x05: DEC B - Decrementa registro B (con actualización de flags Z, N, H)
    - 0x06: LD B, d8 - Load immediate en registro B
    - 0x0C: INC C - Incrementa registro C (con actualización de flags Z, N, H)
    - 0x0E: LD C, d8 - Load immediate en registro C
    - 0x11: LD DE, d16 - Load immediate 16-bit en DE
    - 0x17: RLA - Rotate A left through carry
    - 0x1A: LD A, (DE) - Load en A desde memoria apuntada por DE
    - 0x20: JR NZ, s8 - Jump relativo si Z flag = 0
    - 0x21: LD HL, n16 - Load immediate 16-bit en HL
    - 0x26: LD H, d8 - Load immediate en registro H
    - 0x31: LD SP, n16 - Load immediate 16-bit en Stack Pointer
    - 0x32: LD (HL-), A - Store A en dirección HL y decrementar HL
    - 0x3E: LD A, d8 - Load immediate en registro A
    - 0x40: LD B, B - Load B en B
    - 0x41: LD B, C - Load C en B
    - 0x45: LD B, L - Load L en B
    - 0x4F: LD C, A - Load A en C
    - 0x77: LD (HL), A - Store A en memoria apuntada por HL
    - 0xAF: XOR A - XOR de A consigo mismo (resultado siempre 0)
    - 0xC1: POP BC - Pop del stack a BC
    - 0xC5: PUSH BC - Push BC al stack
    - 0xCD: CALL a16 - Call a dirección 16-bit
    - 0xE0: LD (a8), A - Store A en 0xFF00 + a8
    - 0xE2: LD (C), A - Store A en dirección 0xFF00 + C (I/O ports)
    - 0xCB11: RLC C - Rotate Left C
    - 0xCB7C: BIT 7, H - Test bit 7 del registro H
  - Utilidades implementadas:
    - MovePC() - Movimiento del Program Counter
    - jointBytesToUInt16() - Combinar bytes a uint16
    - splitUInt16ToBytes() - Dividir uint16 en bytes
    - calculateHalfFlagAdd() - Calcula half-carry flag para sumas
    - calculateHalfFlagSubtract() - Calcula half-carry flag para restas
    - calculateHalfFlagIncrement() - Calcula half-carry flag para incrementos
    - calculateHalfFlagDecrement() - Calcula half-carry flag para decrementos
    - bool2u8() - Convierte booleanos a uint8

### Memoria
- Sistema de direccionamiento de 16-bit (0x0000 - 0xFFFF)
- **Estado actual**: ✅ Implementado y refinado
  - Mapa de memoria completo del Game Boy (corregido y preciso):
    - 0x0000-0x3FFF: ROM Bank #0 (16KB) / Boot ROM
    - 0x4000-0x7FFF: ROM Bank #1 switchable (16KB)
    - 0x8000-0x9FFF: Video RAM (8KB)
    - 0xA000-0xBFFF: External RAM switchable (8KB)
    - 0xC000-0xCFFF: Work RAM Bank 0 (4KB)
    - 0xD000-0xDFFF: Work RAM Bank 1 switchable (4KB)
    - 0xE000-0xFDFF: Echo RAM (7680 bytes, mirror de WRAM)
    - 0xFE00-0xFE9F: OAM (Sprite Attribute Memory)
    - 0xFEA0-0xFEFF: Prohibido (no usable)
    - 0xFF00-0xFF7F: I/O Ports
    - 0xFF80-0xFFFE: High RAM (HRAM, 127 bytes)
    - 0xFFFF: Interrupt Enable Register (IE)
  - **Lectura y escritura de memoria completamente funcionales**
  - Función getMemoryAddress() optimizada con punteros
  - Soporte para Boot ROM con switch automático
  - Echo RAM correctamente mapeado a WRAM
  - Sistema de bancos de memoria preparado (pendiente MBC)

### GPU/PPU (Picture Processing Unit)
- Resolución: 160x144 píxeles
- 4 tonos de gris
- Sprites y backgrounds
- **Estado actual**: ❌ Pendiente de implementación

### Rendering y Ventana
- **Estado actual**: ✅ Loop principal implementado
  - Ebiten v2 para rendering 2D
  - Estructura Game con métodos Update, Draw y Layout
  - Loop principal ejecutándose (~70224 ciclos por frame)
  - Ventana configurada (512x480) con pantalla lógica de 160x144
  - Gestión de pausa implementada
  - Rendering básico (pantalla negra, pendiente integración con PPU)

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
- Sistema de CPU con registros, flags (Z, N, H, C) documentados, PC y SP inicializados
- **Stack completamente funcional**: push/pop de 8 y 16 bits (stack.go)
- Sistema de ejecución de instrucciones con soporte para opcodes de 1 y 2 bytes
- **Mapa de memoria refinado y preciso** (WRAM correctamente dividido en bancos, Echo RAM, IE register)
- **Lectura y escritura de memoria completamente funcionales**
- Carga de ROMs y Boot ROM en memoria
- **Loop principal del emulador funcional** con Ebiten v2
- **29 instrucciones del CPU** implementadas (LD, INC, DEC, JR, XOR, BIT, PUSH, POP, CALL, RLA, RLC)
- Funciones auxiliares para manipulación de datos (split/join bytes, half-carry flags para add/sub/inc/dec, bool2u8)
- Método MovePC para gestión del Program Counter
- Tabla de instrucciones avanzadas (prefijo CB) con 2 instrucciones
- Sistema de cálculo de half-carry flags para operaciones aritméticas (suma, resta, incremento, decremento)
- Instrucciones de control de flujo: CALL con manejo automático del stack
- **Game loop ejecutándose** con Update (70224 ciclos/frame), Draw y Layout implementados

### ⚠️ En Desarrollo
- Integración del PPU con el loop principal (estructura lista, pendiente rendering real)
- Sistema de bancos de memoria conmutables (MBC1, MBC3, MBC5)
- Expansión del set de instrucciones del CPU (~219 restantes)

### ❌ Pendiente
- Implementación completa del set de instrucciones del CPU (~219 instrucciones restantes)
- Instrucciones CB restantes (~254 instrucciones)
- PPU/GPU para rendering de gráficos (tiles, sprites, backgrounds)
- Sistema de entrada (controles/joypad)
- Audio (APU)
- Sistema de interrupciones completo
- Timers
- Debugging tools
- Tests unitarios y de integración
- Instrucciones de retorno (RET, RETI) y otras de control de flujo
- Sincronización precisa de timing (actualmente ~70224 ciclos fijos por frame)

### 📝 Notas Técnicas
- ✅ El mapa de memoria completamente adaptado al Game Boy con direccionamiento preciso
- ✅ **Stack Pointer inicializado en 0xFFFE** (valor estándar del Game Boy al inicio)
- ✅ **Sistema de stack funcional** con operaciones push/pop correctamente implementadas
- ✅ Flags del CPU implementados como booleanos separados con documentación detallada
- ✅ Soporte para instrucciones de 2 bytes con prefijo CB implementado
- ✅ Funciones auxiliares para conversión byte ↔ uint16 (little-endian)
- ✅ Sistema de cálculo de half-carry flag para operaciones aritméticas (suma, resta, incremento, decremento)
- ✅ Tabla de instrucciones simplificada (uso de inicialización de structs sin puntero explícito)
- ✅ **Loop principal del emulador ejecutándose** con ciclos por frame (~70224 ciclos)
- ✅ Directorio `roms/` disponible para almacenar archivos ROM (.gb, .gbc)
- ✅ **Función `Write()` completamente implementada** con manejo optimizado de punteros
- ✅ **getMemoryAddress()** retorna punteros para lectura y escritura eficiente
- ✅ WRAM dividido correctamente en Bank 0 (4KB) y Bank 1 (4KB) switchable
- ✅ Echo RAM implementado como espejo de WRAM (direcciones 0xE000-0xFDFF)
- ✅ Registro IE (Interrupt Enable) en 0xFFFF correctamente implementado
- ✅ Instrucciones de control de flujo: CALL implementado con push automático del PC
- Se recomienda revisar el archivo `gbctr.pdf` para especificaciones técnicas del hardware
- El sistema soporta Boot ROM para emular el inicio real del Game Boy
- Referencias de documentación integradas en el código:
  - [CPU Registers and Flags](https://gbdev.io/pandocs/CPU_Registers_and_Flags.html)
  - [GB Opcodes Generator](https://meganesu.github.io/generate-gb-opcodes/)
  - [RGBDS Instruction Set](https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7)

## Referencias

- [Pan Docs](https://gbdev.io/pandocs/) - Documentación técnica completa del Game Boy
- [Game Boy CPU Manual](http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf)
- [The Ultimate Game Boy Talk](https://www.youtube.com/watch?v=HyzD8pNlpwI)

## Licencia

Por definir
