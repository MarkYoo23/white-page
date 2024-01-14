package commands

import (
	"github.com/spf13/cobra"
	"white-page/internal/db"
	"white-page/internal/di"
)

func init() {
	rootCmd.AddCommand(dbInitCmd)
}

var dbInitCmd = &cobra.Command{
	Use: "generate-schema",
	Run: func(cmd *cobra.Command, args []string) {
		dbContext, err := di.GetService[db.DbContext]()
		if err != nil {
			panic(err)
		}

		err = dbContext.GenerateSchema()
		if err != nil {
			panic(err)
		}
	},
}
