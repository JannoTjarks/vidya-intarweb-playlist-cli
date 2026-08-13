package cmd

import (
	"fmt"

	"github.com/JannoTjarks/vidya-intarweb-playlist-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show the current version of vidya-intarweb-playlist-cli",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("json") {
			fmt.Println(utils.GenerateVersionJson())
			return
		}

		fmt.Println(utils.GenerateVersionSignature())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().Bool("json", false, "If set, the output will be formatted as json")

	viper.BindPFlags(versionCmd.Flags())
}
