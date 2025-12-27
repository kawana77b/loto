package cmd

import (
	"os"

	"github.com/kawana77b/loto/internal/loto"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Displays the available argument names",
	Long:    `Displays the available argument names.`,
	RunE:    runList,
}

func runList(cmd *cobra.Command, args []string) error {
	table, err := loto.Table(os.Stdout)
	if err != nil {
		return err
	}
	return table.Render()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
