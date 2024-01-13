package di

import (
	"errors"
	"fmt"
	"go.uber.org/dig"
	"white-page/internal/db"
	"white-page/internal/entries"
)

var container *dig.Container

func Initialize() error {
	if container != nil {
		return errors.New("already initialized")
	}

	newContainer := dig.New()

	if err := newContainer.Provide(func() *db.DbContext {
		return db.GetDbContext()
	}); err != nil {
		return fmt.Errorf("register service error : %v", err)
	}

	if err := newContainer.Provide(func(ctx *db.DbContext) *entries.EntryService {
		return entries.NewEntryService(ctx)
	}); err != nil {
		return fmt.Errorf("register service error : %v", err)
	}

	container = newContainer
	return nil
}

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return container.Invoke(function, opts...)
}

func GetService[T any]() (*T, error) {
	var service *T
	if err := container.Invoke(func(s *T) {
		service = s
	}); err != nil {
		return nil, err
	}

	return service, nil
}
