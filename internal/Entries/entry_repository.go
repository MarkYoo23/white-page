package Entries

import "white-page/ent"

type EntryRepository struct {
}

func NewEntryRepository() *EntryRepository {
	return &EntryRepository{}
}

func (r *EntryRepository) Save(entry *ent.Entry) {
}

func (r *EntryRepository) Find(key string) []*ent.Entry {
	return []*ent.Entry{}
}
