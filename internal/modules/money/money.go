package money

import (
	"fmt"
	"math"
	"strconv"
)

const Scale int = 5

func StringToInt(val string, scale int) (int, error) {
	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid initial balance: %w", err)
	}
	scaleMod := math.Pow10(scale)
	result := math.Trunc(valFloat * scaleMod)
	return int(result), nil
}

func IntToString(val int, scale int) string {
	scaleMod := math.Pow10(scale)
	result := float64(val) / scaleMod
	return fmt.Sprintf("%.6f", result)
}
