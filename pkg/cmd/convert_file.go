package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var file *string
var extension *string
var path string

var fileCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert file",
	Long:  `Convert file to another type`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file = &args[0]      // passa o endereço
		extension = &args[1] // passa o endereço
		if path != "" {
			log.Println("Converting file", *file, "to", *extension, "at", path)
		}

	},
}

func init() {
	fileCmd.SetArgs([]string{"file"})
	fileCmd.Flags().StringVarP(&path, "path", "p", "", "Path to file")
	rootCmd.AddCommand(fileCmd)
}
