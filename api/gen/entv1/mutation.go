// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"api/gen/entv1/predicate"
	"api/gen/entv1/status"
	"context"
	"fmt"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeStatus = "Status"
)

// StatusMutation represents an operation that mutates the Status nodes in the graph.
type StatusMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Status, error)
	predicates    []predicate.Status
}

var _ ent.Mutation = (*StatusMutation)(nil)

// statusOption allows management of the mutation configuration using functional options.
type statusOption func(*StatusMutation)

// newStatusMutation creates new mutation for the Status entity.
func newStatusMutation(c config, op Op, opts ...statusOption) *StatusMutation {
	m := &StatusMutation{
		config:        c,
		op:            op,
		typ:           TypeStatus,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStatusID sets the ID field of the mutation.
func withStatusID(id int) statusOption {
	return func(m *StatusMutation) {
		var (
			err   error
			once  sync.Once
			value *Status
		)
		m.oldValue = func(ctx context.Context) (*Status, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Status.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStatus sets the old Status of the mutation.
func withStatus(node *Status) statusOption {
	return func(m *StatusMutation) {
		m.oldValue = func(context.Context) (*Status, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StatusMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StatusMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("entv1: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *StatusMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTitle sets the "title" field.
func (m *StatusMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *StatusMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Status entity.
// If the Status object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StatusMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *StatusMutation) ResetTitle() {
	m.title = nil
}

// Op returns the operation name.
func (m *StatusMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Status).
func (m *StatusMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StatusMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.title != nil {
		fields = append(fields, status.FieldTitle)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StatusMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case status.FieldTitle:
		return m.Title()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StatusMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case status.FieldTitle:
		return m.OldTitle(ctx)
	}
	return nil, fmt.Errorf("unknown Status field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StatusMutation) SetField(name string, value ent.Value) error {
	switch name {
	case status.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	}
	return fmt.Errorf("unknown Status field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StatusMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StatusMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StatusMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Status numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StatusMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StatusMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StatusMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Status nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StatusMutation) ResetField(name string) error {
	switch name {
	case status.FieldTitle:
		m.ResetTitle()
		return nil
	}
	return fmt.Errorf("unknown Status field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StatusMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StatusMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StatusMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StatusMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StatusMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StatusMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StatusMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Status unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StatusMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Status edge %s", name)
}
