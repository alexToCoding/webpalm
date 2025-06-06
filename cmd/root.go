package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/alexToCoding/webpalm/v2/core"
	"github.com/alexToCoding/webpalm/v2/shared"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     usage(),
	Short:   "A web scraping tool",
	Long:    long(),
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		options, err := shared.ValidateThenBuildOption(cmd)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		fmt.Print(banner())
		options.PrintBanner()
		crawler := core.NewCrawler(options)
		crawler.Crawl()
	},
	Example: example() + regexestable(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "target url / ex: -u https://google.com")
	if err := rootCmd.MarkFlagRequired("url"); err != nil {
		return
	}
	rootCmd.Flags().IntP("level", "l", 0, "level of palming / ex: -l2")

	rootCmd.Flags().StringP("output", "o", "", "file to export the result (f.json, f.xml, f.txt) / ex: -o result.json")

	// rootCmd.Flags().StringToString("regexes", map[string]string{}, "regexes to match in each page / ex: --regexes comments=\"\\<\\!--.*?-->\"")

	var regexesFlag shared.RegexFlag
	rootCmd.Flags().Var(&regexesFlag, "regexes", "regexes to match in each page / ex: --regexes comments=\"\\<\\!--.*?-->\"")

	rootCmd.Flags().IntSliceP("exclude-code", "x", []int{}, "status codes to exclude / ex : -x 404,500")

	rootCmd.Flags().StringSliceP("include", "i", []string{}, "include only domains / ex : -i google.com,facebook.com")

	rootCmd.Flags().IntP("worker", "w", 0, "number of workers for multi-threading  / ex: -w 10")

	rootCmd.Flags().IntP("delay", "d", 0, "delay (ms) between each request / ex: -d 200")

	rootCmd.Flags().StringP("proxy", "p", "", "proxy to use / ex: -p http://proxy.com:8080")

	rootCmd.Flags().IntP("timeout", "t", 10, "timeout in seconds / ex: -t 10")

	rootCmd.Flags().StringP("user-agent", "a", "", "user agent to use / ex: -a chrome, firefox, safari, ie, edge, opera, android, ios, custom")

}
