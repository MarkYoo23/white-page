// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EntriesColumns holds the columns for the "entries" table.
	EntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "surname", Type: field.TypeString, Default: "unknown"},
		{Name: "tel", Type: field.TypeString, Default: "unknown"},
	}
	// EntriesTable holds the schema information for the "entries" table.
	EntriesTable = &schema.Table{
		Name:       "entries",
		Columns:    EntriesColumns,
		PrimaryKey: []*schema.Column{EntriesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EntriesTable,
	}
)

func init() {
}