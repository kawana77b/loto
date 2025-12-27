package loto_test

import (
	"testing"

	"github.com/kawana77b/loto/internal/loto"
)

// TestLotteryType_Validate tests the Validate method of LotteryType
func TestLotteryType_Validate(t *testing.T) {
	tests := []struct {
		name        string
		lotteryType loto.LotteryType
		wantErr     bool
	}{
		{
			name:        "valid loto6",
			lotteryType: loto.LOTO_6,
			wantErr:     false,
		},
		{
			name:        "valid loto7",
			lotteryType: loto.LOTO_7,
			wantErr:     false,
		},
		{
			name:        "valid miniloto",
			lotteryType: loto.LOTO_MINI,
			wantErr:     false,
		},
		{
			name:        "valid numbers3",
			lotteryType: loto.NUMBERS_3,
			wantErr:     false,
		},
		{
			name:        "valid numbers4",
			lotteryType: loto.NUMBERS_4,
			wantErr:     false,
		},
		{
			name:        "invalid lottery type",
			lotteryType: loto.LotteryType("invalid"),
			wantErr:     true,
		},
		{
			name:        "empty lottery type",
			lotteryType: loto.LotteryType(""),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.lotteryType.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("LotteryType.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestGetKind tests the GetKind function
func TestGetKind(t *testing.T) {
	tests := []struct {
		name        string
		lotteryType loto.LotteryType
		want        loto.LotteryCategory
	}{
		{
			name:        "loto6 returns LOTO category",
			lotteryType: loto.LOTO_6,
			want:        loto.LOTO,
		},
		{
			name:        "loto7 returns LOTO category",
			lotteryType: loto.LOTO_7,
			want:        loto.LOTO,
		},
		{
			name:        "miniloto returns LOTO category",
			lotteryType: loto.LOTO_MINI,
			want:        loto.LOTO,
		},
		{
			name:        "numbers3 returns NUMBERS category",
			lotteryType: loto.NUMBERS_3,
			want:        loto.NUMBERS,
		},
		{
			name:        "numbers4 returns NUMBERS category",
			lotteryType: loto.NUMBERS_4,
			want:        loto.NUMBERS,
		},
		{
			name:        "invalid type returns empty",
			lotteryType: loto.LotteryType("invalid"),
			want:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := loto.GetCategory(tt.lotteryType)
			if got != tt.want {
				t.Errorf("GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestNewBox tests the NewBox function
func TestNewBox(t *testing.T) {
	tests := []struct {
		name       string
		min        int
		max        int
		wantLength int
	}{
		{
			name:       "loto range 1-43",
			min:        1,
			max:        43,
			wantLength: 43,
		},
		{
			name:       "numbers range 0-9",
			min:        0,
			max:        9,
			wantLength: 10,
		},
		{
			name:       "single element",
			min:        5,
			max:        5,
			wantLength: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := loto.NewBox(tt.min, tt.max)
			if box.Length() != tt.wantLength {
				t.Errorf("NewBox().Length() = %v, want %v", box.Length(), tt.wantLength)
			}
		})
	}
}

// TestBox_PickN tests the PickN method of Box
func TestBox_PickN(t *testing.T) {
	box := loto.NewBox(1, 43)

	tests := []struct {
		name      string
		n         int
		wantLen   int
		checkUniq bool
	}{
		{
			name:      "pick 6 numbers",
			n:         6,
			wantLen:   6,
			checkUniq: true,
		},
		{
			name:      "pick 0 numbers",
			n:         0,
			wantLen:   0,
			checkUniq: false,
		},
		{
			name:      "pick negative numbers",
			n:         -1,
			wantLen:   0,
			checkUniq: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := box.PickN(tt.n)
			if len(result) != tt.wantLen {
				t.Errorf("Box.PickN() length = %v, want %v", len(result), tt.wantLen)
			}

			// Check uniqueness if required
			if tt.checkUniq && len(result) > 0 {
				seen := make(map[int]bool)
				for _, num := range result {
					if seen[num] {
						t.Errorf("Box.PickN() returned duplicate number: %d", num)
					}
					seen[num] = true

					// Check range
					if num < 1 || num > 43 {
						t.Errorf("Box.PickN() returned out of range number: %d", num)
					}
				}
			}
		})
	}
}

// TestBox_PickDupN tests the PickDupN method of Box
func TestBox_PickDupN(t *testing.T) {
	box := loto.NewBox(0, 9)

	tests := []struct {
		name    string
		n       int
		wantLen int
	}{
		{
			name:    "pick 3 digits",
			n:       3,
			wantLen: 3,
		},
		{
			name:    "pick 4 digits",
			n:       4,
			wantLen: 4,
		},
		{
			name:    "pick 0 digits",
			n:       0,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := box.PickDupN(tt.n)
			if len(result) != tt.wantLen {
				t.Errorf("Box.PickDupN() length = %v, want %v", len(result), tt.wantLen)
			}

			// Check range
			for _, num := range result {
				if num < 0 || num > 9 {
					t.Errorf("Box.PickDupN() returned out of range number: %d", num)
				}
			}
		})
	}
}

// TestBox_Operations tests various Box operations
func TestBox_Operations(t *testing.T) {
	t.Run("Append", func(t *testing.T) {
		box := loto.NewBox(1, 5)
		initialLen := box.Length()
		box.Append(6, 7, 8)
		if box.Length() != initialLen+3 {
			t.Errorf("Box.Append() length = %v, want %v", box.Length(), initialLen+3)
		}
	})

	t.Run("Clear", func(t *testing.T) {
		box := loto.NewBox(1, 10)
		box.Clear()
		if box.Length() != 0 {
			t.Errorf("Box.Clear() length = %v, want 0", box.Length())
		}
	})

	t.Run("Clone", func(t *testing.T) {
		box := loto.NewBox(1, 5)
		clone := box.Clone()
		if clone.Length() != box.Length() {
			t.Errorf("Box.Clone() length = %v, want %v", clone.Length(), box.Length())
		}
		// Modify clone shouldn't affect original
		clone.Append(99)
		if box.Length() == clone.Length() {
			t.Error("Box.Clone() didn't create independent copy")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		box := loto.NewBox(1, 10)
		if !box.Contains(5) {
			t.Error("Box.Contains(5) = false, want true")
		}
		if box.Contains(99) {
			t.Error("Box.Contains(99) = true, want false")
		}
	})
}

// TestNewLottery tests the NewLottery function
func TestNewLottery(t *testing.T) {
	tests := []struct {
		name        string
		lotteryType loto.LotteryType
		wantNil     bool
	}{
		{
			name:        "create loto6",
			lotteryType: loto.LOTO_6,
			wantNil:     false,
		},
		{
			name:        "create numbers3",
			lotteryType: loto.NUMBERS_3,
			wantNil:     false,
		},
		{
			name:        "invalid type",
			lotteryType: loto.LotteryType("invalid"),
			wantNil:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lottery := loto.NewLottery(tt.lotteryType)
			if (lottery == nil) != tt.wantNil {
				t.Errorf("NewLottery() = %v, wantNil %v", lottery, tt.wantNil)
			}
		})
	}
}

// TestLotteryGame_Pick tests the Pick method of LotteryGame
func TestLotteryGame_Pick(t *testing.T) {
	tests := []struct {
		name        string
		lotteryType loto.LotteryType
		wantCount   int
		allowDup    bool
		checkRange  func(int) bool
	}{
		{
			name:        "loto6 pick",
			lotteryType: loto.LOTO_6,
			wantCount:   6,
			allowDup:    false,
			checkRange: func(n int) bool {
				return n >= 1 && n <= 43
			},
		},
		{
			name:        "loto7 pick",
			lotteryType: loto.LOTO_7,
			wantCount:   7,
			allowDup:    false,
			checkRange: func(n int) bool {
				return n >= 1 && n <= 43
			},
		},
		{
			name:        "miniloto pick",
			lotteryType: loto.LOTO_MINI,
			wantCount:   5,
			allowDup:    false,
			checkRange: func(n int) bool {
				return n >= 1 && n <= 37
			},
		},
		{
			name:        "numbers3 pick",
			lotteryType: loto.NUMBERS_3,
			wantCount:   3,
			allowDup:    true,
			checkRange: func(n int) bool {
				return n >= 0 && n <= 9
			},
		},
		{
			name:        "numbers4 pick",
			lotteryType: loto.NUMBERS_4,
			wantCount:   4,
			allowDup:    true,
			checkRange: func(n int) bool {
				return n >= 0 && n <= 9
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lottery := loto.NewLottery(tt.lotteryType)
			if lottery == nil {
				t.Fatal("NewLottery() returned nil")
			}

			result := lottery.Pick()

			// Check count
			if len(result) != tt.wantCount {
				t.Errorf("Pick() length = %v, want %v", len(result), tt.wantCount)
			}

			// Check range
			for _, num := range result {
				if !tt.checkRange(num) {
					t.Errorf("Pick() returned out of range number: %d", num)
				}
			}

			// Check uniqueness for non-duplicate lotteries
			if !tt.allowDup {
				seen := make(map[int]bool)
				for _, num := range result {
					if seen[num] {
						t.Errorf("Pick() returned duplicate number in non-duplicate lottery: %d", num)
					}
					seen[num] = true
				}

				// Check that loto results are sorted in ascending order
				for i := 1; i < len(result); i++ {
					if result[i] <= result[i-1] {
						t.Errorf("Pick() result not sorted: %v", result)
						break
					}
				}
			}
		})
	}
}

// TestLotteryGame_PickN tests the PickN method of LotteryGame
func TestLotteryGame_PickN(t *testing.T) {
	tests := []struct {
		name        string
		lotteryType loto.LotteryType
		count       int
		wantCount   int
	}{
		{
			name:        "loto6 pick 5 times",
			lotteryType: loto.LOTO_6,
			count:       5,
			wantCount:   5,
		},
		{
			name:        "numbers3 pick 3 times",
			lotteryType: loto.NUMBERS_3,
			count:       3,
			wantCount:   3,
		},
		{
			name:        "pick 1 time",
			lotteryType: loto.LOTO_6,
			count:       1,
			wantCount:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lottery := loto.NewLottery(tt.lotteryType)
			if lottery == nil {
				t.Fatal("NewLottery() returned nil")
			}

			results := lottery.PickN(tt.count)

			// Check count
			if len(results) != tt.wantCount {
				t.Errorf("PickN() length = %v, want %v", len(results), tt.wantCount)
			}

			// Check that each result is unique (no duplicate result sets)
			seen := make(map[string]bool)
			for _, result := range results {
				// Convert result to string for comparison
				key := ""
				for _, num := range result {
					key += string(rune(num))
				}
				if seen[key] {
					t.Errorf("PickN() returned duplicate result set")
				}
				seen[key] = true
			}
		})
	}
}

// TestLotteryConfigs tests that all lottery configs are properly configured
func TestLotteryConfigs(t *testing.T) {
	expectedConfigs := []loto.LotteryType{
		loto.LOTO_6,
		loto.LOTO_7,
		loto.LOTO_MINI,
		loto.NUMBERS_3,
		loto.NUMBERS_4,
	}

	for _, lotteryType := range expectedConfigs {
		t.Run(string(lotteryType), func(t *testing.T) {
			config, ok := loto.LotteryConfigs[lotteryType]
			if !ok {
				t.Fatalf("LotteryConfigs missing entry for %s", lotteryType)
			}

			// Check that config has valid values
			if config.Count <= 0 {
				t.Errorf("Config.Count = %v, want > 0", config.Count)
			}
			if config.Max <= config.Min {
				t.Errorf("Config.Max = %v, want > Min (%v)", config.Max, config.Min)
			}
			if config.Category == "" {
				t.Error("Config.Category is empty")
			}

			// Check category consistency
			expectedCategory := loto.GetCategory(lotteryType)
			if config.Category != expectedCategory {
				t.Errorf("Config.Category = %v, want %v", config.Category, expectedCategory)
			}
		})
	}
}
