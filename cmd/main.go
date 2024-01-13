package main

import (
	"bufio"
	"os"
	"strings"
	"white-page/cmd/exec"
	"white-page/internal/db"
	"white-page/internal/di"
)

var execMap map[string]exec.Exec

func init() {
	execMap = make(map[string]exec.Exec)
	execMap["search"] = exec.NewSearchExec()
	execMap["list"] = exec.NewListExec()
	execMap["insert"] = exec.NewInsertExec()
	execMap["delete"] = exec.NewDeleteExec()
	execMap["export"] = exec.NewExportExec()
	execMap["import"] = exec.NewImportExec()
	execMap["map"] = exec.NewMapExec()
}

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

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		splitValues := strings.Split(strings.Replace(input, "\r\n", "", -1), " ")

		ex, ok := execMap[splitValues[0]]
		if ok {
			err = ex.Execute(splitValues[1:])
		} else {
			break
		}
	}

	err = dbContext.Close()
	if err != nil {
		panic(err)
	}
}
