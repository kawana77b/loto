package loto

import (
	"fmt"
	"io"
	"sort"

	"github.com/olekukonko/tablewriter"
)

// Names returns all available lottery type names sorted alphabetically.
func Names() []string {
	names := make([]string, 0, len(LotteryConfigs))
	for lotteryType := range LotteryConfigs {
		names = append(names, lotteryType.String())
	}
	sort.Strings(names)
	return names
}

// Table creates a formatted table of all available lottery types with their configurations.
func Table(w io.Writer) (*tablewriter.Table, error) {
	table := tablewriter.NewWriter(w)
	table.Header([]string{"Name", "Count", "Min", "Max", "Allow Duplicates"})

	for _, name := range Names() {
		config, ok := LotteryConfigs[LotteryType(name)]
		if !ok {
			return nil, fmt.Errorf("invalid lottery type: %s", name)
		}

		allowDup := "No"
		if config.AllowDuplicate {
			allowDup = "Yes"
		}

		table.Append([]string{
			name,
			fmt.Sprintf("%d", config.Count),
			fmt.Sprintf("%d", config.Min),
			fmt.Sprintf("%d", config.Max),
			allowDup,
		})
	}
	return table, nil
}
