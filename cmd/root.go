package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var jsonFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "si",
	Short: "SI is a utility that shows your system's information.",
	Long: `SI is a utility that shows your system's information.
Use additional sub commands to display specific information about your system.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello CLI")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&jsonFlag, "json", "", "display information in json form")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "display information in json form.")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	rootCmd.Flags().BoolP("json", "j", false, "display information in json form")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("inside initConfig")

	jsonFlag, _ := rootCmd.Flags().GetBool("json")
	if jsonFlag {
		fmt.Println("json flag:", jsonFlag)
	}

}
