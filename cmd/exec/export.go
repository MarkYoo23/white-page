package exec

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"white-page/internal/di"
	"white-page/internal/entries"
)

type ExportExec struct {
}

func NewExportExec() *ExportExec {
	return &ExportExec{}
}

func (*ExportExec) Execute([]string) error {

	service, err := di.GetService[entries.EntryService]()
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	results, err := service.GetAll()
	if err != nil {
		return fmt.Errorf("not found error : %v", err)
	}

	// CSV 파일 생성
	file, err := os.Create(fmt.Sprintf(".bin/export_%d.csv", time.Now().Unix()))
	if err != nil {
		return fmt.Errorf("could not create CSV file: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, result := range results {
		if err := writer.Write([]string{result.Name, result.Surname, result.Tel}); err != nil {
			return fmt.Errorf("could not write to CSV file: %v", err)
		}
	}

	return nil
}
