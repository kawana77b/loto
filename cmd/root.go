package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kawana77b/loto/internal/loto"
	"github.com/kawana77b/loto/internal/prompt"
	"github.com/kawana77b/loto/internal/util"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var Version string = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loto",
	Short: "Proposing lottery ticket candidates for Japan (Takarakuji)",
	Long: `Proposing lottery ticket candidates for Japan (Takarakuji).
Applicable to "Loto" or "Numbers".

This tool is purely a complete random pick;
it does not analyze or suggest candidates, nor does it guarantee winning.`,
	Args:    cobra.MatchAll(cobra.RangeArgs(0, 1)),
	PreRunE: preRunRoot,
	RunE:    runRoot,
}

type rootOptions struct {
	lotteryType loto.LotteryType
	length      int
}

var rootOpts rootOptions

func preRunRoot(cmd *cobra.Command, args []string) error {
	// lottery type
	if len(args) == 0 {
		lt, _ := prompt.PromptLotteryType()
		rootOpts.lotteryType = loto.LotteryType(lt)
	} else {
		rootOpts.lotteryType = loto.LotteryType(args[0])
	}

	// --length
	length, _ := cmd.Flags().GetInt("length")
	rootOpts.length = util.Abs(length)

	// validatation
	if err := (rootOpts.lotteryType).Validate(); err != nil {
		return err
	}
	if rootOpts.length <= 0 {
		os.Exit(1)
	}
	return nil
}

func runRoot(cmd *cobra.Command, args []string) error {
	// Create lottery game
	lottery := loto.NewLottery(rootOpts.lotteryType)
	if lottery == nil {
		return fmt.Errorf("failed to create lottery for type: %s", rootOpts.lotteryType)
	}

	// Pick lottery numbers
	results := lottery.PickN(rootOpts.length)

	// Display results in table format
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"No", "Result"})

	// Determine if we need zero-padding based on lottery category
	isLoto := loto.GetCategory(rootOpts.lotteryType) == loto.LOTO

	for i, result := range results {
		numbers := make([]string, len(result))
		for j, num := range result {
			if isLoto {
				// Loto: 2-digit zero-padded format (e.g., 01, 07, 38)
				numbers[j] = fmt.Sprintf("%02d", num)
			} else {
				// Numbers: no padding (e.g., 8, 3, 3)
				numbers[j] = strconv.Itoa(num)
			}
		}

		// Join numbers differently based on type
		var numbersStr string
		if isLoto {
			numbersStr = strings.Join(numbers, ", ")
		} else {
			numbersStr = strings.Join(numbers, "")
		}

		table.Append([]string{
			strconv.Itoa(i + 1),
			numbersStr,
		})
	}

	return table.Render()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Version = Version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const quickPickDefaultCount int = 5

func init() {
	rootCmd.Flags().IntP("length", "n", quickPickDefaultCount, "Specify the number of lottery results to pick")
}
