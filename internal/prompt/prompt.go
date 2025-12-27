package prompt

import (
	"github.com/kawana77b/loto/internal/loto"
	"github.com/manifoldco/promptui"
)

// PromptLotteryType prompts the user to select a lottery type and returns the selected type.
func PromptLotteryType() (string, bool) {
	prompt := promptui.Select{
		Label: "Select Lottery Type",
		Items: loto.Names(),
	}
	_, ans, err := prompt.Run()
	if err != nil {
		return "", false
	}
	return ans, true
}
