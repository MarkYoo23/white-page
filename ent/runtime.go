// Code generated by ent, DO NOT EDIT.

package ent

import (
	"white-page/ent/entry"
	"white-page/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entryFields := schema.Entry{}.Fields()
	_ = entryFields
	// entryDescName is the schema descriptor for name field.
	entryDescName := entryFields[0].Descriptor()
	// entry.DefaultName holds the default value on creation for the name field.
	entry.DefaultName = entryDescName.Default.(string)
	// entryDescSurname is the schema descriptor for surname field.
	entryDescSurname := entryFields[1].Descriptor()
	// entry.DefaultSurname holds the default value on creation for the surname field.
	entry.DefaultSurname = entryDescSurname.Default.(string)
	// entryDescTel is the schema descriptor for tel field.
	entryDescTel := entryFields[2].Descriptor()
	// entry.DefaultTel holds the default value on creation for the tel field.
	entry.DefaultTel = entryDescTel.Default.(string)
}
