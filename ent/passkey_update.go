// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cloudreve/Cloudreve/v4/ent/passkey"
	"github.com/cloudreve/Cloudreve/v4/ent/predicate"
	"github.com/cloudreve/Cloudreve/v4/ent/user"
	"github.com/go-webauthn/webauthn/webauthn"
)

// PasskeyUpdate is the builder for updating Passkey entities.
type PasskeyUpdate struct {
	config
	hooks    []Hook
	mutation *PasskeyMutation
}

// Where appends a list predicates to the PasskeyUpdate builder.
func (pu *PasskeyUpdate) Where(ps ...predicate.Passkey) *PasskeyUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PasskeyUpdate) SetUpdatedAt(t time.Time) *PasskeyUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PasskeyUpdate) SetDeletedAt(t time.Time) *PasskeyUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PasskeyUpdate) SetNillableDeletedAt(t *time.Time) *PasskeyUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *PasskeyUpdate) ClearDeletedAt() *PasskeyUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PasskeyUpdate) SetUserID(i int) *PasskeyUpdate {
	pu.mutation.SetUserID(i)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PasskeyUpdate) SetNillableUserID(i *int) *PasskeyUpdate {
	if i != nil {
		pu.SetUserID(*i)
	}
	return pu
}

// SetCredentialID sets the "credential_id" field.
func (pu *PasskeyUpdate) SetCredentialID(s string) *PasskeyUpdate {
	pu.mutation.SetCredentialID(s)
	return pu
}

// SetNillableCredentialID sets the "credential_id" field if the given value is not nil.
func (pu *PasskeyUpdate) SetNillableCredentialID(s *string) *PasskeyUpdate {
	if s != nil {
		pu.SetCredentialID(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PasskeyUpdate) SetName(s string) *PasskeyUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PasskeyUpdate) SetNillableName(s *string) *PasskeyUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetCredential sets the "credential" field.
func (pu *PasskeyUpdate) SetCredential(w *webauthn.Credential) *PasskeyUpdate {
	pu.mutation.SetCredential(w)
	return pu
}

// SetUsedAt sets the "used_at" field.
func (pu *PasskeyUpdate) SetUsedAt(t time.Time) *PasskeyUpdate {
	pu.mutation.SetUsedAt(t)
	return pu
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (pu *PasskeyUpdate) SetNillableUsedAt(t *time.Time) *PasskeyUpdate {
	if t != nil {
		pu.SetUsedAt(*t)
	}
	return pu
}

// ClearUsedAt clears the value of the "used_at" field.
func (pu *PasskeyUpdate) ClearUsedAt() *PasskeyUpdate {
	pu.mutation.ClearUsedAt()
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PasskeyUpdate) SetUser(u *User) *PasskeyUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PasskeyMutation object of the builder.
func (pu *PasskeyUpdate) Mutation() *PasskeyMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PasskeyUpdate) ClearUser() *PasskeyUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PasskeyUpdate) Save(ctx context.Context) (int, error) {
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PasskeyUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PasskeyUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PasskeyUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PasskeyUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if passkey.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized passkey.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := passkey.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pu *PasskeyUpdate) check() error {
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Passkey.user"`)
	}
	return nil
}

func (pu *PasskeyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(passkey.Table, passkey.Columns, sqlgraph.NewFieldSpec(passkey.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(passkey.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(passkey.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(passkey.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.CredentialID(); ok {
		_spec.SetField(passkey.FieldCredentialID, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(passkey.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Credential(); ok {
		_spec.SetField(passkey.FieldCredential, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.UsedAt(); ok {
		_spec.SetField(passkey.FieldUsedAt, field.TypeTime, value)
	}
	if pu.mutation.UsedAtCleared() {
		_spec.ClearField(passkey.FieldUsedAt, field.TypeTime)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passkey.UserTable,
			Columns: []string{passkey.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passkey.UserTable,
			Columns: []string{passkey.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passkey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PasskeyUpdateOne is the builder for updating a single Passkey entity.
type PasskeyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasskeyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PasskeyUpdateOne) SetUpdatedAt(t time.Time) *PasskeyUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PasskeyUpdateOne) SetDeletedAt(t time.Time) *PasskeyUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PasskeyUpdateOne) SetNillableDeletedAt(t *time.Time) *PasskeyUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *PasskeyUpdateOne) ClearDeletedAt() *PasskeyUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PasskeyUpdateOne) SetUserID(i int) *PasskeyUpdateOne {
	puo.mutation.SetUserID(i)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PasskeyUpdateOne) SetNillableUserID(i *int) *PasskeyUpdateOne {
	if i != nil {
		puo.SetUserID(*i)
	}
	return puo
}

// SetCredentialID sets the "credential_id" field.
func (puo *PasskeyUpdateOne) SetCredentialID(s string) *PasskeyUpdateOne {
	puo.mutation.SetCredentialID(s)
	return puo
}

// SetNillableCredentialID sets the "credential_id" field if the given value is not nil.
func (puo *PasskeyUpdateOne) SetNillableCredentialID(s *string) *PasskeyUpdateOne {
	if s != nil {
		puo.SetCredentialID(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PasskeyUpdateOne) SetName(s string) *PasskeyUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PasskeyUpdateOne) SetNillableName(s *string) *PasskeyUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetCredential sets the "credential" field.
func (puo *PasskeyUpdateOne) SetCredential(w *webauthn.Credential) *PasskeyUpdateOne {
	puo.mutation.SetCredential(w)
	return puo
}

// SetUsedAt sets the "used_at" field.
func (puo *PasskeyUpdateOne) SetUsedAt(t time.Time) *PasskeyUpdateOne {
	puo.mutation.SetUsedAt(t)
	return puo
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (puo *PasskeyUpdateOne) SetNillableUsedAt(t *time.Time) *PasskeyUpdateOne {
	if t != nil {
		puo.SetUsedAt(*t)
	}
	return puo
}

// ClearUsedAt clears the value of the "used_at" field.
func (puo *PasskeyUpdateOne) ClearUsedAt() *PasskeyUpdateOne {
	puo.mutation.ClearUsedAt()
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PasskeyUpdateOne) SetUser(u *User) *PasskeyUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PasskeyMutation object of the builder.
func (puo *PasskeyUpdateOne) Mutation() *PasskeyMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PasskeyUpdateOne) ClearUser() *PasskeyUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PasskeyUpdate builder.
func (puo *PasskeyUpdateOne) Where(ps ...predicate.Passkey) *PasskeyUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PasskeyUpdateOne) Select(field string, fields ...string) *PasskeyUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Passkey entity.
func (puo *PasskeyUpdateOne) Save(ctx context.Context) (*Passkey, error) {
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PasskeyUpdateOne) SaveX(ctx context.Context) *Passkey {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PasskeyUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PasskeyUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PasskeyUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if passkey.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized passkey.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := passkey.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (puo *PasskeyUpdateOne) check() error {
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Passkey.user"`)
	}
	return nil
}

func (puo *PasskeyUpdateOne) sqlSave(ctx context.Context) (_node *Passkey, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(passkey.Table, passkey.Columns, sqlgraph.NewFieldSpec(passkey.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Passkey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passkey.FieldID)
		for _, f := range fields {
			if !passkey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passkey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(passkey.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(passkey.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(passkey.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.CredentialID(); ok {
		_spec.SetField(passkey.FieldCredentialID, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(passkey.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Credential(); ok {
		_spec.SetField(passkey.FieldCredential, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.UsedAt(); ok {
		_spec.SetField(passkey.FieldUsedAt, field.TypeTime, value)
	}
	if puo.mutation.UsedAtCleared() {
		_spec.ClearField(passkey.FieldUsedAt, field.TypeTime)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passkey.UserTable,
			Columns: []string{passkey.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passkey.UserTable,
			Columns: []string{passkey.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Passkey{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passkey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
