package usermigration

import (
	"fmt"

	"iam-examples-go/core/adapters/mysql/connection"
	"iam-examples-go/core/adapters/mysql/repository"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "user-migration",
	Short: "Create records in the iam_user entity",
	Long:  `Create records in the iam_user entity using a specific SQL query`,
	Run: func(cmd *cobra.Command, args []string) {
		err := MigrateUser()

		if err != nil {
			fmt.Println("Error creating iam_user: ", err)
		}
		//	dryrun mode must be handled using transactions

		//	tx := db.Begin()
		//	tx.Rollback()
	},
}

func GetCommand() *cobra.Command {
	return updateCmd
}

func MigrateUser() error {
	con, err := connection.GetConnection()
	if err != nil {
		fmt.Println("Error load connection", err)
		return err
	}

	users := repository.NewUsersRepository(con)
	totalUsersToMigrate, err := users.Total()

	if err != nil {
		fmt.Println("Error get Total iam users", err)
		return err
	}

	fmt.Println("total Users Records: ", totalUsersToMigrate)

	err = users.CreateNewTable()

	if err != nil {
		fmt.Println("Error to create new iam users enty", err)
		return err
	}

	err = users.ExtractIAMUsers()

	if err != nil {
		fmt.Println("Error to extract new iam users enty", err)
		return err
	}

	userKeys := repository.NewUserKeysRepository(con)
	totalUserKeysToMigrate, err := userKeys.Total()

	if err != nil {
		fmt.Println("Error get Total iam users", err)
		return err
	}

	fmt.Println("total User Keys Records: ", totalUserKeysToMigrate)

	err = userKeys.CreateNewTable()

	if err != nil {
		fmt.Println("Error to create new iam users entities", err)
		return err
	}

	err = userKeys.ExtractIAMUserKeys()

	if err != nil {
		fmt.Println("Error to extract new iam users keys entities", err)
		return err
	}

	accessGroup := repository.NewAccessGroupApiKeysRepository(con)
	totalAccessGroupToMigrate, err := accessGroup.Total()

	if err != nil {
		fmt.Println("Error get Total iam access group api key users", err)
		return err
	}

	fmt.Println("total User Keys Records: ", totalAccessGroupToMigrate)

	err = accessGroup.CreateNewTable()

	if err != nil {
		fmt.Println("Error to create new iam users entities", err)
		return err
	}

	err = accessGroup.ExtractIAMAccessGroupApiKeys()

	if err != nil {
		fmt.Println("Error to extract new iam users keys entities", err)
		return err
	}

	fmt.Println("It is done!")

	return nil
}
