// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/includednetwork"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// IncludedNetworkUpdate is the builder for updating IncludedNetwork entities.
type IncludedNetworkUpdate struct {
	config
	hooks    []Hook
	mutation *IncludedNetworkMutation
}

// Where adds a new predicate for the IncludedNetworkUpdate builder.
func (inu *IncludedNetworkUpdate) Where(ps ...predicate.IncludedNetwork) *IncludedNetworkUpdate {
	inu.mutation.predicates = append(inu.mutation.predicates, ps...)
	return inu
}

// SetName sets the "name" field.
func (inu *IncludedNetworkUpdate) SetName(s string) *IncludedNetworkUpdate {
	inu.mutation.SetName(s)
	return inu
}

// SetHosts sets the "hosts" field.
func (inu *IncludedNetworkUpdate) SetHosts(s []string) *IncludedNetworkUpdate {
	inu.mutation.SetHosts(s)
	return inu
}

// AddIncludedNetworkToTagIDs adds the "IncludedNetworkToTag" edge to the Tag entity by IDs.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToTagIDs(ids ...int) *IncludedNetworkUpdate {
	inu.mutation.AddIncludedNetworkToTagIDs(ids...)
	return inu
}

// AddIncludedNetworkToTag adds the "IncludedNetworkToTag" edges to the Tag entity.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inu.AddIncludedNetworkToTagIDs(ids...)
}

// AddIncludedNetworkToEnvironmentIDs adds the "IncludedNetworkToEnvironment" edge to the Environment entity by IDs.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToEnvironmentIDs(ids ...int) *IncludedNetworkUpdate {
	inu.mutation.AddIncludedNetworkToEnvironmentIDs(ids...)
	return inu
}

// AddIncludedNetworkToEnvironment adds the "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inu.AddIncludedNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the IncludedNetworkMutation object of the builder.
func (inu *IncludedNetworkUpdate) Mutation() *IncludedNetworkMutation {
	return inu.mutation
}

// ClearIncludedNetworkToTag clears all "IncludedNetworkToTag" edges to the Tag entity.
func (inu *IncludedNetworkUpdate) ClearIncludedNetworkToTag() *IncludedNetworkUpdate {
	inu.mutation.ClearIncludedNetworkToTag()
	return inu
}

// RemoveIncludedNetworkToTagIDs removes the "IncludedNetworkToTag" edge to Tag entities by IDs.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToTagIDs(ids ...int) *IncludedNetworkUpdate {
	inu.mutation.RemoveIncludedNetworkToTagIDs(ids...)
	return inu
}

// RemoveIncludedNetworkToTag removes "IncludedNetworkToTag" edges to Tag entities.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inu.RemoveIncludedNetworkToTagIDs(ids...)
}

// ClearIncludedNetworkToEnvironment clears all "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inu *IncludedNetworkUpdate) ClearIncludedNetworkToEnvironment() *IncludedNetworkUpdate {
	inu.mutation.ClearIncludedNetworkToEnvironment()
	return inu
}

// RemoveIncludedNetworkToEnvironmentIDs removes the "IncludedNetworkToEnvironment" edge to Environment entities by IDs.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToEnvironmentIDs(ids ...int) *IncludedNetworkUpdate {
	inu.mutation.RemoveIncludedNetworkToEnvironmentIDs(ids...)
	return inu
}

// RemoveIncludedNetworkToEnvironment removes "IncludedNetworkToEnvironment" edges to Environment entities.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inu.RemoveIncludedNetworkToEnvironmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (inu *IncludedNetworkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(inu.hooks) == 0 {
		affected, err = inu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncludedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			inu.mutation = mutation
			affected, err = inu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(inu.hooks) - 1; i >= 0; i-- {
			mut = inu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, inu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (inu *IncludedNetworkUpdate) SaveX(ctx context.Context) int {
	affected, err := inu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (inu *IncludedNetworkUpdate) Exec(ctx context.Context) error {
	_, err := inu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (inu *IncludedNetworkUpdate) ExecX(ctx context.Context) {
	if err := inu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (inu *IncludedNetworkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   includednetwork.Table,
			Columns: includednetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: includednetwork.FieldID,
			},
		},
	}
	if ps := inu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := inu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: includednetwork.FieldName,
		})
	}
	if value, ok := inu.mutation.Hosts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: includednetwork.FieldHosts,
		})
	}
	if inu.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.RemovedIncludedNetworkToTagIDs(); len(nodes) > 0 && !inu.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.IncludedNetworkToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inu.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.RemovedIncludedNetworkToEnvironmentIDs(); len(nodes) > 0 && !inu.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.IncludedNetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, inu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{includednetwork.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// IncludedNetworkUpdateOne is the builder for updating a single IncludedNetwork entity.
type IncludedNetworkUpdateOne struct {
	config
	hooks    []Hook
	mutation *IncludedNetworkMutation
}

// SetName sets the "name" field.
func (inuo *IncludedNetworkUpdateOne) SetName(s string) *IncludedNetworkUpdateOne {
	inuo.mutation.SetName(s)
	return inuo
}

// SetHosts sets the "hosts" field.
func (inuo *IncludedNetworkUpdateOne) SetHosts(s []string) *IncludedNetworkUpdateOne {
	inuo.mutation.SetHosts(s)
	return inuo
}

// AddIncludedNetworkToTagIDs adds the "IncludedNetworkToTag" edge to the Tag entity by IDs.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToTagIDs(ids ...int) *IncludedNetworkUpdateOne {
	inuo.mutation.AddIncludedNetworkToTagIDs(ids...)
	return inuo
}

// AddIncludedNetworkToTag adds the "IncludedNetworkToTag" edges to the Tag entity.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inuo.AddIncludedNetworkToTagIDs(ids...)
}

// AddIncludedNetworkToEnvironmentIDs adds the "IncludedNetworkToEnvironment" edge to the Environment entity by IDs.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToEnvironmentIDs(ids ...int) *IncludedNetworkUpdateOne {
	inuo.mutation.AddIncludedNetworkToEnvironmentIDs(ids...)
	return inuo
}

// AddIncludedNetworkToEnvironment adds the "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inuo.AddIncludedNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the IncludedNetworkMutation object of the builder.
func (inuo *IncludedNetworkUpdateOne) Mutation() *IncludedNetworkMutation {
	return inuo.mutation
}

// ClearIncludedNetworkToTag clears all "IncludedNetworkToTag" edges to the Tag entity.
func (inuo *IncludedNetworkUpdateOne) ClearIncludedNetworkToTag() *IncludedNetworkUpdateOne {
	inuo.mutation.ClearIncludedNetworkToTag()
	return inuo
}

// RemoveIncludedNetworkToTagIDs removes the "IncludedNetworkToTag" edge to Tag entities by IDs.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToTagIDs(ids ...int) *IncludedNetworkUpdateOne {
	inuo.mutation.RemoveIncludedNetworkToTagIDs(ids...)
	return inuo
}

// RemoveIncludedNetworkToTag removes "IncludedNetworkToTag" edges to Tag entities.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inuo.RemoveIncludedNetworkToTagIDs(ids...)
}

// ClearIncludedNetworkToEnvironment clears all "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inuo *IncludedNetworkUpdateOne) ClearIncludedNetworkToEnvironment() *IncludedNetworkUpdateOne {
	inuo.mutation.ClearIncludedNetworkToEnvironment()
	return inuo
}

// RemoveIncludedNetworkToEnvironmentIDs removes the "IncludedNetworkToEnvironment" edge to Environment entities by IDs.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToEnvironmentIDs(ids ...int) *IncludedNetworkUpdateOne {
	inuo.mutation.RemoveIncludedNetworkToEnvironmentIDs(ids...)
	return inuo
}

// RemoveIncludedNetworkToEnvironment removes "IncludedNetworkToEnvironment" edges to Environment entities.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inuo.RemoveIncludedNetworkToEnvironmentIDs(ids...)
}

// Save executes the query and returns the updated IncludedNetwork entity.
func (inuo *IncludedNetworkUpdateOne) Save(ctx context.Context) (*IncludedNetwork, error) {
	var (
		err  error
		node *IncludedNetwork
	)
	if len(inuo.hooks) == 0 {
		node, err = inuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncludedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			inuo.mutation = mutation
			node, err = inuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(inuo.hooks) - 1; i >= 0; i-- {
			mut = inuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, inuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (inuo *IncludedNetworkUpdateOne) SaveX(ctx context.Context) *IncludedNetwork {
	node, err := inuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (inuo *IncludedNetworkUpdateOne) Exec(ctx context.Context) error {
	_, err := inuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (inuo *IncludedNetworkUpdateOne) ExecX(ctx context.Context) {
	if err := inuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (inuo *IncludedNetworkUpdateOne) sqlSave(ctx context.Context) (_node *IncludedNetwork, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   includednetwork.Table,
			Columns: includednetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: includednetwork.FieldID,
			},
		},
	}
	id, ok := inuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing IncludedNetwork.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := inuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := inuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: includednetwork.FieldName,
		})
	}
	if value, ok := inuo.mutation.Hosts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: includednetwork.FieldHosts,
		})
	}
	if inuo.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.RemovedIncludedNetworkToTagIDs(); len(nodes) > 0 && !inuo.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.IncludedNetworkToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inuo.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.RemovedIncludedNetworkToEnvironmentIDs(); len(nodes) > 0 && !inuo.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.IncludedNetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IncludedNetwork{config: inuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, inuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{includednetwork.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
