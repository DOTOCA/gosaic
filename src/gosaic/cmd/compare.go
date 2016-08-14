package cmd

import (
	"gosaic/controller"

	"github.com/spf13/cobra"
)

var (
	macroId int
)

func init() {
	addLocalIntFlag(&macroId, "macro_id", "", 0, "Id of macro for comparison", CompareCmd)
	RootCmd.AddCommand(CompareCmd)
}

var CompareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Build comparisons for macro against index",
	Long:  "Build comparisons for macro against index",
	Run: func(c *cobra.Command, args []string) {
		if macroId == 0 {
			Env.Fatalln("Macro id is required")
		}

		err := Env.Init()
		if err != nil {
			Env.Fatalf("Unable to initialize environment: %s\n", err.Error())
		}
		defer Env.Close()

		controller.Compare(Env, int64(macroId))
	},
}
