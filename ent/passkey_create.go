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
	"github.com/cloudreve/Cloudreve/v4/ent/user"
	"github.com/go-webauthn/webauthn/webauthn"
)

// PasskeyCreate is the builder for creating a Passkey entity.
type PasskeyCreate struct {
	config
	mutation *PasskeyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pc *PasskeyCreate) SetCreatedAt(t time.Time) *PasskeyCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PasskeyCreate) SetNillableCreatedAt(t *time.Time) *PasskeyCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PasskeyCreate) SetUpdatedAt(t time.Time) *PasskeyCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PasskeyCreate) SetNillableUpdatedAt(t *time.Time) *PasskeyCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PasskeyCreate) SetDeletedAt(t time.Time) *PasskeyCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PasskeyCreate) SetNillableDeletedAt(t *time.Time) *PasskeyCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *PasskeyCreate) SetUserID(i int) *PasskeyCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetCredentialID sets the "credential_id" field.
func (pc *PasskeyCreate) SetCredentialID(s string) *PasskeyCreate {
	pc.mutation.SetCredentialID(s)
	return pc
}

// SetName sets the "name" field.
func (pc *PasskeyCreate) SetName(s string) *PasskeyCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetCredential sets the "credential" field.
func (pc *PasskeyCreate) SetCredential(w *webauthn.Credential) *PasskeyCreate {
	pc.mutation.SetCredential(w)
	return pc
}

// SetUsedAt sets the "used_at" field.
func (pc *PasskeyCreate) SetUsedAt(t time.Time) *PasskeyCreate {
	pc.mutation.SetUsedAt(t)
	return pc
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (pc *PasskeyCreate) SetNillableUsedAt(t *time.Time) *PasskeyCreate {
	if t != nil {
		pc.SetUsedAt(*t)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PasskeyCreate) SetUser(u *User) *PasskeyCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the PasskeyMutation object of the builder.
func (pc *PasskeyCreate) Mutation() *PasskeyMutation {
	return pc.mutation
}

// Save creates the Passkey in the database.
func (pc *PasskeyCreate) Save(ctx context.Context) (*Passkey, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PasskeyCreate) SaveX(ctx context.Context) *Passkey {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PasskeyCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PasskeyCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PasskeyCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if passkey.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized passkey.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := passkey.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if passkey.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized passkey.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := passkey.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PasskeyCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Passkey.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Passkey.updated_at"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Passkey.user_id"`)}
	}
	if _, ok := pc.mutation.CredentialID(); !ok {
		return &ValidationError{Name: "credential_id", err: errors.New(`ent: missing required field "Passkey.credential_id"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Passkey.name"`)}
	}
	if _, ok := pc.mutation.Credential(); !ok {
		return &ValidationError{Name: "credential", err: errors.New(`ent: missing required field "Passkey.credential"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Passkey.user"`)}
	}
	return nil
}

func (pc *PasskeyCreate) sqlSave(ctx context.Context) (*Passkey, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PasskeyCreate) createSpec() (*Passkey, *sqlgraph.CreateSpec) {
	var (
		_node = &Passkey{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(passkey.Table, sqlgraph.NewFieldSpec(passkey.FieldID, field.TypeInt))
	)

	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		id64 := int64(id)
		_spec.ID.Value = id64
	}

	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(passkey.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(passkey.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(passkey.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := pc.mutation.CredentialID(); ok {
		_spec.SetField(passkey.FieldCredentialID, field.TypeString, value)
		_node.CredentialID = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(passkey.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Credential(); ok {
		_spec.SetField(passkey.FieldCredential, field.TypeJSON, value)
		_node.Credential = value
	}
	if value, ok := pc.mutation.UsedAt(); ok {
		_spec.SetField(passkey.FieldUsedAt, field.TypeTime, value)
		_node.UsedAt = &value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Passkey.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PasskeyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pc *PasskeyCreate) OnConflict(opts ...sql.ConflictOption) *PasskeyUpsertOne {
	pc.conflict = opts
	return &PasskeyUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Passkey.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PasskeyCreate) OnConflictColumns(columns ...string) *PasskeyUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PasskeyUpsertOne{
		create: pc,
	}
}

type (
	// PasskeyUpsertOne is the builder for "upsert"-ing
	//  one Passkey node.
	PasskeyUpsertOne struct {
		create *PasskeyCreate
	}

	// PasskeyUpsert is the "OnConflict" setter.
	PasskeyUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *PasskeyUpsert) SetUpdatedAt(v time.Time) *PasskeyUpsert {
	u.Set(passkey.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PasskeyUpsert) UpdateUpdatedAt() *PasskeyUpsert {
	u.SetExcluded(passkey.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PasskeyUpsert) SetDeletedAt(v time.Time) *PasskeyUpsert {
	u.Set(passkey.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PasskeyUpsert) UpdateDeletedAt() *PasskeyUpsert {
	u.SetExcluded(passkey.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *PasskeyUpsert) ClearDeletedAt() *PasskeyUpsert {
	u.SetNull(passkey.FieldDeletedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *PasskeyUpsert) SetUserID(v int) *PasskeyUpsert {
	u.Set(passkey.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *PasskeyUpsert) UpdateUserID() *PasskeyUpsert {
	u.SetExcluded(passkey.FieldUserID)
	return u
}

// SetCredentialID sets the "credential_id" field.
func (u *PasskeyUpsert) SetCredentialID(v string) *PasskeyUpsert {
	u.Set(passkey.FieldCredentialID, v)
	return u
}

// UpdateCredentialID sets the "credential_id" field to the value that was provided on create.
func (u *PasskeyUpsert) UpdateCredentialID() *PasskeyUpsert {
	u.SetExcluded(passkey.FieldCredentialID)
	return u
}

// SetName sets the "name" field.
func (u *PasskeyUpsert) SetName(v string) *PasskeyUpsert {
	u.Set(passkey.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PasskeyUpsert) UpdateName() *PasskeyUpsert {
	u.SetExcluded(passkey.FieldName)
	return u
}

// SetCredential sets the "credential" field.
func (u *PasskeyUpsert) SetCredential(v *webauthn.Credential) *PasskeyUpsert {
	u.Set(passkey.FieldCredential, v)
	return u
}

// UpdateCredential sets the "credential" field to the value that was provided on create.
func (u *PasskeyUpsert) UpdateCredential() *PasskeyUpsert {
	u.SetExcluded(passkey.FieldCredential)
	return u
}

// SetUsedAt sets the "used_at" field.
func (u *PasskeyUpsert) SetUsedAt(v time.Time) *PasskeyUpsert {
	u.Set(passkey.FieldUsedAt, v)
	return u
}

// UpdateUsedAt sets the "used_at" field to the value that was provided on create.
func (u *PasskeyUpsert) UpdateUsedAt() *PasskeyUpsert {
	u.SetExcluded(passkey.FieldUsedAt)
	return u
}

// ClearUsedAt clears the value of the "used_at" field.
func (u *PasskeyUpsert) ClearUsedAt() *PasskeyUpsert {
	u.SetNull(passkey.FieldUsedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Passkey.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PasskeyUpsertOne) UpdateNewValues() *PasskeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(passkey.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Passkey.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PasskeyUpsertOne) Ignore() *PasskeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PasskeyUpsertOne) DoNothing() *PasskeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PasskeyCreate.OnConflict
// documentation for more info.
func (u *PasskeyUpsertOne) Update(set func(*PasskeyUpsert)) *PasskeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PasskeyUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PasskeyUpsertOne) SetUpdatedAt(v time.Time) *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PasskeyUpsertOne) UpdateUpdatedAt() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PasskeyUpsertOne) SetDeletedAt(v time.Time) *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PasskeyUpsertOne) UpdateDeletedAt() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *PasskeyUpsertOne) ClearDeletedAt() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *PasskeyUpsertOne) SetUserID(v int) *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *PasskeyUpsertOne) UpdateUserID() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateUserID()
	})
}

// SetCredentialID sets the "credential_id" field.
func (u *PasskeyUpsertOne) SetCredentialID(v string) *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetCredentialID(v)
	})
}

// UpdateCredentialID sets the "credential_id" field to the value that was provided on create.
func (u *PasskeyUpsertOne) UpdateCredentialID() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateCredentialID()
	})
}

// SetName sets the "name" field.
func (u *PasskeyUpsertOne) SetName(v string) *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PasskeyUpsertOne) UpdateName() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateName()
	})
}

// SetCredential sets the "credential" field.
func (u *PasskeyUpsertOne) SetCredential(v *webauthn.Credential) *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetCredential(v)
	})
}

// UpdateCredential sets the "credential" field to the value that was provided on create.
func (u *PasskeyUpsertOne) UpdateCredential() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateCredential()
	})
}

// SetUsedAt sets the "used_at" field.
func (u *PasskeyUpsertOne) SetUsedAt(v time.Time) *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetUsedAt(v)
	})
}

// UpdateUsedAt sets the "used_at" field to the value that was provided on create.
func (u *PasskeyUpsertOne) UpdateUsedAt() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateUsedAt()
	})
}

// ClearUsedAt clears the value of the "used_at" field.
func (u *PasskeyUpsertOne) ClearUsedAt() *PasskeyUpsertOne {
	return u.Update(func(s *PasskeyUpsert) {
		s.ClearUsedAt()
	})
}

// Exec executes the query.
func (u *PasskeyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PasskeyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PasskeyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PasskeyUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PasskeyUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (m *PasskeyCreate) SetRawID(t int) *PasskeyCreate {
	m.mutation.SetRawID(t)
	return m
}

// PasskeyCreateBulk is the builder for creating many Passkey entities in bulk.
type PasskeyCreateBulk struct {
	config
	err      error
	builders []*PasskeyCreate
	conflict []sql.ConflictOption
}

// Save creates the Passkey entities in the database.
func (pcb *PasskeyCreateBulk) Save(ctx context.Context) ([]*Passkey, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Passkey, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PasskeyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PasskeyCreateBulk) SaveX(ctx context.Context) []*Passkey {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PasskeyCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PasskeyCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Passkey.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PasskeyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pcb *PasskeyCreateBulk) OnConflict(opts ...sql.ConflictOption) *PasskeyUpsertBulk {
	pcb.conflict = opts
	return &PasskeyUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Passkey.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PasskeyCreateBulk) OnConflictColumns(columns ...string) *PasskeyUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PasskeyUpsertBulk{
		create: pcb,
	}
}

// PasskeyUpsertBulk is the builder for "upsert"-ing
// a bulk of Passkey nodes.
type PasskeyUpsertBulk struct {
	create *PasskeyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Passkey.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PasskeyUpsertBulk) UpdateNewValues() *PasskeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(passkey.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Passkey.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PasskeyUpsertBulk) Ignore() *PasskeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PasskeyUpsertBulk) DoNothing() *PasskeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PasskeyCreateBulk.OnConflict
// documentation for more info.
func (u *PasskeyUpsertBulk) Update(set func(*PasskeyUpsert)) *PasskeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PasskeyUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PasskeyUpsertBulk) SetUpdatedAt(v time.Time) *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PasskeyUpsertBulk) UpdateUpdatedAt() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PasskeyUpsertBulk) SetDeletedAt(v time.Time) *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PasskeyUpsertBulk) UpdateDeletedAt() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *PasskeyUpsertBulk) ClearDeletedAt() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *PasskeyUpsertBulk) SetUserID(v int) *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *PasskeyUpsertBulk) UpdateUserID() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateUserID()
	})
}

// SetCredentialID sets the "credential_id" field.
func (u *PasskeyUpsertBulk) SetCredentialID(v string) *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetCredentialID(v)
	})
}

// UpdateCredentialID sets the "credential_id" field to the value that was provided on create.
func (u *PasskeyUpsertBulk) UpdateCredentialID() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateCredentialID()
	})
}

// SetName sets the "name" field.
func (u *PasskeyUpsertBulk) SetName(v string) *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PasskeyUpsertBulk) UpdateName() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateName()
	})
}

// SetCredential sets the "credential" field.
func (u *PasskeyUpsertBulk) SetCredential(v *webauthn.Credential) *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetCredential(v)
	})
}

// UpdateCredential sets the "credential" field to the value that was provided on create.
func (u *PasskeyUpsertBulk) UpdateCredential() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateCredential()
	})
}

// SetUsedAt sets the "used_at" field.
func (u *PasskeyUpsertBulk) SetUsedAt(v time.Time) *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.SetUsedAt(v)
	})
}

// UpdateUsedAt sets the "used_at" field to the value that was provided on create.
func (u *PasskeyUpsertBulk) UpdateUsedAt() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.UpdateUsedAt()
	})
}

// ClearUsedAt clears the value of the "used_at" field.
func (u *PasskeyUpsertBulk) ClearUsedAt() *PasskeyUpsertBulk {
	return u.Update(func(s *PasskeyUpsert) {
		s.ClearUsedAt()
	})
}

// Exec executes the query.
func (u *PasskeyUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PasskeyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PasskeyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PasskeyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
