// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/hostdependency"
	"github.com/gen0cide/laforge/ent/includednetwork"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/google/uuid"
)

// NetworkCreate is the builder for creating a Network entity.
type NetworkCreate struct {
	config
	mutation *NetworkMutation
	hooks    []Hook
}

// SetHclID sets the "hcl_id" field.
func (nc *NetworkCreate) SetHclID(s string) *NetworkCreate {
	nc.mutation.SetHclID(s)
	return nc
}

// SetName sets the "name" field.
func (nc *NetworkCreate) SetName(s string) *NetworkCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetCidr sets the "cidr" field.
func (nc *NetworkCreate) SetCidr(s string) *NetworkCreate {
	nc.mutation.SetCidr(s)
	return nc
}

// SetVdiVisible sets the "vdi_visible" field.
func (nc *NetworkCreate) SetVdiVisible(b bool) *NetworkCreate {
	nc.mutation.SetVdiVisible(b)
	return nc
}

// SetVars sets the "vars" field.
func (nc *NetworkCreate) SetVars(m map[string]string) *NetworkCreate {
	nc.mutation.SetVars(m)
	return nc
}

// SetTags sets the "tags" field.
func (nc *NetworkCreate) SetTags(m map[string]string) *NetworkCreate {
	nc.mutation.SetTags(m)
	return nc
}

// SetID sets the "id" field.
func (nc *NetworkCreate) SetID(u uuid.UUID) *NetworkCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNetworkToEnvironmentID sets the "NetworkToEnvironment" edge to the Environment entity by ID.
func (nc *NetworkCreate) SetNetworkToEnvironmentID(id uuid.UUID) *NetworkCreate {
	nc.mutation.SetNetworkToEnvironmentID(id)
	return nc
}

// SetNillableNetworkToEnvironmentID sets the "NetworkToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (nc *NetworkCreate) SetNillableNetworkToEnvironmentID(id *uuid.UUID) *NetworkCreate {
	if id != nil {
		nc = nc.SetNetworkToEnvironmentID(*id)
	}
	return nc
}

// SetNetworkToEnvironment sets the "NetworkToEnvironment" edge to the Environment entity.
func (nc *NetworkCreate) SetNetworkToEnvironment(e *Environment) *NetworkCreate {
	return nc.SetNetworkToEnvironmentID(e.ID)
}

// AddNetworkToHostDependencyIDs adds the "NetworkToHostDependency" edge to the HostDependency entity by IDs.
func (nc *NetworkCreate) AddNetworkToHostDependencyIDs(ids ...uuid.UUID) *NetworkCreate {
	nc.mutation.AddNetworkToHostDependencyIDs(ids...)
	return nc
}

// AddNetworkToHostDependency adds the "NetworkToHostDependency" edges to the HostDependency entity.
func (nc *NetworkCreate) AddNetworkToHostDependency(h ...*HostDependency) *NetworkCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return nc.AddNetworkToHostDependencyIDs(ids...)
}

// AddNetworkToIncludedNetworkIDs adds the "NetworkToIncludedNetwork" edge to the IncludedNetwork entity by IDs.
func (nc *NetworkCreate) AddNetworkToIncludedNetworkIDs(ids ...uuid.UUID) *NetworkCreate {
	nc.mutation.AddNetworkToIncludedNetworkIDs(ids...)
	return nc
}

// AddNetworkToIncludedNetwork adds the "NetworkToIncludedNetwork" edges to the IncludedNetwork entity.
func (nc *NetworkCreate) AddNetworkToIncludedNetwork(i ...*IncludedNetwork) *NetworkCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return nc.AddNetworkToIncludedNetworkIDs(ids...)
}

// Mutation returns the NetworkMutation object of the builder.
func (nc *NetworkCreate) Mutation() *NetworkMutation {
	return nc.mutation
}

// Save creates the Network in the database.
func (nc *NetworkCreate) Save(ctx context.Context) (*Network, error) {
	var (
		err  error
		node *Network
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NetworkCreate) SaveX(ctx context.Context) *Network {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NetworkCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NetworkCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NetworkCreate) defaults() {
	if _, ok := nc.mutation.ID(); !ok {
		v := network.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NetworkCreate) check() error {
	if _, ok := nc.mutation.HclID(); !ok {
		return &ValidationError{Name: "hcl_id", err: errors.New(`ent: missing required field "hcl_id"`)}
	}
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := nc.mutation.Cidr(); !ok {
		return &ValidationError{Name: "cidr", err: errors.New(`ent: missing required field "cidr"`)}
	}
	if _, ok := nc.mutation.VdiVisible(); !ok {
		return &ValidationError{Name: "vdi_visible", err: errors.New(`ent: missing required field "vdi_visible"`)}
	}
	if _, ok := nc.mutation.Vars(); !ok {
		return &ValidationError{Name: "vars", err: errors.New(`ent: missing required field "vars"`)}
	}
	if _, ok := nc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "tags"`)}
	}
	return nil
}

func (nc *NetworkCreate) sqlSave(ctx context.Context) (*Network, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (nc *NetworkCreate) createSpec() (*Network, *sqlgraph.CreateSpec) {
	var (
		_node = &Network{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: network.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: network.FieldID,
			},
		}
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.HclID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldHclID,
		})
		_node.HclID = value
	}
	if value, ok := nc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldName,
		})
		_node.Name = value
	}
	if value, ok := nc.mutation.Cidr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldCidr,
		})
		_node.Cidr = value
	}
	if value, ok := nc.mutation.VdiVisible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: network.FieldVdiVisible,
		})
		_node.VdiVisible = value
	}
	if value, ok := nc.mutation.Vars(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: network.FieldVars,
		})
		_node.Vars = value
	}
	if value, ok := nc.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: network.FieldTags,
		})
		_node.Tags = value
	}
	if nodes := nc.mutation.NetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   network.NetworkToEnvironmentTable,
			Columns: []string{network.NetworkToEnvironmentColumn},
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
		_node.environment_environment_to_network = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NetworkToHostDependencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   network.NetworkToHostDependencyTable,
			Columns: []string{network.NetworkToHostDependencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hostdependency.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NetworkToIncludedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   network.NetworkToIncludedNetworkTable,
			Columns: []string{network.NetworkToIncludedNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: includednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NetworkCreateBulk is the builder for creating many Network entities in bulk.
type NetworkCreateBulk struct {
	config
	builders []*NetworkCreate
}

// Save creates the Network entities in the database.
func (ncb *NetworkCreateBulk) Save(ctx context.Context) ([]*Network, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Network, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NetworkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NetworkCreateBulk) SaveX(ctx context.Context) []*Network {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NetworkCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NetworkCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}