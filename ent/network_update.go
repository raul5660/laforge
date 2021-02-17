// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// NetworkUpdate is the builder for updating Network entities.
type NetworkUpdate struct {
	config
	hooks    []Hook
	mutation *NetworkMutation
}

// Where adds a new predicate for the builder.
func (nu *NetworkUpdate) Where(ps ...predicate.Network) *NetworkUpdate {
	nu.mutation.predicates = append(nu.mutation.predicates, ps...)
	return nu
}

// SetName sets the name field.
func (nu *NetworkUpdate) SetName(s string) *NetworkUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetCidr sets the cidr field.
func (nu *NetworkUpdate) SetCidr(s string) *NetworkUpdate {
	nu.mutation.SetCidr(s)
	return nu
}

// SetVdiVisible sets the vdi_visible field.
func (nu *NetworkUpdate) SetVdiVisible(b bool) *NetworkUpdate {
	nu.mutation.SetVdiVisible(b)
	return nu
}

// SetVars sets the vars field.
func (nu *NetworkUpdate) SetVars(m map[string]string) *NetworkUpdate {
	nu.mutation.SetVars(m)
	return nu
}

// SetTags sets the tags field.
func (nu *NetworkUpdate) SetTags(m map[string]string) *NetworkUpdate {
	nu.mutation.SetTags(m)
	return nu
}

// AddNetworkToTagIDs adds the NetworkToTag edge to Tag by ids.
func (nu *NetworkUpdate) AddNetworkToTagIDs(ids ...int) *NetworkUpdate {
	nu.mutation.AddNetworkToTagIDs(ids...)
	return nu
}

// AddNetworkToTag adds the NetworkToTag edges to Tag.
func (nu *NetworkUpdate) AddNetworkToTag(t ...*Tag) *NetworkUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nu.AddNetworkToTagIDs(ids...)
}

// AddNetworkToEnvironmentIDs adds the NetworkToEnvironment edge to Environment by ids.
func (nu *NetworkUpdate) AddNetworkToEnvironmentIDs(ids ...int) *NetworkUpdate {
	nu.mutation.AddNetworkToEnvironmentIDs(ids...)
	return nu
}

// AddNetworkToEnvironment adds the NetworkToEnvironment edges to Environment.
func (nu *NetworkUpdate) AddNetworkToEnvironment(e ...*Environment) *NetworkUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nu.AddNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the NetworkMutation object of the builder.
func (nu *NetworkUpdate) Mutation() *NetworkMutation {
	return nu.mutation
}

// ClearNetworkToTag clears all "NetworkToTag" edges to type Tag.
func (nu *NetworkUpdate) ClearNetworkToTag() *NetworkUpdate {
	nu.mutation.ClearNetworkToTag()
	return nu
}

// RemoveNetworkToTagIDs removes the NetworkToTag edge to Tag by ids.
func (nu *NetworkUpdate) RemoveNetworkToTagIDs(ids ...int) *NetworkUpdate {
	nu.mutation.RemoveNetworkToTagIDs(ids...)
	return nu
}

// RemoveNetworkToTag removes NetworkToTag edges to Tag.
func (nu *NetworkUpdate) RemoveNetworkToTag(t ...*Tag) *NetworkUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nu.RemoveNetworkToTagIDs(ids...)
}

// ClearNetworkToEnvironment clears all "NetworkToEnvironment" edges to type Environment.
func (nu *NetworkUpdate) ClearNetworkToEnvironment() *NetworkUpdate {
	nu.mutation.ClearNetworkToEnvironment()
	return nu
}

// RemoveNetworkToEnvironmentIDs removes the NetworkToEnvironment edge to Environment by ids.
func (nu *NetworkUpdate) RemoveNetworkToEnvironmentIDs(ids ...int) *NetworkUpdate {
	nu.mutation.RemoveNetworkToEnvironmentIDs(ids...)
	return nu
}

// RemoveNetworkToEnvironment removes NetworkToEnvironment edges to Environment.
func (nu *NetworkUpdate) RemoveNetworkToEnvironment(e ...*Environment) *NetworkUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nu.RemoveNetworkToEnvironmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NetworkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NetworkUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NetworkUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NetworkUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NetworkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   network.Table,
			Columns: network.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: network.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldName,
		})
	}
	if value, ok := nu.mutation.Cidr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldCidr,
		})
	}
	if value, ok := nu.mutation.VdiVisible(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: network.FieldVdiVisible,
		})
	}
	if value, ok := nu.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: network.FieldVars,
		})
	}
	if value, ok := nu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: network.FieldTags,
		})
	}
	if nu.mutation.NetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   network.NetworkToTagTable,
			Columns: []string{network.NetworkToTagColumn},
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
	if nodes := nu.mutation.RemovedNetworkToTagIDs(); len(nodes) > 0 && !nu.mutation.NetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   network.NetworkToTagTable,
			Columns: []string{network.NetworkToTagColumn},
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
	if nodes := nu.mutation.NetworkToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   network.NetworkToTagTable,
			Columns: []string{network.NetworkToTagColumn},
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
	if nu.mutation.NetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   network.NetworkToEnvironmentTable,
			Columns: network.NetworkToEnvironmentPrimaryKey,
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
	if nodes := nu.mutation.RemovedNetworkToEnvironmentIDs(); len(nodes) > 0 && !nu.mutation.NetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   network.NetworkToEnvironmentTable,
			Columns: network.NetworkToEnvironmentPrimaryKey,
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
	if nodes := nu.mutation.NetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   network.NetworkToEnvironmentTable,
			Columns: network.NetworkToEnvironmentPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{network.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// NetworkUpdateOne is the builder for updating a single Network entity.
type NetworkUpdateOne struct {
	config
	hooks    []Hook
	mutation *NetworkMutation
}

// SetName sets the name field.
func (nuo *NetworkUpdateOne) SetName(s string) *NetworkUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetCidr sets the cidr field.
func (nuo *NetworkUpdateOne) SetCidr(s string) *NetworkUpdateOne {
	nuo.mutation.SetCidr(s)
	return nuo
}

// SetVdiVisible sets the vdi_visible field.
func (nuo *NetworkUpdateOne) SetVdiVisible(b bool) *NetworkUpdateOne {
	nuo.mutation.SetVdiVisible(b)
	return nuo
}

// SetVars sets the vars field.
func (nuo *NetworkUpdateOne) SetVars(m map[string]string) *NetworkUpdateOne {
	nuo.mutation.SetVars(m)
	return nuo
}

// SetTags sets the tags field.
func (nuo *NetworkUpdateOne) SetTags(m map[string]string) *NetworkUpdateOne {
	nuo.mutation.SetTags(m)
	return nuo
}

// AddNetworkToTagIDs adds the NetworkToTag edge to Tag by ids.
func (nuo *NetworkUpdateOne) AddNetworkToTagIDs(ids ...int) *NetworkUpdateOne {
	nuo.mutation.AddNetworkToTagIDs(ids...)
	return nuo
}

// AddNetworkToTag adds the NetworkToTag edges to Tag.
func (nuo *NetworkUpdateOne) AddNetworkToTag(t ...*Tag) *NetworkUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nuo.AddNetworkToTagIDs(ids...)
}

// AddNetworkToEnvironmentIDs adds the NetworkToEnvironment edge to Environment by ids.
func (nuo *NetworkUpdateOne) AddNetworkToEnvironmentIDs(ids ...int) *NetworkUpdateOne {
	nuo.mutation.AddNetworkToEnvironmentIDs(ids...)
	return nuo
}

// AddNetworkToEnvironment adds the NetworkToEnvironment edges to Environment.
func (nuo *NetworkUpdateOne) AddNetworkToEnvironment(e ...*Environment) *NetworkUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nuo.AddNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the NetworkMutation object of the builder.
func (nuo *NetworkUpdateOne) Mutation() *NetworkMutation {
	return nuo.mutation
}

// ClearNetworkToTag clears all "NetworkToTag" edges to type Tag.
func (nuo *NetworkUpdateOne) ClearNetworkToTag() *NetworkUpdateOne {
	nuo.mutation.ClearNetworkToTag()
	return nuo
}

// RemoveNetworkToTagIDs removes the NetworkToTag edge to Tag by ids.
func (nuo *NetworkUpdateOne) RemoveNetworkToTagIDs(ids ...int) *NetworkUpdateOne {
	nuo.mutation.RemoveNetworkToTagIDs(ids...)
	return nuo
}

// RemoveNetworkToTag removes NetworkToTag edges to Tag.
func (nuo *NetworkUpdateOne) RemoveNetworkToTag(t ...*Tag) *NetworkUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return nuo.RemoveNetworkToTagIDs(ids...)
}

// ClearNetworkToEnvironment clears all "NetworkToEnvironment" edges to type Environment.
func (nuo *NetworkUpdateOne) ClearNetworkToEnvironment() *NetworkUpdateOne {
	nuo.mutation.ClearNetworkToEnvironment()
	return nuo
}

// RemoveNetworkToEnvironmentIDs removes the NetworkToEnvironment edge to Environment by ids.
func (nuo *NetworkUpdateOne) RemoveNetworkToEnvironmentIDs(ids ...int) *NetworkUpdateOne {
	nuo.mutation.RemoveNetworkToEnvironmentIDs(ids...)
	return nuo
}

// RemoveNetworkToEnvironment removes NetworkToEnvironment edges to Environment.
func (nuo *NetworkUpdateOne) RemoveNetworkToEnvironment(e ...*Environment) *NetworkUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nuo.RemoveNetworkToEnvironmentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (nuo *NetworkUpdateOne) Save(ctx context.Context) (*Network, error) {
	var (
		err  error
		node *Network
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NetworkUpdateOne) SaveX(ctx context.Context) *Network {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NetworkUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NetworkUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NetworkUpdateOne) sqlSave(ctx context.Context) (_node *Network, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   network.Table,
			Columns: network.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: network.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Network.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := nuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldName,
		})
	}
	if value, ok := nuo.mutation.Cidr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldCidr,
		})
	}
	if value, ok := nuo.mutation.VdiVisible(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: network.FieldVdiVisible,
		})
	}
	if value, ok := nuo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: network.FieldVars,
		})
	}
	if value, ok := nuo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: network.FieldTags,
		})
	}
	if nuo.mutation.NetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   network.NetworkToTagTable,
			Columns: []string{network.NetworkToTagColumn},
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
	if nodes := nuo.mutation.RemovedNetworkToTagIDs(); len(nodes) > 0 && !nuo.mutation.NetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   network.NetworkToTagTable,
			Columns: []string{network.NetworkToTagColumn},
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
	if nodes := nuo.mutation.NetworkToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   network.NetworkToTagTable,
			Columns: []string{network.NetworkToTagColumn},
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
	if nuo.mutation.NetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   network.NetworkToEnvironmentTable,
			Columns: network.NetworkToEnvironmentPrimaryKey,
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
	if nodes := nuo.mutation.RemovedNetworkToEnvironmentIDs(); len(nodes) > 0 && !nuo.mutation.NetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   network.NetworkToEnvironmentTable,
			Columns: network.NetworkToEnvironmentPrimaryKey,
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
	if nodes := nuo.mutation.NetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   network.NetworkToEnvironmentTable,
			Columns: network.NetworkToEnvironmentPrimaryKey,
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
	_node = &Network{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{network.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
