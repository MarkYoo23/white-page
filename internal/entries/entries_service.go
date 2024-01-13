package entries

import (
	"context"
	"white-page/ent"
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

func (r *EntryService) AddEntry(name string, surName string, tel string) (*ent.Entry, error) {
	client := r.dbContext.Client

	entry, err := client.Entry.Create().
		SetName(name).
		SetSurname(surName).
		SetTel(tel).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return entry, nil
}
