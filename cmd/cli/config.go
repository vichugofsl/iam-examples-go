package cli

import (
	"iam-examples-go/cmd/cli/usermigration"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "user-uploader-go",
	Short: "CLI command for creating records in the iam_user, iam_user_keys and iam_access_group_api_key entity",
	Long: `user-uploader-go is a CLI tool built in Go, which creates new records in the
iam_user entity after running a specific SQL query on a MySQL database.`,
}

func init() {
	rootCmd.AddCommand(usermigration.GetCommand())
}

func Execute() error {
	return rootCmd.Execute()
}
