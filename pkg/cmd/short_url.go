package cmd

import (
	"log"
	"os"
	exec "os/exec"
	"runtime"

	"github.com/CarlosEduardoAD/shortie_cli/internal/application/commands"
	"github.com/spf13/cobra"
)

var originalUrl *string
var shortUrl *string

var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "Short Url",
	Long:  `Pass an url to be shortened and get the short url back.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		originalUrl = &args[0]
		shortUrl, err := commands.CreateShortUrl(*originalUrl)

		if err != nil {
			log.Panic(err)
			os.Exit(1)
		}

		log.Println(shortUrl)
	},
}

var acessCmd = &cobra.Command{
	Use:   "acess",
	Short: "Acess Short Url",
	Long:  `Pass an shorten url to be acessed.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shortUrl = &args[0]
		originalUrl, err := commands.FindOriginalUrl(*shortUrl)

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		switch runtime.GOOS {
		case "windows":
			exec.Command("explorer", originalUrl).Start()
		case "darwin":
			exec.Command("open", originalUrl).Start()
		default:
			exec.Command("xdg-open", originalUrl).Start()
		}
	},
}

func init() {
	shortCmd.SetArgs([]string{"originalUrl"})
	acessCmd.SetArgs([]string{"shortUrl"})
	rootCmd.AddCommand(acessCmd)
	rootCmd.AddCommand(shortCmd)
}
