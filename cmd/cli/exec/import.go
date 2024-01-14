package exec

import (
	"fmt"
	"os"
	"strings"
	"white-page/internal/di"
	"white-page/internal/entries"
)

type ImportExec struct {
}

func NewImportExec() *ImportExec {
	return &ImportExec{}
}

func (*ImportExec) Execute([]string) error {

	service, err := di.GetService[entries.EntryService]()
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	openFile, err := os.Open(".bin/export.csv")
	if err != nil {
		return err
	}

	var buf = make([]byte, 1024)

	_, err = openFile.Read(buf)
	if err != nil {
		return err
	}

	err = openFile.Close()
	if err != nil {
		return err
	}

	err = service.DeleteAll()
	if err != nil {
		return fmt.Errorf("db exec fail : %v", err)
	}

	line := strings.Split(string(buf), "\n")

	for _, v := range line {
		split := strings.Split(v, ",")
		if len(split) == 1 {
			break
		}

		_, err = service.Add(split[0], split[1], split[2])
		if err != nil {
			return fmt.Errorf("db error : %v", err)
		}
	}

	fmt.Println("import success")

	return nil
}
