package entries

import (
	"context"
	"white-page/ent"
	"white-page/ent/entry"
	"white-page/internal/db"
)

type EntryService struct {
	dbContext *db.DbContext
}

func NewEntryService(dbContext *db.DbContext) *EntryService {
	return &EntryService{
		dbContext: dbContext,
	}
}

func (r *EntryService) Add(name string, surName string, tel string) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.Create().
		SetName(name).
		SetSurname(surName).
		SetTel(tel).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *EntryService) Get(id int) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.Get(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *EntryService) GetAll() ([]*ent.Entry, error) {
	client := r.dbContext.Client

	entries, err := client.Entry.Query().All(context.Background())

	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *EntryService) Update(id int, name string, surName string, tel string) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.UpdateOneID(id).
		SetName(name).
		SetSurname(surName).
		SetTel(tel).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *EntryService) Delete(id int) error {
	client := r.dbContext.Client

	err := client.Entry.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (r *EntryService) FindByPhone(tel string) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.Query().
		Where(entry.TelIn(tel)).
		First(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}
