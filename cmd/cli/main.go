package main

import (
	"fmt"
	"strings"
	"white-page/cmd/cli/commands"
	"white-page/internal/constants"
	"white-page/internal/db"
	"white-page/internal/di"
)

// https://manytools.org/hacker-tools/ascii-banner/

const banner = `

  █████████  ██████████   ██████   █████
 ███░░░░░███░░███░░░░███ ░░██████ ░░███ 
░███    ░░░  ░███   ░░███ ░███░███ ░███ 
░░█████████  ░███    ░███ ░███░░███░███ 
 ░░░░░░░░███ ░███    ░███ ░███ ░░██████ 
 ███    ░███ ░███    ███  ░███  ░░█████ 
░░█████████  ██████████   █████  ░░█████
 ░░░░░░░░░  ░░░░░░░░░░   ░░░░░    ░░░░░ 
-------------------------------------------------------------------------------
Write page application v{ApplicationVersion}, created by 'Saturday Day Night'
-------------------------------------------------------------------------------
`

func init() {
}

func main() {
	fmt.Println(strings.ReplaceAll(banner, "{ApplicationVersion}", constants.Version))

	err := di.Initialize()
	if err != nil {
		panic(err)
	}

	dbContext, err := di.GetService[db.DbContext]()
	if err != nil {
		panic(err)
	}

	err = dbContext.Open()
	if err != nil {
		panic(err)
	}

	err = commands.Execute()
	if err != nil {
		panic(err)
	}

	err = dbContext.Close()
	if err != nil {
		panic(err)
	}
}
