package main

import "white-page/internal/Infrastructure"

func main() {
	dbContext := Infrastructure.NewDbContext()
	err := dbContext.Open()
	if err != nil {
		panic(err)
	}

	err = dbContext.GenerateSchema()
	if err != nil {
		panic(err)
	}

	err = dbContext.Close()
	if err != nil {
		panic(err)
	}
}
