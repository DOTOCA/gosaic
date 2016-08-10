package cmd

import (
	"gosaic/controller"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	coverAspectWidth  int
	coverAspectHeight int
	coverAspect       string
	coverAspectNum    int
)

func init() {
	addLocalIntFlag(&coverAspectWidth, "width", "", 0, "Pixel width of cover", CoverAspectCmd)
	addLocalIntFlag(&coverAspectHeight, "height", "", 0, "Pixel height of cover", CoverAspectCmd)
	addLocalFlag(&coverAspect, "aspect", "a", "1x1", "Aspect of cover partials (CxR)", CoverAspectCmd)
	addLocalIntFlag(&coverAspectNum, "size", "s", 0, "Number of partials in smallest dimension", CoverAspectCmd)
	RootCmd.AddCommand(CoverAspectCmd)
}

var CoverAspectCmd = &cobra.Command{
	Use:   "cover_aspect NAME",
	Short: "Create a aspect cover",
	Long:  "Create a aspect cover",
	Run: func(c *cobra.Command, args []string) {
		if len(args) != 1 {
			Env.Fatalln("name is required")
		}

		if args[0] == "" {
			Env.Fatalln("name is required")
		}

		if coverAspectWidth == 0 {
			Env.Fatalln("width is required")
		} else if coverAspectWidth < 0 {
			Env.Fatalln("width must be greater than zero")
		}

		if coverAspectHeight == 0 {
			Env.Fatalln("height is required")
		} else if coverAspectHeight < 0 {
			Env.Fatalln("height must be greater than zero")
		}

		if coverAspect == "" {
			Env.Fatalln("aspect is required")
		}

		aspectStrings := strings.Split(coverAspect, "x")
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

		if coverAspectNum == 0 {
			Env.Fatalln("num is required")
		} else if coverAspectNum < 0 {
			Env.Fatalln("num must be greater than zero")
		}

		err = Env.Init()
		if err != nil {
			Env.Fatalf("Unable to initialize environment: %s\n", err.Error())
		}
		defer Env.Close()

		controller.CoverAspect(Env, args[0], coverAspectWidth, coverAspectHeight, aw, ah, coverAspectNum)
	},
}