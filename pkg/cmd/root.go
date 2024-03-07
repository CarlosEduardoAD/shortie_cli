package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shortie",
	Short: "shortie_cli is utility cli shorting urls",
	Long:  `A fast and pratical cli for shorting urls, you can acess them via cli or in your browser!.`,
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
