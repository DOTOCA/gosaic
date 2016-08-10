package cmd

import (
	"gosaic/controller"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	macroAspectWidth  int
	macroAspectHeight int
	macroAspect       string
	macroAspectNum    int
)

func init() {
	addLocalIntFlag(&macroAspectWidth, "width", "", 0, "Pixel width of cover", MacroAspectCmd)
	addLocalIntFlag(&macroAspectHeight, "height", "", 0, "Pixel height of cover", MacroAspectCmd)
	addLocalFlag(&macroAspect, "aspect", "a", "1x1", "Aspect of cover partials (CxR)", MacroAspectCmd)
	addLocalIntFlag(&macroAspectNum, "size", "s", 0, "Number of partials in smallest dimension", MacroAspectCmd)
	RootCmd.AddCommand(MacroAspectCmd)
}

var MacroAspectCmd = &cobra.Command{
	Use:   "macro_aspect PATH",
	Short: "Add cover aspect and macro",
	Long:  "Add cover aspect and macro",
	Run: func(c *cobra.Command, args []string) {
		if len(args) != 1 {
			Env.Fatalln("PATH is required")
		}

		if args[0] == "" {
			Env.Fatalln("Macro path is required")
		}

		if macroAspectWidth == 0 {
			Env.Fatalln("width is required")
		} else if macroAspectWidth < 0 {
			Env.Fatalln("width must be greater than zero")
		}

		if macroAspectHeight == 0 {
			Env.Fatalln("height is required")
		} else if macroAspectHeight < 0 {
			Env.Fatalln("height must be greater than zero")
		}

		if macroAspect == "" {
			Env.Fatalln("aspect is required")
		}

		aspectStrings := strings.Split(macroAspect, "x")
		if len(aspectStrings) != 2 {
			Env.Fatalln("aspect format must be CxR")
		}

		aw, err := strconv.Atoi(aspectStrings[0])
		if err != nil {
			Env.Fatalf("Error converting aspect columns: %s\n", err.Error())
		}

		if aw == 0 {
			Env.Fatalln("aspect columns cannot be zero")
		} else if aw < 0 {
			Env.Fatalln("aspect columns must be greater than zero")
		}

		ah, err := strconv.Atoi(aspectStrings[1])
		if err != nil {
			Env.Fatalf("Error converting aspect rows: %s\n", err.Error())
		}

		if ah == 0 {
			Env.Fatalln("aspect rows cannot be zero")
		} else if ah < 0 {
			Env.Fatalln("aspect rows must be greater than zero")
		}

		if macroAspectNum == 0 {
			Env.Fatalln("num is required")
		} else if macroAspectNum < 0 {
			Env.Fatalln("num must be greater than zero")
		}

		err = Env.Init()
		if err != nil {
			Env.Fatalf("Unable to initialize environment: %s\n", err.Error())
		}
		defer Env.Close()

		controller.MacroAspect(Env, args[0], macroAspectWidth, macroAspectHeight, aw, ah, macroAspectNum)
	},
}