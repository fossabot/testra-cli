package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"github.com/testra-tech/testra-cli/internal/utils"
)

func GetResolvedProjectId(cmd *cobra.Command) string {
	projectId, _ := cmd.Flags().GetString(PROJECT_ID_FLAG_NAME)

	if EMPTY_STR == projectId {
		projectId = viper.GetString("defaultProjectId")
		utils.Info("ProjectId flag not passed, using default ProjectId from testra config")
	}

	if EMPTY_STR == projectId {
		utils.Danger("Default ProjectId in testra config is empty. Set using command `testra config set -p={PROJECT_ID}`")
		os.Exit(1)
	}

	return projectId
}
