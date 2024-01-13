// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"white-page/ent/entry"
	"white-page/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEntry = "Entry"
)

// EntryMutation represents an operation that mutates the Entry nodes in the graph.
type EntryMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	surname       *string
	tel           *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Entry, error)
	predicates    []predicate.Entry
}

var _ ent.Mutation = (*EntryMutation)(nil)

// entryOption allows management of the mutation configuration using functional options.
type entryOption func(*EntryMutation)

// newEntryMutation creates new mutation for the Entry entity.
func newEntryMutation(c config, op Op, opts ...entryOption) *EntryMutation {
	m := &EntryMutation{
		config:        c,
		op:            op,
		typ:           TypeEntry,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEntryID sets the ID field of the mutation.
func withEntryID(id int) entryOption {
	return func(m *EntryMutation) {
		var (
			err   error
			once  sync.Once
			value *Entry
		)
		m.oldValue = func(ctx context.Context) (*Entry, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Entry.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEntry sets the old Entry of the mutation.
func withEntry(node *Entry) entryOption {
	return func(m *EntryMutation) {
		m.oldValue = func(context.Context) (*Entry, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EntryMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EntryMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EntryMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EntryMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Entry.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *EntryMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *EntryMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Entry entity.
// If the Entry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntryMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *EntryMutation) ResetName() {
	m.name = nil
}

// SetSurname sets the "surname" field.
func (m *EntryMutation) SetSurname(s string) {
	m.surname = &s
}

// Surname returns the value of the "surname" field in the mutation.
func (m *EntryMutation) Surname() (r string, exists bool) {
	v := m.surname
	if v == nil {
		return
	}
	return *v, true
}

// OldSurname returns the old "surname" field's value of the Entry entity.
// If the Entry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntryMutation) OldSurname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSurname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSurname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSurname: %w", err)
	}
	return oldValue.Surname, nil
}

// ResetSurname resets all changes to the "surname" field.
func (m *EntryMutation) ResetSurname() {
	m.surname = nil
}

// SetTel sets the "tel" field.
func (m *EntryMutation) SetTel(s string) {
	m.tel = &s
}

// Tel returns the value of the "tel" field in the mutation.
func (m *EntryMutation) Tel() (r string, exists bool) {
	v := m.tel
	if v == nil {
		return
	}
	return *v, true
}

// OldTel returns the old "tel" field's value of the Entry entity.
// If the Entry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntryMutation) OldTel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTel: %w", err)
	}
	return oldValue.Tel, nil
}

// ResetTel resets all changes to the "tel" field.
func (m *EntryMutation) ResetTel() {
	m.tel = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *EntryMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EntryMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Entry entity.
// If the Entry object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntryMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EntryMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the EntryMutation builder.
func (m *EntryMutation) Where(ps ...predicate.Entry) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EntryMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EntryMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Entry, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EntryMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EntryMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Entry).
func (m *EntryMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EntryMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, entry.FieldName)
	}
	if m.surname != nil {
		fields = append(fields, entry.FieldSurname)
	}
	if m.tel != nil {
		fields = append(fields, entry.FieldTel)
	}
	if m.created_at != nil {
		fields = append(fields, entry.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EntryMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case entry.FieldName:
		return m.Name()
	case entry.FieldSurname:
		return m.Surname()
	case entry.FieldTel:
		return m.Tel()
	case entry.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EntryMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case entry.FieldName:
		return m.OldName(ctx)
	case entry.FieldSurname:
		return m.OldSurname(ctx)
	case entry.FieldTel:
		return m.OldTel(ctx)
	case entry.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Entry field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EntryMutation) SetField(name string, value ent.Value) error {
	switch name {
	case entry.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case entry.FieldSurname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSurname(v)
		return nil
	case entry.FieldTel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTel(v)
		return nil
	case entry.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Entry field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EntryMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EntryMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EntryMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Entry numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EntryMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EntryMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EntryMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Entry nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EntryMutation) ResetField(name string) error {
	switch name {
	case entry.FieldName:
		m.ResetName()
		return nil
	case entry.FieldSurname:
		m.ResetSurname()
		return nil
	case entry.FieldTel:
		m.ResetTel()
		return nil
	case entry.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Entry field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EntryMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EntryMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EntryMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EntryMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EntryMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EntryMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EntryMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Entry unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EntryMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Entry edge %s", name)
}
