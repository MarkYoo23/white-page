package exec

import (
	"fmt"
	"strconv"
	"white-page/internal/di"
	"white-page/internal/entries"
)

type DeleteExec struct {
}

func NewDeleteExec() *DeleteExec {
	return &DeleteExec{}
}

func (*DeleteExec) Execute(args []string) error {
	service, err := di.GetService[entries.EntryService]()
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	idx, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("convert string to int error : %v", err)
	}

	err = service.Delete(idx)
	if err != nil {
		return fmt.Errorf("db error : %v", err)
	}

	fmt.Println("entry deleted")
	return nil
}
