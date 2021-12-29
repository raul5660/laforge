// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/finding"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/script"
	"github.com/gen0cide/laforge/ent/user"
	"github.com/google/uuid"
)

// ScriptUpdate is the builder for updating Script entities.
type ScriptUpdate struct {
	config
	hooks    []Hook
	mutation *ScriptMutation
}

// Where appends a list predicates to the ScriptUpdate builder.
func (su *ScriptUpdate) Where(ps ...predicate.Script) *ScriptUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetHclID sets the "hcl_id" field.
func (su *ScriptUpdate) SetHclID(s string) *ScriptUpdate {
	su.mutation.SetHclID(s)
	return su
}

// SetName sets the "name" field.
func (su *ScriptUpdate) SetName(s string) *ScriptUpdate {
	su.mutation.SetName(s)
	return su
}

// SetLanguage sets the "language" field.
func (su *ScriptUpdate) SetLanguage(s string) *ScriptUpdate {
	su.mutation.SetLanguage(s)
	return su
}

// SetDescription sets the "description" field.
func (su *ScriptUpdate) SetDescription(s string) *ScriptUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetSource sets the "source" field.
func (su *ScriptUpdate) SetSource(s string) *ScriptUpdate {
	su.mutation.SetSource(s)
	return su
}

// SetSourceType sets the "source_type" field.
func (su *ScriptUpdate) SetSourceType(s string) *ScriptUpdate {
	su.mutation.SetSourceType(s)
	return su
}

// SetCooldown sets the "cooldown" field.
func (su *ScriptUpdate) SetCooldown(i int) *ScriptUpdate {
	su.mutation.ResetCooldown()
	su.mutation.SetCooldown(i)
	return su
}

// AddCooldown adds i to the "cooldown" field.
func (su *ScriptUpdate) AddCooldown(i int) *ScriptUpdate {
	su.mutation.AddCooldown(i)
	return su
}

// SetTimeout sets the "timeout" field.
func (su *ScriptUpdate) SetTimeout(i int) *ScriptUpdate {
	su.mutation.ResetTimeout()
	su.mutation.SetTimeout(i)
	return su
}

// AddTimeout adds i to the "timeout" field.
func (su *ScriptUpdate) AddTimeout(i int) *ScriptUpdate {
	su.mutation.AddTimeout(i)
	return su
}

// SetIgnoreErrors sets the "ignore_errors" field.
func (su *ScriptUpdate) SetIgnoreErrors(b bool) *ScriptUpdate {
	su.mutation.SetIgnoreErrors(b)
	return su
}

// SetArgs sets the "args" field.
func (su *ScriptUpdate) SetArgs(s []string) *ScriptUpdate {
	su.mutation.SetArgs(s)
	return su
}

// SetDisabled sets the "disabled" field.
func (su *ScriptUpdate) SetDisabled(b bool) *ScriptUpdate {
	su.mutation.SetDisabled(b)
	return su
}

// SetVars sets the "vars" field.
func (su *ScriptUpdate) SetVars(m map[string]string) *ScriptUpdate {
	su.mutation.SetVars(m)
	return su
}

// SetAbsPath sets the "abs_path" field.
func (su *ScriptUpdate) SetAbsPath(s string) *ScriptUpdate {
	su.mutation.SetAbsPath(s)
	return su
}

// SetTags sets the "tags" field.
func (su *ScriptUpdate) SetTags(m map[string]string) *ScriptUpdate {
	su.mutation.SetTags(m)
	return su
}

// AddScriptToUserIDs adds the "ScriptToUser" edge to the User entity by IDs.
func (su *ScriptUpdate) AddScriptToUserIDs(ids ...uuid.UUID) *ScriptUpdate {
	su.mutation.AddScriptToUserIDs(ids...)
	return su
}

// AddScriptToUser adds the "ScriptToUser" edges to the User entity.
func (su *ScriptUpdate) AddScriptToUser(u ...*User) *ScriptUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddScriptToUserIDs(ids...)
}

// AddScriptToFindingIDs adds the "ScriptToFinding" edge to the Finding entity by IDs.
func (su *ScriptUpdate) AddScriptToFindingIDs(ids ...uuid.UUID) *ScriptUpdate {
	su.mutation.AddScriptToFindingIDs(ids...)
	return su
}

// AddScriptToFinding adds the "ScriptToFinding" edges to the Finding entity.
func (su *ScriptUpdate) AddScriptToFinding(f ...*Finding) *ScriptUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.AddScriptToFindingIDs(ids...)
}

// SetScriptToEnvironmentID sets the "ScriptToEnvironment" edge to the Environment entity by ID.
func (su *ScriptUpdate) SetScriptToEnvironmentID(id uuid.UUID) *ScriptUpdate {
	su.mutation.SetScriptToEnvironmentID(id)
	return su
}

// SetNillableScriptToEnvironmentID sets the "ScriptToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (su *ScriptUpdate) SetNillableScriptToEnvironmentID(id *uuid.UUID) *ScriptUpdate {
	if id != nil {
		su = su.SetScriptToEnvironmentID(*id)
	}
	return su
}

// SetScriptToEnvironment sets the "ScriptToEnvironment" edge to the Environment entity.
func (su *ScriptUpdate) SetScriptToEnvironment(e *Environment) *ScriptUpdate {
	return su.SetScriptToEnvironmentID(e.ID)
}

// Mutation returns the ScriptMutation object of the builder.
func (su *ScriptUpdate) Mutation() *ScriptMutation {
	return su.mutation
}

// ClearScriptToUser clears all "ScriptToUser" edges to the User entity.
func (su *ScriptUpdate) ClearScriptToUser() *ScriptUpdate {
	su.mutation.ClearScriptToUser()
	return su
}

// RemoveScriptToUserIDs removes the "ScriptToUser" edge to User entities by IDs.
func (su *ScriptUpdate) RemoveScriptToUserIDs(ids ...uuid.UUID) *ScriptUpdate {
	su.mutation.RemoveScriptToUserIDs(ids...)
	return su
}

// RemoveScriptToUser removes "ScriptToUser" edges to User entities.
func (su *ScriptUpdate) RemoveScriptToUser(u ...*User) *ScriptUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveScriptToUserIDs(ids...)
}

// ClearScriptToFinding clears all "ScriptToFinding" edges to the Finding entity.
func (su *ScriptUpdate) ClearScriptToFinding() *ScriptUpdate {
	su.mutation.ClearScriptToFinding()
	return su
}

// RemoveScriptToFindingIDs removes the "ScriptToFinding" edge to Finding entities by IDs.
func (su *ScriptUpdate) RemoveScriptToFindingIDs(ids ...uuid.UUID) *ScriptUpdate {
	su.mutation.RemoveScriptToFindingIDs(ids...)
	return su
}

// RemoveScriptToFinding removes "ScriptToFinding" edges to Finding entities.
func (su *ScriptUpdate) RemoveScriptToFinding(f ...*Finding) *ScriptUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.RemoveScriptToFindingIDs(ids...)
}

// ClearScriptToEnvironment clears the "ScriptToEnvironment" edge to the Environment entity.
func (su *ScriptUpdate) ClearScriptToEnvironment() *ScriptUpdate {
	su.mutation.ClearScriptToEnvironment()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScriptUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScriptUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScriptUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScriptUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ScriptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   script.Table,
			Columns: script.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: script.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldHclID,
		})
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldName,
		})
	}
	if value, ok := su.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldLanguage,
		})
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldDescription,
		})
	}
	if value, ok := su.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSource,
		})
	}
	if value, ok := su.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSourceType,
		})
	}
	if value, ok := su.mutation.Cooldown(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := su.mutation.AddedCooldown(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := su.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := su.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := su.mutation.IgnoreErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldIgnoreErrors,
		})
	}
	if value, ok := su.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldArgs,
		})
	}
	if value, ok := su.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldDisabled,
		})
	}
	if value, ok := su.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldVars,
		})
	}
	if value, ok := su.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldAbsPath,
		})
	}
	if value, ok := su.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldTags,
		})
	}
	if su.mutation.ScriptToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToUserTable,
			Columns: []string{script.ScriptToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedScriptToUserIDs(); len(nodes) > 0 && !su.mutation.ScriptToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToUserTable,
			Columns: []string{script.ScriptToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ScriptToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToUserTable,
			Columns: []string{script.ScriptToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ScriptToFindingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToFindingTable,
			Columns: []string{script.ScriptToFindingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: finding.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedScriptToFindingIDs(); len(nodes) > 0 && !su.mutation.ScriptToFindingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToFindingTable,
			Columns: []string{script.ScriptToFindingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: finding.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ScriptToFindingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToFindingTable,
			Columns: []string{script.ScriptToFindingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: finding.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ScriptToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   script.ScriptToEnvironmentTable,
			Columns: []string{script.ScriptToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ScriptToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   script.ScriptToEnvironmentTable,
			Columns: []string{script.ScriptToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{script.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ScriptUpdateOne is the builder for updating a single Script entity.
type ScriptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScriptMutation
}

// SetHclID sets the "hcl_id" field.
func (suo *ScriptUpdateOne) SetHclID(s string) *ScriptUpdateOne {
	suo.mutation.SetHclID(s)
	return suo
}

// SetName sets the "name" field.
func (suo *ScriptUpdateOne) SetName(s string) *ScriptUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetLanguage sets the "language" field.
func (suo *ScriptUpdateOne) SetLanguage(s string) *ScriptUpdateOne {
	suo.mutation.SetLanguage(s)
	return suo
}

// SetDescription sets the "description" field.
func (suo *ScriptUpdateOne) SetDescription(s string) *ScriptUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetSource sets the "source" field.
func (suo *ScriptUpdateOne) SetSource(s string) *ScriptUpdateOne {
	suo.mutation.SetSource(s)
	return suo
}

// SetSourceType sets the "source_type" field.
func (suo *ScriptUpdateOne) SetSourceType(s string) *ScriptUpdateOne {
	suo.mutation.SetSourceType(s)
	return suo
}

// SetCooldown sets the "cooldown" field.
func (suo *ScriptUpdateOne) SetCooldown(i int) *ScriptUpdateOne {
	suo.mutation.ResetCooldown()
	suo.mutation.SetCooldown(i)
	return suo
}

// AddCooldown adds i to the "cooldown" field.
func (suo *ScriptUpdateOne) AddCooldown(i int) *ScriptUpdateOne {
	suo.mutation.AddCooldown(i)
	return suo
}

// SetTimeout sets the "timeout" field.
func (suo *ScriptUpdateOne) SetTimeout(i int) *ScriptUpdateOne {
	suo.mutation.ResetTimeout()
	suo.mutation.SetTimeout(i)
	return suo
}

// AddTimeout adds i to the "timeout" field.
func (suo *ScriptUpdateOne) AddTimeout(i int) *ScriptUpdateOne {
	suo.mutation.AddTimeout(i)
	return suo
}

// SetIgnoreErrors sets the "ignore_errors" field.
func (suo *ScriptUpdateOne) SetIgnoreErrors(b bool) *ScriptUpdateOne {
	suo.mutation.SetIgnoreErrors(b)
	return suo
}

// SetArgs sets the "args" field.
func (suo *ScriptUpdateOne) SetArgs(s []string) *ScriptUpdateOne {
	suo.mutation.SetArgs(s)
	return suo
}

// SetDisabled sets the "disabled" field.
func (suo *ScriptUpdateOne) SetDisabled(b bool) *ScriptUpdateOne {
	suo.mutation.SetDisabled(b)
	return suo
}

// SetVars sets the "vars" field.
func (suo *ScriptUpdateOne) SetVars(m map[string]string) *ScriptUpdateOne {
	suo.mutation.SetVars(m)
	return suo
}

// SetAbsPath sets the "abs_path" field.
func (suo *ScriptUpdateOne) SetAbsPath(s string) *ScriptUpdateOne {
	suo.mutation.SetAbsPath(s)
	return suo
}

// SetTags sets the "tags" field.
func (suo *ScriptUpdateOne) SetTags(m map[string]string) *ScriptUpdateOne {
	suo.mutation.SetTags(m)
	return suo
}

// AddScriptToUserIDs adds the "ScriptToUser" edge to the User entity by IDs.
func (suo *ScriptUpdateOne) AddScriptToUserIDs(ids ...uuid.UUID) *ScriptUpdateOne {
	suo.mutation.AddScriptToUserIDs(ids...)
	return suo
}

// AddScriptToUser adds the "ScriptToUser" edges to the User entity.
func (suo *ScriptUpdateOne) AddScriptToUser(u ...*User) *ScriptUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddScriptToUserIDs(ids...)
}

// AddScriptToFindingIDs adds the "ScriptToFinding" edge to the Finding entity by IDs.
func (suo *ScriptUpdateOne) AddScriptToFindingIDs(ids ...uuid.UUID) *ScriptUpdateOne {
	suo.mutation.AddScriptToFindingIDs(ids...)
	return suo
}

// AddScriptToFinding adds the "ScriptToFinding" edges to the Finding entity.
func (suo *ScriptUpdateOne) AddScriptToFinding(f ...*Finding) *ScriptUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.AddScriptToFindingIDs(ids...)
}

// SetScriptToEnvironmentID sets the "ScriptToEnvironment" edge to the Environment entity by ID.
func (suo *ScriptUpdateOne) SetScriptToEnvironmentID(id uuid.UUID) *ScriptUpdateOne {
	suo.mutation.SetScriptToEnvironmentID(id)
	return suo
}

// SetNillableScriptToEnvironmentID sets the "ScriptToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (suo *ScriptUpdateOne) SetNillableScriptToEnvironmentID(id *uuid.UUID) *ScriptUpdateOne {
	if id != nil {
		suo = suo.SetScriptToEnvironmentID(*id)
	}
	return suo
}

// SetScriptToEnvironment sets the "ScriptToEnvironment" edge to the Environment entity.
func (suo *ScriptUpdateOne) SetScriptToEnvironment(e *Environment) *ScriptUpdateOne {
	return suo.SetScriptToEnvironmentID(e.ID)
}

// Mutation returns the ScriptMutation object of the builder.
func (suo *ScriptUpdateOne) Mutation() *ScriptMutation {
	return suo.mutation
}

// ClearScriptToUser clears all "ScriptToUser" edges to the User entity.
func (suo *ScriptUpdateOne) ClearScriptToUser() *ScriptUpdateOne {
	suo.mutation.ClearScriptToUser()
	return suo
}

// RemoveScriptToUserIDs removes the "ScriptToUser" edge to User entities by IDs.
func (suo *ScriptUpdateOne) RemoveScriptToUserIDs(ids ...uuid.UUID) *ScriptUpdateOne {
	suo.mutation.RemoveScriptToUserIDs(ids...)
	return suo
}

// RemoveScriptToUser removes "ScriptToUser" edges to User entities.
func (suo *ScriptUpdateOne) RemoveScriptToUser(u ...*User) *ScriptUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveScriptToUserIDs(ids...)
}

// ClearScriptToFinding clears all "ScriptToFinding" edges to the Finding entity.
func (suo *ScriptUpdateOne) ClearScriptToFinding() *ScriptUpdateOne {
	suo.mutation.ClearScriptToFinding()
	return suo
}

// RemoveScriptToFindingIDs removes the "ScriptToFinding" edge to Finding entities by IDs.
func (suo *ScriptUpdateOne) RemoveScriptToFindingIDs(ids ...uuid.UUID) *ScriptUpdateOne {
	suo.mutation.RemoveScriptToFindingIDs(ids...)
	return suo
}

// RemoveScriptToFinding removes "ScriptToFinding" edges to Finding entities.
func (suo *ScriptUpdateOne) RemoveScriptToFinding(f ...*Finding) *ScriptUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.RemoveScriptToFindingIDs(ids...)
}

// ClearScriptToEnvironment clears the "ScriptToEnvironment" edge to the Environment entity.
func (suo *ScriptUpdateOne) ClearScriptToEnvironment() *ScriptUpdateOne {
	suo.mutation.ClearScriptToEnvironment()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScriptUpdateOne) Select(field string, fields ...string) *ScriptUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Script entity.
func (suo *ScriptUpdateOne) Save(ctx context.Context) (*Script, error) {
	var (
		err  error
		node *Script
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScriptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScriptUpdateOne) SaveX(ctx context.Context) *Script {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScriptUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScriptUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ScriptUpdateOne) sqlSave(ctx context.Context) (_node *Script, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   script.Table,
			Columns: script.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: script.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Script.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, script.FieldID)
		for _, f := range fields {
			if !script.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != script.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldHclID,
		})
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldName,
		})
	}
	if value, ok := suo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldLanguage,
		})
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldDescription,
		})
	}
	if value, ok := suo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSource,
		})
	}
	if value, ok := suo.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldSourceType,
		})
	}
	if value, ok := suo.mutation.Cooldown(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := suo.mutation.AddedCooldown(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldCooldown,
		})
	}
	if value, ok := suo.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := suo.mutation.AddedTimeout(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: script.FieldTimeout,
		})
	}
	if value, ok := suo.mutation.IgnoreErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldIgnoreErrors,
		})
	}
	if value, ok := suo.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldArgs,
		})
	}
	if value, ok := suo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: script.FieldDisabled,
		})
	}
	if value, ok := suo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldVars,
		})
	}
	if value, ok := suo.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: script.FieldAbsPath,
		})
	}
	if value, ok := suo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: script.FieldTags,
		})
	}
	if suo.mutation.ScriptToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToUserTable,
			Columns: []string{script.ScriptToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedScriptToUserIDs(); len(nodes) > 0 && !suo.mutation.ScriptToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToUserTable,
			Columns: []string{script.ScriptToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ScriptToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToUserTable,
			Columns: []string{script.ScriptToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ScriptToFindingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToFindingTable,
			Columns: []string{script.ScriptToFindingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: finding.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedScriptToFindingIDs(); len(nodes) > 0 && !suo.mutation.ScriptToFindingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToFindingTable,
			Columns: []string{script.ScriptToFindingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: finding.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ScriptToFindingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   script.ScriptToFindingTable,
			Columns: []string{script.ScriptToFindingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: finding.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ScriptToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   script.ScriptToEnvironmentTable,
			Columns: []string{script.ScriptToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ScriptToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   script.ScriptToEnvironmentTable,
			Columns: []string{script.ScriptToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Script{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{script.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}