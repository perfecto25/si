package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// osCmd represents the os command
var osCmd = &cobra.Command{
	Use:   "os",
	Short: "shows details about your Operating System and Kernel",
	Long:  `shows details about your Operating System and Kernel`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("os called")

	},
}

func init() {
	rootCmd.AddCommand(osCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// osCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// osCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
