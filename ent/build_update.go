// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/gen0cide/laforge/ent/user"
)

// BuildUpdate is the builder for updating Build entities.
type BuildUpdate struct {
	config
	hooks    []Hook
	mutation *BuildMutation
}

// Where adds a new predicate for the BuildUpdate builder.
func (bu *BuildUpdate) Where(ps ...predicate.Build) *BuildUpdate {
	bu.mutation.predicates = append(bu.mutation.predicates, ps...)
	return bu
}

// SetRevision sets the "revision" field.
func (bu *BuildUpdate) SetRevision(i int) *BuildUpdate {
	bu.mutation.ResetRevision()
	bu.mutation.SetRevision(i)
	return bu
}

// AddRevision adds i to the "revision" field.
func (bu *BuildUpdate) AddRevision(i int) *BuildUpdate {
	bu.mutation.AddRevision(i)
	return bu
}

// SetConfig sets the "config" field.
func (bu *BuildUpdate) SetConfig(m map[string]string) *BuildUpdate {
	bu.mutation.SetConfig(m)
	return bu
}

// AddBuildToUserIDs adds the "BuildToUser" edge to the User entity by IDs.
func (bu *BuildUpdate) AddBuildToUserIDs(ids ...int) *BuildUpdate {
	bu.mutation.AddBuildToUserIDs(ids...)
	return bu
}

// AddBuildToUser adds the "BuildToUser" edges to the User entity.
func (bu *BuildUpdate) AddBuildToUser(u ...*User) *BuildUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.AddBuildToUserIDs(ids...)
}

// AddBuildToTagIDs adds the "BuildToTag" edge to the Tag entity by IDs.
func (bu *BuildUpdate) AddBuildToTagIDs(ids ...int) *BuildUpdate {
	bu.mutation.AddBuildToTagIDs(ids...)
	return bu
}

// AddBuildToTag adds the "BuildToTag" edges to the Tag entity.
func (bu *BuildUpdate) AddBuildToTag(t ...*Tag) *BuildUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddBuildToTagIDs(ids...)
}

// AddBuildToProvisionedNetworkIDs adds the "BuildToProvisionedNetwork" edge to the ProvisionedNetwork entity by IDs.
func (bu *BuildUpdate) AddBuildToProvisionedNetworkIDs(ids ...int) *BuildUpdate {
	bu.mutation.AddBuildToProvisionedNetworkIDs(ids...)
	return bu
}

// AddBuildToProvisionedNetwork adds the "BuildToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (bu *BuildUpdate) AddBuildToProvisionedNetwork(p ...*ProvisionedNetwork) *BuildUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddBuildToProvisionedNetworkIDs(ids...)
}

// AddBuildToTeamIDs adds the "BuildToTeam" edge to the Team entity by IDs.
func (bu *BuildUpdate) AddBuildToTeamIDs(ids ...int) *BuildUpdate {
	bu.mutation.AddBuildToTeamIDs(ids...)
	return bu
}

// AddBuildToTeam adds the "BuildToTeam" edges to the Team entity.
func (bu *BuildUpdate) AddBuildToTeam(t ...*Team) *BuildUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddBuildToTeamIDs(ids...)
}

// AddBuildToEnvironmentIDs adds the "BuildToEnvironment" edge to the Environment entity by IDs.
func (bu *BuildUpdate) AddBuildToEnvironmentIDs(ids ...int) *BuildUpdate {
	bu.mutation.AddBuildToEnvironmentIDs(ids...)
	return bu
}

// AddBuildToEnvironment adds the "BuildToEnvironment" edges to the Environment entity.
func (bu *BuildUpdate) AddBuildToEnvironment(e ...*Environment) *BuildUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return bu.AddBuildToEnvironmentIDs(ids...)
}

// AddBuildToPlanIDs adds the "BuildToPlan" edge to the Plan entity by IDs.
func (bu *BuildUpdate) AddBuildToPlanIDs(ids ...int) *BuildUpdate {
	bu.mutation.AddBuildToPlanIDs(ids...)
	return bu
}

// AddBuildToPlan adds the "BuildToPlan" edges to the Plan entity.
func (bu *BuildUpdate) AddBuildToPlan(p ...*Plan) *BuildUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddBuildToPlanIDs(ids...)
}

// Mutation returns the BuildMutation object of the builder.
func (bu *BuildUpdate) Mutation() *BuildMutation {
	return bu.mutation
}

// ClearBuildToUser clears all "BuildToUser" edges to the User entity.
func (bu *BuildUpdate) ClearBuildToUser() *BuildUpdate {
	bu.mutation.ClearBuildToUser()
	return bu
}

// RemoveBuildToUserIDs removes the "BuildToUser" edge to User entities by IDs.
func (bu *BuildUpdate) RemoveBuildToUserIDs(ids ...int) *BuildUpdate {
	bu.mutation.RemoveBuildToUserIDs(ids...)
	return bu
}

// RemoveBuildToUser removes "BuildToUser" edges to User entities.
func (bu *BuildUpdate) RemoveBuildToUser(u ...*User) *BuildUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.RemoveBuildToUserIDs(ids...)
}

// ClearBuildToTag clears all "BuildToTag" edges to the Tag entity.
func (bu *BuildUpdate) ClearBuildToTag() *BuildUpdate {
	bu.mutation.ClearBuildToTag()
	return bu
}

// RemoveBuildToTagIDs removes the "BuildToTag" edge to Tag entities by IDs.
func (bu *BuildUpdate) RemoveBuildToTagIDs(ids ...int) *BuildUpdate {
	bu.mutation.RemoveBuildToTagIDs(ids...)
	return bu
}

// RemoveBuildToTag removes "BuildToTag" edges to Tag entities.
func (bu *BuildUpdate) RemoveBuildToTag(t ...*Tag) *BuildUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveBuildToTagIDs(ids...)
}

// ClearBuildToProvisionedNetwork clears all "BuildToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (bu *BuildUpdate) ClearBuildToProvisionedNetwork() *BuildUpdate {
	bu.mutation.ClearBuildToProvisionedNetwork()
	return bu
}

// RemoveBuildToProvisionedNetworkIDs removes the "BuildToProvisionedNetwork" edge to ProvisionedNetwork entities by IDs.
func (bu *BuildUpdate) RemoveBuildToProvisionedNetworkIDs(ids ...int) *BuildUpdate {
	bu.mutation.RemoveBuildToProvisionedNetworkIDs(ids...)
	return bu
}

// RemoveBuildToProvisionedNetwork removes "BuildToProvisionedNetwork" edges to ProvisionedNetwork entities.
func (bu *BuildUpdate) RemoveBuildToProvisionedNetwork(p ...*ProvisionedNetwork) *BuildUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemoveBuildToProvisionedNetworkIDs(ids...)
}

// ClearBuildToTeam clears all "BuildToTeam" edges to the Team entity.
func (bu *BuildUpdate) ClearBuildToTeam() *BuildUpdate {
	bu.mutation.ClearBuildToTeam()
	return bu
}

// RemoveBuildToTeamIDs removes the "BuildToTeam" edge to Team entities by IDs.
func (bu *BuildUpdate) RemoveBuildToTeamIDs(ids ...int) *BuildUpdate {
	bu.mutation.RemoveBuildToTeamIDs(ids...)
	return bu
}

// RemoveBuildToTeam removes "BuildToTeam" edges to Team entities.
func (bu *BuildUpdate) RemoveBuildToTeam(t ...*Team) *BuildUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveBuildToTeamIDs(ids...)
}

// ClearBuildToEnvironment clears all "BuildToEnvironment" edges to the Environment entity.
func (bu *BuildUpdate) ClearBuildToEnvironment() *BuildUpdate {
	bu.mutation.ClearBuildToEnvironment()
	return bu
}

// RemoveBuildToEnvironmentIDs removes the "BuildToEnvironment" edge to Environment entities by IDs.
func (bu *BuildUpdate) RemoveBuildToEnvironmentIDs(ids ...int) *BuildUpdate {
	bu.mutation.RemoveBuildToEnvironmentIDs(ids...)
	return bu
}

// RemoveBuildToEnvironment removes "BuildToEnvironment" edges to Environment entities.
func (bu *BuildUpdate) RemoveBuildToEnvironment(e ...*Environment) *BuildUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return bu.RemoveBuildToEnvironmentIDs(ids...)
}

// ClearBuildToPlan clears all "BuildToPlan" edges to the Plan entity.
func (bu *BuildUpdate) ClearBuildToPlan() *BuildUpdate {
	bu.mutation.ClearBuildToPlan()
	return bu
}

// RemoveBuildToPlanIDs removes the "BuildToPlan" edge to Plan entities by IDs.
func (bu *BuildUpdate) RemoveBuildToPlanIDs(ids ...int) *BuildUpdate {
	bu.mutation.RemoveBuildToPlanIDs(ids...)
	return bu
}

// RemoveBuildToPlan removes "BuildToPlan" edges to Plan entities.
func (bu *BuildUpdate) RemoveBuildToPlan(p ...*Plan) *BuildUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemoveBuildToPlanIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BuildUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BuildUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BuildUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BuildUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BuildUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   build.Table,
			Columns: build.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: build.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldRevision,
		})
	}
	if value, ok := bu.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldRevision,
		})
	}
	if value, ok := bu.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: build.FieldConfig,
		})
	}
	if bu.mutation.BuildToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToUserTable,
			Columns: []string{build.BuildToUserColumn},
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
	if nodes := bu.mutation.RemovedBuildToUserIDs(); len(nodes) > 0 && !bu.mutation.BuildToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToUserTable,
			Columns: []string{build.BuildToUserColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BuildToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToUserTable,
			Columns: []string{build.BuildToUserColumn},
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
	if bu.mutation.BuildToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToTagTable,
			Columns: []string{build.BuildToTagColumn},
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
	if nodes := bu.mutation.RemovedBuildToTagIDs(); len(nodes) > 0 && !bu.mutation.BuildToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToTagTable,
			Columns: []string{build.BuildToTagColumn},
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
	if nodes := bu.mutation.BuildToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToTagTable,
			Columns: []string{build.BuildToTagColumn},
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
	if bu.mutation.BuildToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   build.BuildToProvisionedNetworkTable,
			Columns: build.BuildToProvisionedNetworkPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedBuildToProvisionedNetworkIDs(); len(nodes) > 0 && !bu.mutation.BuildToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   build.BuildToProvisionedNetworkTable,
			Columns: build.BuildToProvisionedNetworkPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BuildToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   build.BuildToProvisionedNetworkTable,
			Columns: build.BuildToProvisionedNetworkPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.BuildToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToTeamTable,
			Columns: build.BuildToTeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedBuildToTeamIDs(); len(nodes) > 0 && !bu.mutation.BuildToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToTeamTable,
			Columns: build.BuildToTeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BuildToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToTeamTable,
			Columns: build.BuildToTeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.BuildToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToEnvironmentTable,
			Columns: build.BuildToEnvironmentPrimaryKey,
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
	if nodes := bu.mutation.RemovedBuildToEnvironmentIDs(); len(nodes) > 0 && !bu.mutation.BuildToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToEnvironmentTable,
			Columns: build.BuildToEnvironmentPrimaryKey,
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
	if nodes := bu.mutation.BuildToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToEnvironmentTable,
			Columns: build.BuildToEnvironmentPrimaryKey,
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
	if bu.mutation.BuildToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToPlanTable,
			Columns: build.BuildToPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedBuildToPlanIDs(); len(nodes) > 0 && !bu.mutation.BuildToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToPlanTable,
			Columns: build.BuildToPlanPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BuildToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToPlanTable,
			Columns: build.BuildToPlanPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{build.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BuildUpdateOne is the builder for updating a single Build entity.
type BuildUpdateOne struct {
	config
	hooks    []Hook
	mutation *BuildMutation
}

// SetRevision sets the "revision" field.
func (buo *BuildUpdateOne) SetRevision(i int) *BuildUpdateOne {
	buo.mutation.ResetRevision()
	buo.mutation.SetRevision(i)
	return buo
}

// AddRevision adds i to the "revision" field.
func (buo *BuildUpdateOne) AddRevision(i int) *BuildUpdateOne {
	buo.mutation.AddRevision(i)
	return buo
}

// SetConfig sets the "config" field.
func (buo *BuildUpdateOne) SetConfig(m map[string]string) *BuildUpdateOne {
	buo.mutation.SetConfig(m)
	return buo
}

// AddBuildToUserIDs adds the "BuildToUser" edge to the User entity by IDs.
func (buo *BuildUpdateOne) AddBuildToUserIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.AddBuildToUserIDs(ids...)
	return buo
}

// AddBuildToUser adds the "BuildToUser" edges to the User entity.
func (buo *BuildUpdateOne) AddBuildToUser(u ...*User) *BuildUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.AddBuildToUserIDs(ids...)
}

// AddBuildToTagIDs adds the "BuildToTag" edge to the Tag entity by IDs.
func (buo *BuildUpdateOne) AddBuildToTagIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.AddBuildToTagIDs(ids...)
	return buo
}

// AddBuildToTag adds the "BuildToTag" edges to the Tag entity.
func (buo *BuildUpdateOne) AddBuildToTag(t ...*Tag) *BuildUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddBuildToTagIDs(ids...)
}

// AddBuildToProvisionedNetworkIDs adds the "BuildToProvisionedNetwork" edge to the ProvisionedNetwork entity by IDs.
func (buo *BuildUpdateOne) AddBuildToProvisionedNetworkIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.AddBuildToProvisionedNetworkIDs(ids...)
	return buo
}

// AddBuildToProvisionedNetwork adds the "BuildToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (buo *BuildUpdateOne) AddBuildToProvisionedNetwork(p ...*ProvisionedNetwork) *BuildUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddBuildToProvisionedNetworkIDs(ids...)
}

// AddBuildToTeamIDs adds the "BuildToTeam" edge to the Team entity by IDs.
func (buo *BuildUpdateOne) AddBuildToTeamIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.AddBuildToTeamIDs(ids...)
	return buo
}

// AddBuildToTeam adds the "BuildToTeam" edges to the Team entity.
func (buo *BuildUpdateOne) AddBuildToTeam(t ...*Team) *BuildUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddBuildToTeamIDs(ids...)
}

// AddBuildToEnvironmentIDs adds the "BuildToEnvironment" edge to the Environment entity by IDs.
func (buo *BuildUpdateOne) AddBuildToEnvironmentIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.AddBuildToEnvironmentIDs(ids...)
	return buo
}

// AddBuildToEnvironment adds the "BuildToEnvironment" edges to the Environment entity.
func (buo *BuildUpdateOne) AddBuildToEnvironment(e ...*Environment) *BuildUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return buo.AddBuildToEnvironmentIDs(ids...)
}

// AddBuildToPlanIDs adds the "BuildToPlan" edge to the Plan entity by IDs.
func (buo *BuildUpdateOne) AddBuildToPlanIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.AddBuildToPlanIDs(ids...)
	return buo
}

// AddBuildToPlan adds the "BuildToPlan" edges to the Plan entity.
func (buo *BuildUpdateOne) AddBuildToPlan(p ...*Plan) *BuildUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddBuildToPlanIDs(ids...)
}

// Mutation returns the BuildMutation object of the builder.
func (buo *BuildUpdateOne) Mutation() *BuildMutation {
	return buo.mutation
}

// ClearBuildToUser clears all "BuildToUser" edges to the User entity.
func (buo *BuildUpdateOne) ClearBuildToUser() *BuildUpdateOne {
	buo.mutation.ClearBuildToUser()
	return buo
}

// RemoveBuildToUserIDs removes the "BuildToUser" edge to User entities by IDs.
func (buo *BuildUpdateOne) RemoveBuildToUserIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.RemoveBuildToUserIDs(ids...)
	return buo
}

// RemoveBuildToUser removes "BuildToUser" edges to User entities.
func (buo *BuildUpdateOne) RemoveBuildToUser(u ...*User) *BuildUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.RemoveBuildToUserIDs(ids...)
}

// ClearBuildToTag clears all "BuildToTag" edges to the Tag entity.
func (buo *BuildUpdateOne) ClearBuildToTag() *BuildUpdateOne {
	buo.mutation.ClearBuildToTag()
	return buo
}

// RemoveBuildToTagIDs removes the "BuildToTag" edge to Tag entities by IDs.
func (buo *BuildUpdateOne) RemoveBuildToTagIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.RemoveBuildToTagIDs(ids...)
	return buo
}

// RemoveBuildToTag removes "BuildToTag" edges to Tag entities.
func (buo *BuildUpdateOne) RemoveBuildToTag(t ...*Tag) *BuildUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveBuildToTagIDs(ids...)
}

// ClearBuildToProvisionedNetwork clears all "BuildToProvisionedNetwork" edges to the ProvisionedNetwork entity.
func (buo *BuildUpdateOne) ClearBuildToProvisionedNetwork() *BuildUpdateOne {
	buo.mutation.ClearBuildToProvisionedNetwork()
	return buo
}

// RemoveBuildToProvisionedNetworkIDs removes the "BuildToProvisionedNetwork" edge to ProvisionedNetwork entities by IDs.
func (buo *BuildUpdateOne) RemoveBuildToProvisionedNetworkIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.RemoveBuildToProvisionedNetworkIDs(ids...)
	return buo
}

// RemoveBuildToProvisionedNetwork removes "BuildToProvisionedNetwork" edges to ProvisionedNetwork entities.
func (buo *BuildUpdateOne) RemoveBuildToProvisionedNetwork(p ...*ProvisionedNetwork) *BuildUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemoveBuildToProvisionedNetworkIDs(ids...)
}

// ClearBuildToTeam clears all "BuildToTeam" edges to the Team entity.
func (buo *BuildUpdateOne) ClearBuildToTeam() *BuildUpdateOne {
	buo.mutation.ClearBuildToTeam()
	return buo
}

// RemoveBuildToTeamIDs removes the "BuildToTeam" edge to Team entities by IDs.
func (buo *BuildUpdateOne) RemoveBuildToTeamIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.RemoveBuildToTeamIDs(ids...)
	return buo
}

// RemoveBuildToTeam removes "BuildToTeam" edges to Team entities.
func (buo *BuildUpdateOne) RemoveBuildToTeam(t ...*Team) *BuildUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveBuildToTeamIDs(ids...)
}

// ClearBuildToEnvironment clears all "BuildToEnvironment" edges to the Environment entity.
func (buo *BuildUpdateOne) ClearBuildToEnvironment() *BuildUpdateOne {
	buo.mutation.ClearBuildToEnvironment()
	return buo
}

// RemoveBuildToEnvironmentIDs removes the "BuildToEnvironment" edge to Environment entities by IDs.
func (buo *BuildUpdateOne) RemoveBuildToEnvironmentIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.RemoveBuildToEnvironmentIDs(ids...)
	return buo
}

// RemoveBuildToEnvironment removes "BuildToEnvironment" edges to Environment entities.
func (buo *BuildUpdateOne) RemoveBuildToEnvironment(e ...*Environment) *BuildUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return buo.RemoveBuildToEnvironmentIDs(ids...)
}

// ClearBuildToPlan clears all "BuildToPlan" edges to the Plan entity.
func (buo *BuildUpdateOne) ClearBuildToPlan() *BuildUpdateOne {
	buo.mutation.ClearBuildToPlan()
	return buo
}

// RemoveBuildToPlanIDs removes the "BuildToPlan" edge to Plan entities by IDs.
func (buo *BuildUpdateOne) RemoveBuildToPlanIDs(ids ...int) *BuildUpdateOne {
	buo.mutation.RemoveBuildToPlanIDs(ids...)
	return buo
}

// RemoveBuildToPlan removes "BuildToPlan" edges to Plan entities.
func (buo *BuildUpdateOne) RemoveBuildToPlan(p ...*Plan) *BuildUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemoveBuildToPlanIDs(ids...)
}

// Save executes the query and returns the updated Build entity.
func (buo *BuildUpdateOne) Save(ctx context.Context) (*Build, error) {
	var (
		err  error
		node *Build
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BuildUpdateOne) SaveX(ctx context.Context) *Build {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BuildUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BuildUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BuildUpdateOne) sqlSave(ctx context.Context) (_node *Build, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   build.Table,
			Columns: build.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: build.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Build.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldRevision,
		})
	}
	if value, ok := buo.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldRevision,
		})
	}
	if value, ok := buo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: build.FieldConfig,
		})
	}
	if buo.mutation.BuildToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToUserTable,
			Columns: []string{build.BuildToUserColumn},
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
	if nodes := buo.mutation.RemovedBuildToUserIDs(); len(nodes) > 0 && !buo.mutation.BuildToUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToUserTable,
			Columns: []string{build.BuildToUserColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BuildToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToUserTable,
			Columns: []string{build.BuildToUserColumn},
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
	if buo.mutation.BuildToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToTagTable,
			Columns: []string{build.BuildToTagColumn},
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
	if nodes := buo.mutation.RemovedBuildToTagIDs(); len(nodes) > 0 && !buo.mutation.BuildToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToTagTable,
			Columns: []string{build.BuildToTagColumn},
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
	if nodes := buo.mutation.BuildToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   build.BuildToTagTable,
			Columns: []string{build.BuildToTagColumn},
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
	if buo.mutation.BuildToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   build.BuildToProvisionedNetworkTable,
			Columns: build.BuildToProvisionedNetworkPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provisionednetwork.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedBuildToProvisionedNetworkIDs(); len(nodes) > 0 && !buo.mutation.BuildToProvisionedNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   build.BuildToProvisionedNetworkTable,
			Columns: build.BuildToProvisionedNetworkPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BuildToProvisionedNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   build.BuildToProvisionedNetworkTable,
			Columns: build.BuildToProvisionedNetworkPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.BuildToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToTeamTable,
			Columns: build.BuildToTeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedBuildToTeamIDs(); len(nodes) > 0 && !buo.mutation.BuildToTeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToTeamTable,
			Columns: build.BuildToTeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BuildToTeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToTeamTable,
			Columns: build.BuildToTeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.BuildToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToEnvironmentTable,
			Columns: build.BuildToEnvironmentPrimaryKey,
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
	if nodes := buo.mutation.RemovedBuildToEnvironmentIDs(); len(nodes) > 0 && !buo.mutation.BuildToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToEnvironmentTable,
			Columns: build.BuildToEnvironmentPrimaryKey,
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
	if nodes := buo.mutation.BuildToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToEnvironmentTable,
			Columns: build.BuildToEnvironmentPrimaryKey,
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
	if buo.mutation.BuildToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToPlanTable,
			Columns: build.BuildToPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedBuildToPlanIDs(); len(nodes) > 0 && !buo.mutation.BuildToPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToPlanTable,
			Columns: build.BuildToPlanPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BuildToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   build.BuildToPlanTable,
			Columns: build.BuildToPlanPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Build{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{build.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
