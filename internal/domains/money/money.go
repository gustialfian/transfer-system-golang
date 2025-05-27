package money

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const Scale int = 5

var ErrMoneyParseFail = errors.New("money parse fail")

// StringToInt converts a string representation of a decimal number to an integer,
// scaling it by the specified number of decimal places (scale).
// For example, StringToInt("12.34", 2) returns 1234.
// Returns an error if the input string cannot be parsed as a float.
func StringToInt(val string, scale int) (int, error) {
	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float: %w", err)
	}
	scaleMod := math.Pow10(scale)
	result := math.Trunc(valFloat * scaleMod)
	return int(result), nil
}

// IntToString converts an integer value representing a scaled monetary amount into its string
// representation with up to scale decimal places. The 'val' parameter is the integer value,
// and 'scale' specifies the number of decimal places the integer should be divided by (as a power of 10).
// For example, IntToString(12345, 2) returns "123.45000".
func IntToString(val int, scale int) string {
	scaleMod := math.Pow10(scale)
	result := float64(val) / scaleMod
	return fmt.Sprintf("%.5f", result)
}
