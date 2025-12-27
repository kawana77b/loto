package loto

import "slices"

// Lottery is an interface for lottery games.
type Lottery interface {
	// Perform a single random draw and obtain the result.
	Pick() []int
	// Perform multiple random draws and obtain the results.
	PickN(count int) [][]int
}

// LotteryGame is a generic lottery game implementation that works for all lottery types.
type LotteryGame struct {
	config LotteryConfig
	box    *Box
}

// NewLottery creates a new lottery game based on the given lottery type.
func NewLottery(t LotteryType) *LotteryGame {
	config, ok := LotteryConfigs[t]
	if !ok {
		return nil
	}
	return &LotteryGame{
		config: config,
		box:    NewBox(config.Min, config.Max),
	}
}

// Pick performs a single random draw and returns the result.
// For loto types (non-duplicate), the result is sorted in ascending order.
// For numbers types (duplicate allowed), the result is returned as-is.
func (l *LotteryGame) Pick() []int {
	var result []int
	if l.config.AllowDuplicate {
		// Numbers: return as-is (no sorting)
		result = l.box.PickDupN(l.config.Count)
	} else {
		// Loto: sort the result
		result = l.box.PickN(l.config.Count)
		slices.Sort(result)
	}
	return result
}

// PickN performs multiple random draws and returns the results.
func (l *LotteryGame) PickN(count int) [][]int {
	return pickN(l, count)
}

// pickN performs multiple random draws from the given lottery. Each result is unique.
func pickN(lottery Lottery, count int) [][]int {
	results := make([][]int, 0, count)
	for len(results) < count {
		picked := lottery.Pick()

		// The result should be that each element is unique.
		if !slices.ContainsFunc(results, func(r []int) bool {
			return slices.Equal(picked, r)
		}) {
			results = append(results, picked)
		}
	}
	return results
}
