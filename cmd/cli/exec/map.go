package exec

import (
	"fmt"
	"white-page/ent"
	"white-page/internal/di"
	"white-page/internal/entries"
)

type MapExec struct {
}

func NewMapExec() *MapExec {
	return &MapExec{}
}

func (*MapExec) Execute([]string) error {

	service, err := di.GetService[entries.EntryService]()
	if err != nil {
		return fmt.Errorf("error : %v", err)
	}

	// 배열을 맵으로 바꾸기

	results, err := service.GetAll()
	if err != nil {
		return fmt.Errorf("not found error : %v", err)
	}

	entryMap := make(map[string]*ent.Entry)

	for _, result := range results {
		entryMap[result.Tel] = result
	}

	fmt.Println(fmt.Sprintf("entries : %v", entryMap))

	// 맵을 두개의 슬라이스로

	var keys []string
	var values []*ent.Entry

	for key, value := range entryMap {
		keys = append(keys, key)
		values = append(values, value)
	}

	fmt.Println(fmt.Sprintf("keys : %v", keys))
	fmt.Println(fmt.Sprintf("values : %v", values))

	return nil
}
