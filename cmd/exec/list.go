package exec

import (
	"fmt"
	"white-page/internal/di"
	"white-page/internal/entries"
)

type ListExec struct {
}

func NewListExec() *ListExec {
	return &ListExec{}
}

func (*ListExec) Execute([]string) error {
	service, err := di.GetService[entries.EntryService]()
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	results, err := service.GetAll()
	if err != nil {
		return fmt.Errorf("not found error : %v", err)
	}

	fmt.Println(fmt.Sprintf("entries : %v", results))

	return nil
}
