package bytes

import (
	"bulletproof-go/internal/utils"
	"log"
)

type SizeUnit string

const (
	Bytes     SizeUnit = "bytes"
	Kilobytes SizeUnit = "kilobytes"
	Megabytes SizeUnit = "megabytes"
	Gigabytes SizeUnit = "gigabytes"
)

func ConvertToBytes(size int64, unit SizeUnit) int64 {
	s, err := utils.ConvertToBytes(float64(size), utils.SizeUnit(unit))
	if err != nil {
		log.Fatalf("Failed to convert bytes: %v", err)
	}
	return int64(s)
}
