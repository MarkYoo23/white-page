package db

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
	Client *ent.Client
}

var instance *DbContext = nil

func GetDbContext() *DbContext {
	if instance == nil {
		instance = &DbContext{Client: nil}
	}

	return instance
}

func (r *DbContext) GenerateSchema() error {
	if r.Client == nil {
		return errors.New("Client is not opened")
	}

	if err := r.Client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	return nil
}

func (r *DbContext) Open() error {
	if r.Client != nil {
		return errors.New("already opened")
	}

	if _, err := os.Stat(".bin"); os.IsNotExist(err) {
		err = os.Mkdir(".bin", 0755)
		if err != nil {
			return fmt.Errorf("failed creating folder: %v", err)
		}
	}

	databasePath := filepath.Join(".bin", "write_page.db")
	client, err := ent.Open("sqlite3", "file:"+databasePath+"?cache=shared&_fk=1")
	if err != nil {
		_ = r.Client.Close()
		return fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	r.Client = client
	return nil
}

func (r *DbContext) Close() error {
	if r.Client == nil {
		return errors.New("not opened")
	}

	err := r.Client.Close()
	if err != nil {
		return fmt.Errorf("failed closing connection to sqlite: %v", err)
	}

	r.Client = nil
	return err
}
