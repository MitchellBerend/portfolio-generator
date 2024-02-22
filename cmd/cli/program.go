package cli

import (
	"os"

	"github.com/spf13/cobra"
)

type Input string
type InputType string
type Output string

var (
	file InputType = "file"
	dir InputType = "dir"
	none InputType = "none"
)

// A Program contains the data for the files that need to
// be used to generate new html pages
type Program struct {
	Input     Input
	InputType InputType
	Output    Output
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "portfolio-generator",
	Short: "A program that generates html pages",
	Long: `This program is designed to be able to convert json files into html pages that can
susequently get uploaded to a static file store.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCommand)
	generateCommand.Flags().StringP("input", "i", ".", "File or dir that contains json files")
	generateCommand.Flags().StringP("output", "o", ".", "Location of output target")

	rootCmd.AddCommand(newCommand)
}
