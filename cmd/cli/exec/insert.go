package exec

import (
	"errors"
	"fmt"
	"regexp"
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

	var name = args[0]
	var surname = args[1]
	var tel = args[2]

	match, err := regexp.Match(`^\d{3}-\d{3,4}-\d{4}$`, []byte(tel))
	if err != nil {
		return err
	}

	if !match {
		return errors.New("tel format error")
	}

	tel = strings.ReplaceAll(tel, "-", "")

	result, err := service.Add(name, surname, tel)
	if err != nil {
		return fmt.Errorf("db error : %v", err)
	}

	fmt.Println(fmt.Sprintf("entry : %v", result))

	return nil
}
