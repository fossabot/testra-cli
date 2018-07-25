package commons

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/testra-tech/testra-cli/internal/utils/notif"
	"os"
	"github.com/testra-tech/testra-cli/internal/constants/strs"
	"github.com/testra-tech/testra-cli/internal/constants/flags"
)

func GetResolvedProjectId(cmd *cobra.Command) string {
	projectId, _ := cmd.Flags().GetString(flags.PROJECT_ID)

	if strs.EmptyStr == projectId {
		projectId = viper.GetString("defaultProjectId")
		notif.Info("ProjectId flag not passed, using default ProjectId from testra config")
	}

	if strs.EmptyStr == projectId {
		notif.Danger("Default ProjectId in testra config is empty. Set using command `testra config set -p={PROJECT_ID}`")
		os.Exit(1)
	}

	return projectId
}