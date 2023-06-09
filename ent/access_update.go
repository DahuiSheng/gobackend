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
	"github.com/DahuiSheng/gobackend/ent/access"
	"github.com/DahuiSheng/gobackend/ent/predicate"
)

// AccessUpdate is the builder for updating Access entities.
type AccessUpdate struct {
	config
	hooks    []Hook
	mutation *AccessMutation
}

// Where appends a list predicates to the AccessUpdate builder.
func (au *AccessUpdate) Where(ps ...predicate.Access) *AccessUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdateTime sets the "update_time" field.
func (au *AccessUpdate) SetUpdateTime(t time.Time) *AccessUpdate {
	au.mutation.SetUpdateTime(t)
	return au
}

// SetName sets the "name" field.
func (au *AccessUpdate) SetName(s string) *AccessUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AccessUpdate) SetNillableName(s *string) *AccessUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetCheckIn sets the "check_in" field.
func (au *AccessUpdate) SetCheckIn(t time.Time) *AccessUpdate {
	au.mutation.SetCheckIn(t)
	return au
}

// SetNillableCheckIn sets the "check_in" field if the given value is not nil.
func (au *AccessUpdate) SetNillableCheckIn(t *time.Time) *AccessUpdate {
	if t != nil {
		au.SetCheckIn(*t)
	}
	return au
}

// SetCheckOut sets the "check_out" field.
func (au *AccessUpdate) SetCheckOut(t time.Time) *AccessUpdate {
	au.mutation.SetCheckOut(t)
	return au
}

// SetNillableCheckOut sets the "check_out" field if the given value is not nil.
func (au *AccessUpdate) SetNillableCheckOut(t *time.Time) *AccessUpdate {
	if t != nil {
		au.SetCheckOut(*t)
	}
	return au
}

// ClearCheckOut clears the value of the "check_out" field.
func (au *AccessUpdate) ClearCheckOut() *AccessUpdate {
	au.mutation.ClearCheckOut()
	return au
}

// Mutation returns the AccessMutation object of the builder.
func (au *AccessUpdate) Mutation() *AccessMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccessUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks[int, AccessMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccessUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccessUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccessUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccessUpdate) defaults() {
	if _, ok := au.mutation.UpdateTime(); !ok {
		v := access.UpdateDefaultUpdateTime()
		au.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccessUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := access.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Access.name": %w`, err)}
		}
	}
	return nil
}

func (au *AccessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdateTime(); ok {
		_spec.SetField(access.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(access.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.CheckIn(); ok {
		_spec.SetField(access.FieldCheckIn, field.TypeTime, value)
	}
	if value, ok := au.mutation.CheckOut(); ok {
		_spec.SetField(access.FieldCheckOut, field.TypeTime, value)
	}
	if au.mutation.CheckOutCleared() {
		_spec.ClearField(access.FieldCheckOut, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccessUpdateOne is the builder for updating a single Access entity.
type AccessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessMutation
}

// SetUpdateTime sets the "update_time" field.
func (auo *AccessUpdateOne) SetUpdateTime(t time.Time) *AccessUpdateOne {
	auo.mutation.SetUpdateTime(t)
	return auo
}

// SetName sets the "name" field.
func (auo *AccessUpdateOne) SetName(s string) *AccessUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableName(s *string) *AccessUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetCheckIn sets the "check_in" field.
func (auo *AccessUpdateOne) SetCheckIn(t time.Time) *AccessUpdateOne {
	auo.mutation.SetCheckIn(t)
	return auo
}

// SetNillableCheckIn sets the "check_in" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableCheckIn(t *time.Time) *AccessUpdateOne {
	if t != nil {
		auo.SetCheckIn(*t)
	}
	return auo
}

// SetCheckOut sets the "check_out" field.
func (auo *AccessUpdateOne) SetCheckOut(t time.Time) *AccessUpdateOne {
	auo.mutation.SetCheckOut(t)
	return auo
}

// SetNillableCheckOut sets the "check_out" field if the given value is not nil.
func (auo *AccessUpdateOne) SetNillableCheckOut(t *time.Time) *AccessUpdateOne {
	if t != nil {
		auo.SetCheckOut(*t)
	}
	return auo
}

// ClearCheckOut clears the value of the "check_out" field.
func (auo *AccessUpdateOne) ClearCheckOut() *AccessUpdateOne {
	auo.mutation.ClearCheckOut()
	return auo
}

// Mutation returns the AccessMutation object of the builder.
func (auo *AccessUpdateOne) Mutation() *AccessMutation {
	return auo.mutation
}

// Where appends a list predicates to the AccessUpdate builder.
func (auo *AccessUpdateOne) Where(ps ...predicate.Access) *AccessUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccessUpdateOne) Select(field string, fields ...string) *AccessUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Access entity.
func (auo *AccessUpdateOne) Save(ctx context.Context) (*Access, error) {
	auo.defaults()
	return withHooks[*Access, AccessMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccessUpdateOne) SaveX(ctx context.Context) *Access {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccessUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccessUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccessUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdateTime(); !ok {
		v := access.UpdateDefaultUpdateTime()
		auo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccessUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := access.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Access.name": %w`, err)}
		}
	}
	return nil
}

func (auo *AccessUpdateOne) sqlSave(ctx context.Context) (_node *Access, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(access.Table, access.Columns, sqlgraph.NewFieldSpec(access.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Access.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, access.FieldID)
		for _, f := range fields {
			if !access.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != access.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdateTime(); ok {
		_spec.SetField(access.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(access.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.CheckIn(); ok {
		_spec.SetField(access.FieldCheckIn, field.TypeTime, value)
	}
	if value, ok := auo.mutation.CheckOut(); ok {
		_spec.SetField(access.FieldCheckOut, field.TypeTime, value)
	}
	if auo.mutation.CheckOutCleared() {
		_spec.ClearField(access.FieldCheckOut, field.TypeTime)
	}
	_node = &Access{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{access.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
