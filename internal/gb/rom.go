package gb

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type GBHeader struct {
	Magic      [4]byte // Should be "NES^Z" or [0x4E, 0x45, 0x53, 0x1A]
	PRGROMSize byte    // Size of PRG ROM in 16 KB units
	CHRROMSize byte    // Size of CHR ROM in 8 KB units
	Flags6     byte    // Mapper, mirroring, battery, trainer
	Flags7     byte    // Mapper, VS/Playchoice, NES 2.0
	Flags8     byte    // PRG-RAM size (rarely used)
	Flags9     byte    // TV system (rarely used)
	Flags10    byte    // TV system, PRG-RAM presence (rarely used)
	Reserved   [5]byte // Reserved, should be zero
}

// PRGROM represents the Program ROM data of an NES ROM file
type PRGROM struct {
	Size int64  // Size in bytes
	Data []byte // The actual PRG ROM data
}

// readFileBytes lee todos los bytes de un archivo y los devuelve
// Parámetros:
//   - filePath: ruta del archivo a leer
//
// Retorna:
//   - []byte: contenido del archivo en bytes
//   - error: error si ocurre algún problema durante la lectura
func ReadFileBytes(filePath string) ([]byte, error) {
	// Abrir el archivo
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer file.Close()

	// Leer todos los bytes del archivo
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %w", err)
	}

	return data, nil
}

func ReadGBFile(filePath string) (*GBHeader, *PRGROM, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	header := &GBHeader{}
	err = binary.Read(file, binary.LittleEndian, header)
	if err != nil {
		return nil, nil, err
	}

	// Verify magic number
	//if string(header.Magic[:]) != "GB\x1A" {
	//	return nil, nil, fmt.Errorf("not a valid GB ROM file")
	//}

	// Load PRG ROM data (16 x PRGROMSize KB)
	prgROMSize := int64(header.PRGROMSize) * 16 * 1024 // Convert to bytes

	// Check if there's a trainer (512 bytes) that we need to skip
	if header.Flags6&0x04 != 0 {
		_, err = file.Seek(512, 1) // Skip trainer (current position + 512 bytes)
		if err != nil {
			return nil, nil, fmt.Errorf("error skipping trainer: %v", err)
		}
	}

	// Allocate a buffer for PRG ROM data
	prgROMData := make([]byte, prgROMSize)

	// Read PRG ROM data
	bytesRead, err := file.Read(prgROMData)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading PRG ROM data: %v", err)
	}

	if int64(bytesRead) != prgROMSize {
		return nil, nil, fmt.Errorf("unexpected PRG ROM size: got %d bytes, expected %d", bytesRead, prgROMSize)
	}

	// Create and populate the PRGROM struct
	prgROM := &PRGROM{
		Size: prgROMSize,
		Data: prgROMData,
	}

	return header, prgROM, nil
}
