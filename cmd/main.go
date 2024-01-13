package main

import (
	"white-page/internal/db"
	"white-page/internal/di"
	"white-page/internal/entries"
)

func main() {
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

	err = dbContext.GenerateSchema()
	if err != nil {
		panic(err)
	}

	entryService, err := di.GetService[entries.EntryService]()
	if err != nil {
		panic(err)
	}

	entry, err := entryService.AddEntry("John", "Doe", "1234567890")
	if err != nil {
		return
	}

	println(entry)

	err = dbContext.Close()
	if err != nil {
		panic(err)
	}

}
