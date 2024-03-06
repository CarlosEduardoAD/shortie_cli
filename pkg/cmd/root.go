package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filler",
	Short: "Filler is utility cli for converting files",
	Long:  `A fast and pratical cli for converting files of any type in your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello, World!")
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
