package exec

import (
	"fmt"
	"white-page/internal/di"
	"white-page/internal/entries"
)

type SearchExec struct {
}

func NewSearchExec() *SearchExec {
	return &SearchExec{}
}

func (*SearchExec) Execute(args []string) error {
	service, err := di.GetService[entries.EntryService]()
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	surName := args[0]

	result, err := service.FindByPhone(surName)
	if err != nil {
		return fmt.Errorf("not found error : %v", err)
	}

	fmt.Println(fmt.Sprintf("entry : %v", result))

	return nil
}
