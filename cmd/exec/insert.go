package exec

import (
	"fmt"
	"strings"
	"white-page/internal/di"
	"white-page/internal/entries"
)

type InsertExec struct {
}

func NewInsertExec() *InsertExec {
	return &InsertExec{}
}

func (*InsertExec) Execute(args []string) error {
	service, err := di.GetService[entries.EntryService]()
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	result, err := service.Add(args[0], args[1], strings.ReplaceAll(args[2], "-", ""))
	if err != nil {
		return fmt.Errorf("db error : %v", err)
	}

	fmt.Println(fmt.Sprintf("entry : %v", result))

	return nil
}
