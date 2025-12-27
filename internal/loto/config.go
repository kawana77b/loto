package loto

// LotteryConfig holds the configuration for a lottery type.
type LotteryConfig struct {
	Category       LotteryCategory // Category of the lottery (LOTO or NUMBERS)
	Count          int             // Number of numbers/digits to pick
	Min            int             // Minimum value in range
	Max            int             // Maximum value in range
	AllowDuplicate bool            // Whether duplicates are allowed (true for Numbers, false for Loto)
}

/**
 * About the Rules of Japanese Lottery (Takarakuji),
 * References:
 *  https://ja.wikipedia.org/wiki/%E3%83%8A%E3%83%B3%E3%83%90%E3%83%BC%E3%82%BA_(%E5%AE%9D%E3%81%8F%E3%81%98)
 *  https://ja.wikipedia.org/wiki/%E3%83%AD%E3%83%886
 */

// LotteryConfigs holds all lottery type configurations.
var LotteryConfigs = map[LotteryType]LotteryConfig{
	LOTO_6: {
		Category:       LOTO,
		Count:          6,
		Min:            1,
		Max:            43,
		AllowDuplicate: false,
	},
	LOTO_7: {
		Category:       LOTO,
		Count:          7,
		Min:            1,
		Max:            37,
		AllowDuplicate: false,
	},
	LOTO_MINI: {
		Category:       LOTO,
		Count:          5,
		Min:            1,
		Max:            31,
		AllowDuplicate: false,
	},
	NUMBERS_3: {
		Category:       NUMBERS,
		Count:          3,
		Min:            0,
		Max:            9,
		AllowDuplicate: true,
	},
	NUMBERS_4: {
		Category:       NUMBERS,
		Count:          4,
		Min:            0,
		Max:            9,
		AllowDuplicate: true,
	},
}
