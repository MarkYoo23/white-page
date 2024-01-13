package Infrastructure

// ent.Open 메서드를 인식하지 못하면 우선 build 하면 오류가 해결 됨

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"white-page/ent"

	_ "github.com/mattn/go-sqlite3"
)

type DbContext struct {
	client *ent.Client
}

func NewDbContext() *DbContext {
	return &DbContext{
		client: nil,
	}
}

func (r *DbContext) GenerateSchema() error {
	if r.client == nil {
		return errors.New("client is not opened")
	}

	if err := r.client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	return nil
}

func (r *DbContext) Open() error {
	if r.client != nil {
		return errors.New("already opened")
	}

	err := os.Mkdir(".bin", 0755)

	if err != nil {
		return fmt.Errorf("failed creating folder: %v", err)
	}

	databasePath := filepath.Join(".bin", "write_page.db")
	client, err := ent.Open("sqlite3", "file:"+databasePath+"?cache=shared&_fk=1")
	if err != nil {
		_ = r.client.Close()
		return fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	r.client = client
	return nil
}

func (r *DbContext) Close() error {
	if r.client == nil {
		return errors.New("not opened")
	}

	err := r.client.Close()
	if err != nil {
		return fmt.Errorf("failed closing connection to sqlite: %v", err)
	}

	r.client = nil
	return err
}
