package cmdruntimes

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/runtimes"
)

var addLong = `add a runtime to your system or workspace`

func AddRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = runtimes.RunAddFromArgs(args)

	return err
}

var AddCmd = &cobra.Command{

	Use: "add",

	Short: "add a runtime to your system or workspace",

	Long: addLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = AddRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := AddCmd.HelpFunc()
	usage := AddCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/help", "<omit>", 0)
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/usage", "<omit>", 0)
		return usage(cmd)
	}
	AddCmd.SetHelpFunc(thelp)
	AddCmd.SetUsageFunc(tusage)

}
