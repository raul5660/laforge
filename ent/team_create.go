// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/gen0cide/laforge/ent/user"
)

// TeamCreate is the builder for creating a Team entity.
type TeamCreate struct {
	config
	mutation *TeamMutation
	hooks    []Hook
}

// SetTeamNumber sets the "team_number" field.
func (tc *TeamCreate) SetTeamNumber(i int) *TeamCreate {
	tc.mutation.SetTeamNumber(i)
	return tc
}

// SetConfig sets the "config" field.
func (tc *TeamCreate) SetConfig(m map[string]string) *TeamCreate {
	tc.mutation.SetConfig(m)
	return tc
}

// SetRevision sets the "revision" field.
func (tc *TeamCreate) SetRevision(i int64) *TeamCreate {
	tc.mutation.SetRevision(i)
	return tc
}

// AddTeamToUserIDs adds the "TeamToUser" edge to the User entity by IDs.
func (tc *TeamCreate) AddTeamToUserIDs(ids ...int) *TeamCreate {
	tc.mutation.AddTeamToUserIDs(ids...)
	return tc
}

// AddTeamToUser adds the "TeamToUser" edges to the User entity.
func (tc *TeamCreate) AddTeamToUser(u ...*User) *TeamCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddTeamToUserIDs(ids...)
}

// AddTeamToBuildIDs adds the "TeamToBuild" edge to the Build entity by IDs.
func (tc *TeamCreate) AddTeamToBuildIDs(ids ...int) *TeamCreate {
	tc.mutation.AddTeamToBuildIDs(ids...)
	return tc
}

// AddTeamToBuild adds the "TeamToBuild" edges to the Build entity.
func (tc *TeamCreate) AddTeamToBuild(b ...*Build) *TeamCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tc.AddTeamToBuildIDs(ids...)
}

// AddTeamToEnvironmentIDs adds the "TeamToEnvironment" edge to the Environment entity by IDs.
func (tc *TeamCreate) AddTeamToEnvironmentIDs(ids ...int) *TeamCreate {
	tc.mutation.AddTeamToEnvironmentIDs(ids...)
	return tc
}

// AddTeamToEnvironment adds the "TeamToEnvironment" edges to the Environment entity.
func (tc *TeamCreate) AddTeamToEnvironment(e ...*Environment) *TeamCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tc.AddTeamToEnvironmentIDs(ids...)
}

// AddTeamToTagIDs adds the "TeamToTag" edge to the Tag entity by IDs.
func (tc *TeamCreate) AddTeamToTagIDs(ids ...int) *TeamCreate {
	tc.mutation.AddTeamToTagIDs(ids...)
	return tc
}

// AddTeamToTag adds the "TeamToTag" edges to the Tag entity.
func (tc *TeamCreate) AddTeamToTag(t ...*Tag) *TeamCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTeamToTagIDs(ids...)
}

// AddTeamToProvisionedNetworkIDs adds the "TeamToProvisionedNetwork" edge to the ProvisionedNetwork entity by IDs.
func (tc *TeamCreate) AddTeamToProvisionedNetworkIDs(ids ...int) *TeamCreate {
	tc.mutation.AddTeamToProvisionedNetworkIDs(ids...)
	return tc
}

// AddTeamToProvisionedNetwork adds the "TeamToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (tc *TeamCreate) AddTeamToProvisionedNetwork(p ...*ProvisionedNetwork) *TeamCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddTeamToProvisionedNetworkIDs(ids...)
}

// AddTeamToPlanIDs adds the "TeamToPlan" edge to the Plan entity by IDs.
func (tc *TeamCreate) AddTeamToPlanIDs(ids ...int) *TeamCreate {
	tc.mutation.AddTeamToPlanIDs(ids...)
	return tc
}

// AddTeamToPlan adds the "TeamToPlan" edges to the Plan entity.
func (tc *TeamCreate) AddTeamToPlan(p ...*Plan) *TeamCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddTeamToPlanIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tc *TeamCreate) Mutation() *TeamMutation {
	return tc.mutation
}

// Save creates the Team in the database.
func (tc *TeamCreate) Save(ctx context.Context) (*Team, error) {
	var (
		err  error
		node *Team
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeamCreate) SaveX(ctx context.Context) *Team {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeamCreate) check() error {
	if _, ok := tc.mutation.TeamNumber(); !ok {
		return &ValidationError{Name: "team_number", err: errors.New("ent: missing required field \"team_number\"")}
	}
	if _, ok := tc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New("ent: missing required field \"config\"")}
	}
	if _, ok := tc.mutation.Revision(); !ok {
		return &ValidationError{Name: "revision", err: errors.New("ent: missing required field \"revision\"")}
	}
	return nil
}

func (tc *TeamCreate) sqlSave(ctx context.Context) (*Team, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TeamCreate) createSpec() (*Team, *sqlgraph.CreateSpec) {
	var (
		_node = &Team{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: team.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: team.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.TeamNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: team.FieldTeamNumber,
		})
		_node.TeamNumber = value
	}
	if value, ok := tc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: team.FieldConfig,
		})
		_node.Config = value
	}
	if value, ok := tc.mutation.Revision(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: team.FieldRevision,
		})
		_node.Revision = value
	}
	if nodes := tc.mutation.TeamToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamToUserTable,
			Columns: []string{team.TeamToUserColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeamToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.TeamToBuildTable,
			Columns: team.TeamToBuildPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeamToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.TeamToEnvironmentTable,
			Columns: team.TeamToEnvironmentPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeamToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamToTagTable,
			Columns: []string{team.TeamToTagColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeamToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   team.TeamToProvisionedNetworkTable,
			Columns: team.TeamToProvisionedNetworkPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TeamToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   team.TeamToPlanTable,
			Columns: team.TeamToPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
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

// TeamCreateBulk is the builder for creating many Team entities in bulk.
type TeamCreateBulk struct {
	config
	builders []*TeamCreate
}

// Save creates the Team entities in the database.
func (tcb *TeamCreateBulk) Save(ctx context.Context) ([]*Team, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Team, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeamMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeamCreateBulk) SaveX(ctx context.Context) []*Team {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
