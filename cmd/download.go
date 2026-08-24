package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/JannoTjarks/vidya-intarweb-playlist-cli/internal/utils"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads tracks from VIP servers",
	Long:  "Download the tracks from the VIP servers of the chosen roster. The files are downloaded as .m4a files and will also get tagged with some metadata, e.g. Title and Album/Game.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DownloadTracksByRoster(
			viper.GetString(rosterName.name),
			viper.GetString(destinationPath.name),
		)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().String(rosterName.name, rosterName.value, rosterName.usage)
	downloadCmd.Flags().String(destinationPath.name, destinationPath.value, destinationPath.usage)

	downloadCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())
		viper.AutomaticEnv()

		viper.BindEnv(rosterName.name, rosterName.envVar)

		fileStat, err := os.Stat(viper.GetString(destinationPath.name))
		if err != nil {
			return err
		}

		if fileStat.IsDir() != true {
			return errors.New("the given destinationPath has to be a directory")
		}

		return nil
	}

}
