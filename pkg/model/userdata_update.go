// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/grin-ch/grin-auth/pkg/model/predicate"
	"github.com/grin-ch/grin-auth/pkg/model/user"
	"github.com/grin-ch/grin-auth/pkg/model/userdata"
)

// UserDataUpdate is the builder for updating UserData entities.
type UserDataUpdate struct {
	config
	hooks    []Hook
	mutation *UserDataMutation
}

// Where appends a list predicates to the UserDataUpdate builder.
func (udu *UserDataUpdate) Where(ps ...predicate.UserData) *UserDataUpdate {
	udu.mutation.Where(ps...)
	return udu
}

// SetNickname sets the "nickname" field.
func (udu *UserDataUpdate) SetNickname(s string) *UserDataUpdate {
	udu.mutation.SetNickname(s)
	return udu
}

// SetAvatarURL sets the "avatar_url" field.
func (udu *UserDataUpdate) SetAvatarURL(s string) *UserDataUpdate {
	udu.mutation.SetAvatarURL(s)
	return udu
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (udu *UserDataUpdate) SetNillableAvatarURL(s *string) *UserDataUpdate {
	if s != nil {
		udu.SetAvatarURL(*s)
	}
	return udu
}

// SetBirthday sets the "birthday" field.
func (udu *UserDataUpdate) SetBirthday(t time.Time) *UserDataUpdate {
	udu.mutation.SetBirthday(t)
	return udu
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (udu *UserDataUpdate) SetNillableBirthday(t *time.Time) *UserDataUpdate {
	if t != nil {
		udu.SetBirthday(*t)
	}
	return udu
}

// SetSex sets the "sex" field.
func (udu *UserDataUpdate) SetSex(u userdata.Sex) *UserDataUpdate {
	udu.mutation.SetSex(u)
	return udu
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (udu *UserDataUpdate) SetNillableSex(u *userdata.Sex) *UserDataUpdate {
	if u != nil {
		udu.SetSex(*u)
	}
	return udu
}

// SetLevelExp sets the "level_exp" field.
func (udu *UserDataUpdate) SetLevelExp(i int) *UserDataUpdate {
	udu.mutation.ResetLevelExp()
	udu.mutation.SetLevelExp(i)
	return udu
}

// SetNillableLevelExp sets the "level_exp" field if the given value is not nil.
func (udu *UserDataUpdate) SetNillableLevelExp(i *int) *UserDataUpdate {
	if i != nil {
		udu.SetLevelExp(*i)
	}
	return udu
}

// AddLevelExp adds i to the "level_exp" field.
func (udu *UserDataUpdate) AddLevelExp(i int) *UserDataUpdate {
	udu.mutation.AddLevelExp(i)
	return udu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (udu *UserDataUpdate) SetUserID(id int) *UserDataUpdate {
	udu.mutation.SetUserID(id)
	return udu
}

// SetUser sets the "user" edge to the User entity.
func (udu *UserDataUpdate) SetUser(u *User) *UserDataUpdate {
	return udu.SetUserID(u.ID)
}

// Mutation returns the UserDataMutation object of the builder.
func (udu *UserDataUpdate) Mutation() *UserDataMutation {
	return udu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (udu *UserDataUpdate) ClearUser() *UserDataUpdate {
	udu.mutation.ClearUser()
	return udu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (udu *UserDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserDataMutation](ctx, udu.sqlSave, udu.mutation, udu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (udu *UserDataUpdate) SaveX(ctx context.Context) int {
	affected, err := udu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (udu *UserDataUpdate) Exec(ctx context.Context) error {
	_, err := udu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udu *UserDataUpdate) ExecX(ctx context.Context) {
	if err := udu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (udu *UserDataUpdate) check() error {
	if v, ok := udu.mutation.Nickname(); ok {
		if err := userdata.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`model: validator failed for field "UserData.nickname": %w`, err)}
		}
	}
	if v, ok := udu.mutation.Sex(); ok {
		if err := userdata.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`model: validator failed for field "UserData.sex": %w`, err)}
		}
	}
	if _, ok := udu.mutation.UserID(); udu.mutation.UserCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "UserData.user"`)
	}
	return nil
}

func (udu *UserDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := udu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userdata.Table,
			Columns: userdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userdata.FieldID,
			},
		},
	}
	if ps := udu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := udu.mutation.Nickname(); ok {
		_spec.SetField(userdata.FieldNickname, field.TypeString, value)
	}
	if value, ok := udu.mutation.AvatarURL(); ok {
		_spec.SetField(userdata.FieldAvatarURL, field.TypeString, value)
	}
	if value, ok := udu.mutation.Birthday(); ok {
		_spec.SetField(userdata.FieldBirthday, field.TypeTime, value)
	}
	if value, ok := udu.mutation.Sex(); ok {
		_spec.SetField(userdata.FieldSex, field.TypeEnum, value)
	}
	if value, ok := udu.mutation.LevelExp(); ok {
		_spec.SetField(userdata.FieldLevelExp, field.TypeInt, value)
	}
	if value, ok := udu.mutation.AddedLevelExp(); ok {
		_spec.AddField(userdata.FieldLevelExp, field.TypeInt, value)
	}
	if udu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userdata.UserTable,
			Columns: []string{userdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := udu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userdata.UserTable,
			Columns: []string{userdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, udu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	udu.mutation.done = true
	return n, nil
}

// UserDataUpdateOne is the builder for updating a single UserData entity.
type UserDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserDataMutation
}

// SetNickname sets the "nickname" field.
func (uduo *UserDataUpdateOne) SetNickname(s string) *UserDataUpdateOne {
	uduo.mutation.SetNickname(s)
	return uduo
}

// SetAvatarURL sets the "avatar_url" field.
func (uduo *UserDataUpdateOne) SetAvatarURL(s string) *UserDataUpdateOne {
	uduo.mutation.SetAvatarURL(s)
	return uduo
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uduo *UserDataUpdateOne) SetNillableAvatarURL(s *string) *UserDataUpdateOne {
	if s != nil {
		uduo.SetAvatarURL(*s)
	}
	return uduo
}

// SetBirthday sets the "birthday" field.
func (uduo *UserDataUpdateOne) SetBirthday(t time.Time) *UserDataUpdateOne {
	uduo.mutation.SetBirthday(t)
	return uduo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uduo *UserDataUpdateOne) SetNillableBirthday(t *time.Time) *UserDataUpdateOne {
	if t != nil {
		uduo.SetBirthday(*t)
	}
	return uduo
}

// SetSex sets the "sex" field.
func (uduo *UserDataUpdateOne) SetSex(u userdata.Sex) *UserDataUpdateOne {
	uduo.mutation.SetSex(u)
	return uduo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (uduo *UserDataUpdateOne) SetNillableSex(u *userdata.Sex) *UserDataUpdateOne {
	if u != nil {
		uduo.SetSex(*u)
	}
	return uduo
}

// SetLevelExp sets the "level_exp" field.
func (uduo *UserDataUpdateOne) SetLevelExp(i int) *UserDataUpdateOne {
	uduo.mutation.ResetLevelExp()
	uduo.mutation.SetLevelExp(i)
	return uduo
}

// SetNillableLevelExp sets the "level_exp" field if the given value is not nil.
func (uduo *UserDataUpdateOne) SetNillableLevelExp(i *int) *UserDataUpdateOne {
	if i != nil {
		uduo.SetLevelExp(*i)
	}
	return uduo
}

// AddLevelExp adds i to the "level_exp" field.
func (uduo *UserDataUpdateOne) AddLevelExp(i int) *UserDataUpdateOne {
	uduo.mutation.AddLevelExp(i)
	return uduo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uduo *UserDataUpdateOne) SetUserID(id int) *UserDataUpdateOne {
	uduo.mutation.SetUserID(id)
	return uduo
}

// SetUser sets the "user" edge to the User entity.
func (uduo *UserDataUpdateOne) SetUser(u *User) *UserDataUpdateOne {
	return uduo.SetUserID(u.ID)
}

// Mutation returns the UserDataMutation object of the builder.
func (uduo *UserDataUpdateOne) Mutation() *UserDataMutation {
	return uduo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uduo *UserDataUpdateOne) ClearUser() *UserDataUpdateOne {
	uduo.mutation.ClearUser()
	return uduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uduo *UserDataUpdateOne) Select(field string, fields ...string) *UserDataUpdateOne {
	uduo.fields = append([]string{field}, fields...)
	return uduo
}

// Save executes the query and returns the updated UserData entity.
func (uduo *UserDataUpdateOne) Save(ctx context.Context) (*UserData, error) {
	return withHooks[*UserData, UserDataMutation](ctx, uduo.sqlSave, uduo.mutation, uduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uduo *UserDataUpdateOne) SaveX(ctx context.Context) *UserData {
	node, err := uduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uduo *UserDataUpdateOne) Exec(ctx context.Context) error {
	_, err := uduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uduo *UserDataUpdateOne) ExecX(ctx context.Context) {
	if err := uduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uduo *UserDataUpdateOne) check() error {
	if v, ok := uduo.mutation.Nickname(); ok {
		if err := userdata.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`model: validator failed for field "UserData.nickname": %w`, err)}
		}
	}
	if v, ok := uduo.mutation.Sex(); ok {
		if err := userdata.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`model: validator failed for field "UserData.sex": %w`, err)}
		}
	}
	if _, ok := uduo.mutation.UserID(); uduo.mutation.UserCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "UserData.user"`)
	}
	return nil
}

func (uduo *UserDataUpdateOne) sqlSave(ctx context.Context) (_node *UserData, err error) {
	if err := uduo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userdata.Table,
			Columns: userdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userdata.FieldID,
			},
		},
	}
	id, ok := uduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "UserData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userdata.FieldID)
		for _, f := range fields {
			if !userdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != userdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uduo.mutation.Nickname(); ok {
		_spec.SetField(userdata.FieldNickname, field.TypeString, value)
	}
	if value, ok := uduo.mutation.AvatarURL(); ok {
		_spec.SetField(userdata.FieldAvatarURL, field.TypeString, value)
	}
	if value, ok := uduo.mutation.Birthday(); ok {
		_spec.SetField(userdata.FieldBirthday, field.TypeTime, value)
	}
	if value, ok := uduo.mutation.Sex(); ok {
		_spec.SetField(userdata.FieldSex, field.TypeEnum, value)
	}
	if value, ok := uduo.mutation.LevelExp(); ok {
		_spec.SetField(userdata.FieldLevelExp, field.TypeInt, value)
	}
	if value, ok := uduo.mutation.AddedLevelExp(); ok {
		_spec.AddField(userdata.FieldLevelExp, field.TypeInt, value)
	}
	if uduo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userdata.UserTable,
			Columns: []string{userdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uduo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userdata.UserTable,
			Columns: []string{userdata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserData{config: uduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uduo.mutation.done = true
	return _node, nil
}
