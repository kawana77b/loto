package loto

import (
	"fmt"
	"slices"
	"strings"
)

// LotteryCategory represents the category of lottery (LOTO or NUMBERS).
type LotteryCategory string

// LotteryType represents a specific lottery type.
type LotteryType string

const (
	// Lottery categories
	LOTO    = LotteryCategory("loto")
	NUMBERS = LotteryCategory("numbers")

	// Lottery types
	LOTO_6    = LotteryType("loto6")
	LOTO_7    = LotteryType("loto7")
	LOTO_MINI = LotteryType("miniloto")
	NUMBERS_3 = LotteryType("numbers3")
	NUMBERS_4 = LotteryType("numbers4")
)

// Validate checks if the lottery type is valid.
func (c LotteryType) Validate() error {
	if _, ok := LotteryConfigs[c]; !ok {
		// Build list of valid types from config
		validTypes := make([]string, 0, len(LotteryConfigs))
		for lotteryType := range LotteryConfigs {
			validTypes = append(validTypes, string(lotteryType))
		}
		slices.Sort(validTypes)
		return fmt.Errorf("invalid lottery type: %s. It must be one of %s", c, strings.Join(validTypes, ", "))
	}
	return nil
}

// String returns the string representation of the lottery type.
func (t LotteryType) String() string {
	return string(t)
}

// GetCategory returns the category of lottery (LOTO or NUMBERS) based on the given LotteryType.
func GetCategory(t LotteryType) LotteryCategory {
	if config, ok := LotteryConfigs[t]; ok {
		return config.Category
	}
	return ""
}
