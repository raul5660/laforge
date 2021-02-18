// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/competition"
	"github.com/gen0cide/laforge/ent/dns"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// CompetitionUpdate is the builder for updating Competition entities.
type CompetitionUpdate struct {
	config
	hooks    []Hook
	mutation *CompetitionMutation
}

// Where adds a new predicate for the CompetitionUpdate builder.
func (cu *CompetitionUpdate) Where(ps ...predicate.Competition) *CompetitionUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetHclID sets the "hcl_id" field.
func (cu *CompetitionUpdate) SetHclID(s string) *CompetitionUpdate {
	cu.mutation.SetHclID(s)
	return cu
}

// SetRootPassword sets the "root_password" field.
func (cu *CompetitionUpdate) SetRootPassword(s string) *CompetitionUpdate {
	cu.mutation.SetRootPassword(s)
	return cu
}

// SetConfig sets the "config" field.
func (cu *CompetitionUpdate) SetConfig(m map[string]string) *CompetitionUpdate {
	cu.mutation.SetConfig(m)
	return cu
}

// SetTags sets the "tags" field.
func (cu *CompetitionUpdate) SetTags(m map[string]string) *CompetitionUpdate {
	cu.mutation.SetTags(m)
	return cu
}

// AddCompetitionToTagIDs adds the "CompetitionToTag" edge to the Tag entity by IDs.
func (cu *CompetitionUpdate) AddCompetitionToTagIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.AddCompetitionToTagIDs(ids...)
	return cu
}

// AddCompetitionToTag adds the "CompetitionToTag" edges to the Tag entity.
func (cu *CompetitionUpdate) AddCompetitionToTag(t ...*Tag) *CompetitionUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddCompetitionToTagIDs(ids...)
}

// AddCompetitionToDNSIDs adds the "CompetitionToDNS" edge to the DNS entity by IDs.
func (cu *CompetitionUpdate) AddCompetitionToDNSIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.AddCompetitionToDNSIDs(ids...)
	return cu
}

// AddCompetitionToDNS adds the "CompetitionToDNS" edges to the DNS entity.
func (cu *CompetitionUpdate) AddCompetitionToDNS(d ...*DNS) *CompetitionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.AddCompetitionToDNSIDs(ids...)
}

// AddCompetitionToEnvironmentIDs adds the "CompetitionToEnvironment" edge to the Environment entity by IDs.
func (cu *CompetitionUpdate) AddCompetitionToEnvironmentIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.AddCompetitionToEnvironmentIDs(ids...)
	return cu
}

// AddCompetitionToEnvironment adds the "CompetitionToEnvironment" edges to the Environment entity.
func (cu *CompetitionUpdate) AddCompetitionToEnvironment(e ...*Environment) *CompetitionUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.AddCompetitionToEnvironmentIDs(ids...)
}

// Mutation returns the CompetitionMutation object of the builder.
func (cu *CompetitionUpdate) Mutation() *CompetitionMutation {
	return cu.mutation
}

// ClearCompetitionToTag clears all "CompetitionToTag" edges to the Tag entity.
func (cu *CompetitionUpdate) ClearCompetitionToTag() *CompetitionUpdate {
	cu.mutation.ClearCompetitionToTag()
	return cu
}

// RemoveCompetitionToTagIDs removes the "CompetitionToTag" edge to Tag entities by IDs.
func (cu *CompetitionUpdate) RemoveCompetitionToTagIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.RemoveCompetitionToTagIDs(ids...)
	return cu
}

// RemoveCompetitionToTag removes "CompetitionToTag" edges to Tag entities.
func (cu *CompetitionUpdate) RemoveCompetitionToTag(t ...*Tag) *CompetitionUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveCompetitionToTagIDs(ids...)
}

// ClearCompetitionToDNS clears all "CompetitionToDNS" edges to the DNS entity.
func (cu *CompetitionUpdate) ClearCompetitionToDNS() *CompetitionUpdate {
	cu.mutation.ClearCompetitionToDNS()
	return cu
}

// RemoveCompetitionToDNSIDs removes the "CompetitionToDNS" edge to DNS entities by IDs.
func (cu *CompetitionUpdate) RemoveCompetitionToDNSIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.RemoveCompetitionToDNSIDs(ids...)
	return cu
}

// RemoveCompetitionToDNS removes "CompetitionToDNS" edges to DNS entities.
func (cu *CompetitionUpdate) RemoveCompetitionToDNS(d ...*DNS) *CompetitionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.RemoveCompetitionToDNSIDs(ids...)
}

// ClearCompetitionToEnvironment clears all "CompetitionToEnvironment" edges to the Environment entity.
func (cu *CompetitionUpdate) ClearCompetitionToEnvironment() *CompetitionUpdate {
	cu.mutation.ClearCompetitionToEnvironment()
	return cu
}

// RemoveCompetitionToEnvironmentIDs removes the "CompetitionToEnvironment" edge to Environment entities by IDs.
func (cu *CompetitionUpdate) RemoveCompetitionToEnvironmentIDs(ids ...int) *CompetitionUpdate {
	cu.mutation.RemoveCompetitionToEnvironmentIDs(ids...)
	return cu
}

// RemoveCompetitionToEnvironment removes "CompetitionToEnvironment" edges to Environment entities.
func (cu *CompetitionUpdate) RemoveCompetitionToEnvironment(e ...*Environment) *CompetitionUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.RemoveCompetitionToEnvironmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CompetitionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompetitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompetitionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompetitionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompetitionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CompetitionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   competition.Table,
			Columns: competition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: competition.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldHclID,
		})
	}
	if value, ok := cu.mutation.RootPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldRootPassword,
		})
	}
	if value, ok := cu.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: competition.FieldConfig,
		})
	}
	if value, ok := cu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: competition.FieldTags,
		})
	}
	if cu.mutation.CompetitionToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTagTable,
			Columns: []string{competition.CompetitionToTagColumn},
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
	if nodes := cu.mutation.RemovedCompetitionToTagIDs(); len(nodes) > 0 && !cu.mutation.CompetitionToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTagTable,
			Columns: []string{competition.CompetitionToTagColumn},
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
	if nodes := cu.mutation.CompetitionToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTagTable,
			Columns: []string{competition.CompetitionToTagColumn},
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
	if cu.mutation.CompetitionToDNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToDNSTable,
			Columns: []string{competition.CompetitionToDNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCompetitionToDNSIDs(); len(nodes) > 0 && !cu.mutation.CompetitionToDNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToDNSTable,
			Columns: []string{competition.CompetitionToDNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CompetitionToDNSIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToDNSTable,
			Columns: []string{competition.CompetitionToDNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CompetitionToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   competition.CompetitionToEnvironmentTable,
			Columns: competition.CompetitionToEnvironmentPrimaryKey,
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
	if nodes := cu.mutation.RemovedCompetitionToEnvironmentIDs(); len(nodes) > 0 && !cu.mutation.CompetitionToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   competition.CompetitionToEnvironmentTable,
			Columns: competition.CompetitionToEnvironmentPrimaryKey,
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
	if nodes := cu.mutation.CompetitionToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   competition.CompetitionToEnvironmentTable,
			Columns: competition.CompetitionToEnvironmentPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{competition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CompetitionUpdateOne is the builder for updating a single Competition entity.
type CompetitionUpdateOne struct {
	config
	hooks    []Hook
	mutation *CompetitionMutation
}

// SetHclID sets the "hcl_id" field.
func (cuo *CompetitionUpdateOne) SetHclID(s string) *CompetitionUpdateOne {
	cuo.mutation.SetHclID(s)
	return cuo
}

// SetRootPassword sets the "root_password" field.
func (cuo *CompetitionUpdateOne) SetRootPassword(s string) *CompetitionUpdateOne {
	cuo.mutation.SetRootPassword(s)
	return cuo
}

// SetConfig sets the "config" field.
func (cuo *CompetitionUpdateOne) SetConfig(m map[string]string) *CompetitionUpdateOne {
	cuo.mutation.SetConfig(m)
	return cuo
}

// SetTags sets the "tags" field.
func (cuo *CompetitionUpdateOne) SetTags(m map[string]string) *CompetitionUpdateOne {
	cuo.mutation.SetTags(m)
	return cuo
}

// AddCompetitionToTagIDs adds the "CompetitionToTag" edge to the Tag entity by IDs.
func (cuo *CompetitionUpdateOne) AddCompetitionToTagIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.AddCompetitionToTagIDs(ids...)
	return cuo
}

// AddCompetitionToTag adds the "CompetitionToTag" edges to the Tag entity.
func (cuo *CompetitionUpdateOne) AddCompetitionToTag(t ...*Tag) *CompetitionUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddCompetitionToTagIDs(ids...)
}

// AddCompetitionToDNSIDs adds the "CompetitionToDNS" edge to the DNS entity by IDs.
func (cuo *CompetitionUpdateOne) AddCompetitionToDNSIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.AddCompetitionToDNSIDs(ids...)
	return cuo
}

// AddCompetitionToDNS adds the "CompetitionToDNS" edges to the DNS entity.
func (cuo *CompetitionUpdateOne) AddCompetitionToDNS(d ...*DNS) *CompetitionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.AddCompetitionToDNSIDs(ids...)
}

// AddCompetitionToEnvironmentIDs adds the "CompetitionToEnvironment" edge to the Environment entity by IDs.
func (cuo *CompetitionUpdateOne) AddCompetitionToEnvironmentIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.AddCompetitionToEnvironmentIDs(ids...)
	return cuo
}

// AddCompetitionToEnvironment adds the "CompetitionToEnvironment" edges to the Environment entity.
func (cuo *CompetitionUpdateOne) AddCompetitionToEnvironment(e ...*Environment) *CompetitionUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.AddCompetitionToEnvironmentIDs(ids...)
}

// Mutation returns the CompetitionMutation object of the builder.
func (cuo *CompetitionUpdateOne) Mutation() *CompetitionMutation {
	return cuo.mutation
}

// ClearCompetitionToTag clears all "CompetitionToTag" edges to the Tag entity.
func (cuo *CompetitionUpdateOne) ClearCompetitionToTag() *CompetitionUpdateOne {
	cuo.mutation.ClearCompetitionToTag()
	return cuo
}

// RemoveCompetitionToTagIDs removes the "CompetitionToTag" edge to Tag entities by IDs.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToTagIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.RemoveCompetitionToTagIDs(ids...)
	return cuo
}

// RemoveCompetitionToTag removes "CompetitionToTag" edges to Tag entities.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToTag(t ...*Tag) *CompetitionUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveCompetitionToTagIDs(ids...)
}

// ClearCompetitionToDNS clears all "CompetitionToDNS" edges to the DNS entity.
func (cuo *CompetitionUpdateOne) ClearCompetitionToDNS() *CompetitionUpdateOne {
	cuo.mutation.ClearCompetitionToDNS()
	return cuo
}

// RemoveCompetitionToDNSIDs removes the "CompetitionToDNS" edge to DNS entities by IDs.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToDNSIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.RemoveCompetitionToDNSIDs(ids...)
	return cuo
}

// RemoveCompetitionToDNS removes "CompetitionToDNS" edges to DNS entities.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToDNS(d ...*DNS) *CompetitionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.RemoveCompetitionToDNSIDs(ids...)
}

// ClearCompetitionToEnvironment clears all "CompetitionToEnvironment" edges to the Environment entity.
func (cuo *CompetitionUpdateOne) ClearCompetitionToEnvironment() *CompetitionUpdateOne {
	cuo.mutation.ClearCompetitionToEnvironment()
	return cuo
}

// RemoveCompetitionToEnvironmentIDs removes the "CompetitionToEnvironment" edge to Environment entities by IDs.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToEnvironmentIDs(ids ...int) *CompetitionUpdateOne {
	cuo.mutation.RemoveCompetitionToEnvironmentIDs(ids...)
	return cuo
}

// RemoveCompetitionToEnvironment removes "CompetitionToEnvironment" edges to Environment entities.
func (cuo *CompetitionUpdateOne) RemoveCompetitionToEnvironment(e ...*Environment) *CompetitionUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.RemoveCompetitionToEnvironmentIDs(ids...)
}

// Save executes the query and returns the updated Competition entity.
func (cuo *CompetitionUpdateOne) Save(ctx context.Context) (*Competition, error) {
	var (
		err  error
		node *Competition
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompetitionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompetitionUpdateOne) SaveX(ctx context.Context) *Competition {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CompetitionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompetitionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CompetitionUpdateOne) sqlSave(ctx context.Context) (_node *Competition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   competition.Table,
			Columns: competition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: competition.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Competition.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldHclID,
		})
	}
	if value, ok := cuo.mutation.RootPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: competition.FieldRootPassword,
		})
	}
	if value, ok := cuo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: competition.FieldConfig,
		})
	}
	if value, ok := cuo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: competition.FieldTags,
		})
	}
	if cuo.mutation.CompetitionToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTagTable,
			Columns: []string{competition.CompetitionToTagColumn},
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
	if nodes := cuo.mutation.RemovedCompetitionToTagIDs(); len(nodes) > 0 && !cuo.mutation.CompetitionToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTagTable,
			Columns: []string{competition.CompetitionToTagColumn},
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
	if nodes := cuo.mutation.CompetitionToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToTagTable,
			Columns: []string{competition.CompetitionToTagColumn},
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
	if cuo.mutation.CompetitionToDNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToDNSTable,
			Columns: []string{competition.CompetitionToDNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCompetitionToDNSIDs(); len(nodes) > 0 && !cuo.mutation.CompetitionToDNSCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToDNSTable,
			Columns: []string{competition.CompetitionToDNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CompetitionToDNSIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   competition.CompetitionToDNSTable,
			Columns: []string{competition.CompetitionToDNSColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dns.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CompetitionToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   competition.CompetitionToEnvironmentTable,
			Columns: competition.CompetitionToEnvironmentPrimaryKey,
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
	if nodes := cuo.mutation.RemovedCompetitionToEnvironmentIDs(); len(nodes) > 0 && !cuo.mutation.CompetitionToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   competition.CompetitionToEnvironmentTable,
			Columns: competition.CompetitionToEnvironmentPrimaryKey,
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
	if nodes := cuo.mutation.CompetitionToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   competition.CompetitionToEnvironmentTable,
			Columns: competition.CompetitionToEnvironmentPrimaryKey,
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
	_node = &Competition{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{competition.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
