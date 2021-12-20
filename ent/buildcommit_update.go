// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/buildcommit"
	"github.com/gen0cide/laforge/ent/plandiff"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// BuildCommitUpdate is the builder for updating BuildCommit entities.
type BuildCommitUpdate struct {
	config
	hooks    []Hook
	mutation *BuildCommitMutation
}

// Where appends a list predicates to the BuildCommitUpdate builder.
func (bcu *BuildCommitUpdate) Where(ps ...predicate.BuildCommit) *BuildCommitUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetType sets the "type" field.
func (bcu *BuildCommitUpdate) SetType(b buildcommit.Type) *BuildCommitUpdate {
	bcu.mutation.SetType(b)
	return bcu
}

// SetRevision sets the "revision" field.
func (bcu *BuildCommitUpdate) SetRevision(i int) *BuildCommitUpdate {
	bcu.mutation.ResetRevision()
	bcu.mutation.SetRevision(i)
	return bcu
}

// AddRevision adds i to the "revision" field.
func (bcu *BuildCommitUpdate) AddRevision(i int) *BuildCommitUpdate {
	bcu.mutation.AddRevision(i)
	return bcu
}

// SetState sets the "state" field.
func (bcu *BuildCommitUpdate) SetState(b buildcommit.State) *BuildCommitUpdate {
	bcu.mutation.SetState(b)
	return bcu
}

// SetCreatedAt sets the "created_at" field.
func (bcu *BuildCommitUpdate) SetCreatedAt(t time.Time) *BuildCommitUpdate {
	bcu.mutation.SetCreatedAt(t)
	return bcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bcu *BuildCommitUpdate) SetNillableCreatedAt(t *time.Time) *BuildCommitUpdate {
	if t != nil {
		bcu.SetCreatedAt(*t)
	}
	return bcu
}

// SetBuildCommitToBuildID sets the "BuildCommitToBuild" edge to the Build entity by ID.
func (bcu *BuildCommitUpdate) SetBuildCommitToBuildID(id uuid.UUID) *BuildCommitUpdate {
	bcu.mutation.SetBuildCommitToBuildID(id)
	return bcu
}

// SetBuildCommitToBuild sets the "BuildCommitToBuild" edge to the Build entity.
func (bcu *BuildCommitUpdate) SetBuildCommitToBuild(b *Build) *BuildCommitUpdate {
	return bcu.SetBuildCommitToBuildID(b.ID)
}

// AddBuildCommitToPlanDiffIDs adds the "BuildCommitToPlanDiffs" edge to the PlanDiff entity by IDs.
func (bcu *BuildCommitUpdate) AddBuildCommitToPlanDiffIDs(ids ...uuid.UUID) *BuildCommitUpdate {
	bcu.mutation.AddBuildCommitToPlanDiffIDs(ids...)
	return bcu
}

// AddBuildCommitToPlanDiffs adds the "BuildCommitToPlanDiffs" edges to the PlanDiff entity.
func (bcu *BuildCommitUpdate) AddBuildCommitToPlanDiffs(p ...*PlanDiff) *BuildCommitUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bcu.AddBuildCommitToPlanDiffIDs(ids...)
}

// Mutation returns the BuildCommitMutation object of the builder.
func (bcu *BuildCommitUpdate) Mutation() *BuildCommitMutation {
	return bcu.mutation
}

// ClearBuildCommitToBuild clears the "BuildCommitToBuild" edge to the Build entity.
func (bcu *BuildCommitUpdate) ClearBuildCommitToBuild() *BuildCommitUpdate {
	bcu.mutation.ClearBuildCommitToBuild()
	return bcu
}

// ClearBuildCommitToPlanDiffs clears all "BuildCommitToPlanDiffs" edges to the PlanDiff entity.
func (bcu *BuildCommitUpdate) ClearBuildCommitToPlanDiffs() *BuildCommitUpdate {
	bcu.mutation.ClearBuildCommitToPlanDiffs()
	return bcu
}

// RemoveBuildCommitToPlanDiffIDs removes the "BuildCommitToPlanDiffs" edge to PlanDiff entities by IDs.
func (bcu *BuildCommitUpdate) RemoveBuildCommitToPlanDiffIDs(ids ...uuid.UUID) *BuildCommitUpdate {
	bcu.mutation.RemoveBuildCommitToPlanDiffIDs(ids...)
	return bcu
}

// RemoveBuildCommitToPlanDiffs removes "BuildCommitToPlanDiffs" edges to PlanDiff entities.
func (bcu *BuildCommitUpdate) RemoveBuildCommitToPlanDiffs(p ...*PlanDiff) *BuildCommitUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bcu.RemoveBuildCommitToPlanDiffIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BuildCommitUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcu.hooks) == 0 {
		if err = bcu.check(); err != nil {
			return 0, err
		}
		affected, err = bcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildCommitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcu.check(); err != nil {
				return 0, err
			}
			bcu.mutation = mutation
			affected, err = bcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcu.hooks) - 1; i >= 0; i-- {
			if bcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcu *BuildCommitUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BuildCommitUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BuildCommitUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcu *BuildCommitUpdate) check() error {
	if v, ok := bcu.mutation.GetType(); ok {
		if err := buildcommit.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := bcu.mutation.State(); ok {
		if err := buildcommit.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if _, ok := bcu.mutation.BuildCommitToBuildID(); bcu.mutation.BuildCommitToBuildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"BuildCommitToBuild\"")
	}
	return nil
}

func (bcu *BuildCommitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   buildcommit.Table,
			Columns: buildcommit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: buildcommit.FieldID,
			},
		},
	}
	if ps := bcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: buildcommit.FieldType,
		})
	}
	if value, ok := bcu.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: buildcommit.FieldRevision,
		})
	}
	if value, ok := bcu.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: buildcommit.FieldRevision,
		})
	}
	if value, ok := bcu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: buildcommit.FieldState,
		})
	}
	if value, ok := bcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: buildcommit.FieldCreatedAt,
		})
	}
	if bcu.mutation.BuildCommitToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   buildcommit.BuildCommitToBuildTable,
			Columns: []string{buildcommit.BuildCommitToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.BuildCommitToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   buildcommit.BuildCommitToBuildTable,
			Columns: []string{buildcommit.BuildCommitToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcu.mutation.BuildCommitToPlanDiffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToPlanDiffsTable,
			Columns: []string{buildcommit.BuildCommitToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.RemovedBuildCommitToPlanDiffsIDs(); len(nodes) > 0 && !bcu.mutation.BuildCommitToPlanDiffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToPlanDiffsTable,
			Columns: []string{buildcommit.BuildCommitToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.BuildCommitToPlanDiffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToPlanDiffsTable,
			Columns: []string{buildcommit.BuildCommitToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{buildcommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BuildCommitUpdateOne is the builder for updating a single BuildCommit entity.
type BuildCommitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BuildCommitMutation
}

// SetType sets the "type" field.
func (bcuo *BuildCommitUpdateOne) SetType(b buildcommit.Type) *BuildCommitUpdateOne {
	bcuo.mutation.SetType(b)
	return bcuo
}

// SetRevision sets the "revision" field.
func (bcuo *BuildCommitUpdateOne) SetRevision(i int) *BuildCommitUpdateOne {
	bcuo.mutation.ResetRevision()
	bcuo.mutation.SetRevision(i)
	return bcuo
}

// AddRevision adds i to the "revision" field.
func (bcuo *BuildCommitUpdateOne) AddRevision(i int) *BuildCommitUpdateOne {
	bcuo.mutation.AddRevision(i)
	return bcuo
}

// SetState sets the "state" field.
func (bcuo *BuildCommitUpdateOne) SetState(b buildcommit.State) *BuildCommitUpdateOne {
	bcuo.mutation.SetState(b)
	return bcuo
}

// SetCreatedAt sets the "created_at" field.
func (bcuo *BuildCommitUpdateOne) SetCreatedAt(t time.Time) *BuildCommitUpdateOne {
	bcuo.mutation.SetCreatedAt(t)
	return bcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bcuo *BuildCommitUpdateOne) SetNillableCreatedAt(t *time.Time) *BuildCommitUpdateOne {
	if t != nil {
		bcuo.SetCreatedAt(*t)
	}
	return bcuo
}

// SetBuildCommitToBuildID sets the "BuildCommitToBuild" edge to the Build entity by ID.
func (bcuo *BuildCommitUpdateOne) SetBuildCommitToBuildID(id uuid.UUID) *BuildCommitUpdateOne {
	bcuo.mutation.SetBuildCommitToBuildID(id)
	return bcuo
}

// SetBuildCommitToBuild sets the "BuildCommitToBuild" edge to the Build entity.
func (bcuo *BuildCommitUpdateOne) SetBuildCommitToBuild(b *Build) *BuildCommitUpdateOne {
	return bcuo.SetBuildCommitToBuildID(b.ID)
}

// AddBuildCommitToPlanDiffIDs adds the "BuildCommitToPlanDiffs" edge to the PlanDiff entity by IDs.
func (bcuo *BuildCommitUpdateOne) AddBuildCommitToPlanDiffIDs(ids ...uuid.UUID) *BuildCommitUpdateOne {
	bcuo.mutation.AddBuildCommitToPlanDiffIDs(ids...)
	return bcuo
}

// AddBuildCommitToPlanDiffs adds the "BuildCommitToPlanDiffs" edges to the PlanDiff entity.
func (bcuo *BuildCommitUpdateOne) AddBuildCommitToPlanDiffs(p ...*PlanDiff) *BuildCommitUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bcuo.AddBuildCommitToPlanDiffIDs(ids...)
}

// Mutation returns the BuildCommitMutation object of the builder.
func (bcuo *BuildCommitUpdateOne) Mutation() *BuildCommitMutation {
	return bcuo.mutation
}

// ClearBuildCommitToBuild clears the "BuildCommitToBuild" edge to the Build entity.
func (bcuo *BuildCommitUpdateOne) ClearBuildCommitToBuild() *BuildCommitUpdateOne {
	bcuo.mutation.ClearBuildCommitToBuild()
	return bcuo
}

// ClearBuildCommitToPlanDiffs clears all "BuildCommitToPlanDiffs" edges to the PlanDiff entity.
func (bcuo *BuildCommitUpdateOne) ClearBuildCommitToPlanDiffs() *BuildCommitUpdateOne {
	bcuo.mutation.ClearBuildCommitToPlanDiffs()
	return bcuo
}

// RemoveBuildCommitToPlanDiffIDs removes the "BuildCommitToPlanDiffs" edge to PlanDiff entities by IDs.
func (bcuo *BuildCommitUpdateOne) RemoveBuildCommitToPlanDiffIDs(ids ...uuid.UUID) *BuildCommitUpdateOne {
	bcuo.mutation.RemoveBuildCommitToPlanDiffIDs(ids...)
	return bcuo
}

// RemoveBuildCommitToPlanDiffs removes "BuildCommitToPlanDiffs" edges to PlanDiff entities.
func (bcuo *BuildCommitUpdateOne) RemoveBuildCommitToPlanDiffs(p ...*PlanDiff) *BuildCommitUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bcuo.RemoveBuildCommitToPlanDiffIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BuildCommitUpdateOne) Select(field string, fields ...string) *BuildCommitUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BuildCommit entity.
func (bcuo *BuildCommitUpdateOne) Save(ctx context.Context) (*BuildCommit, error) {
	var (
		err  error
		node *BuildCommit
	)
	if len(bcuo.hooks) == 0 {
		if err = bcuo.check(); err != nil {
			return nil, err
		}
		node, err = bcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildCommitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcuo.check(); err != nil {
				return nil, err
			}
			bcuo.mutation = mutation
			node, err = bcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bcuo.hooks) - 1; i >= 0; i-- {
			if bcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcuo *BuildCommitUpdateOne) SaveX(ctx context.Context) *BuildCommit {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BuildCommitUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BuildCommitUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcuo *BuildCommitUpdateOne) check() error {
	if v, ok := bcuo.mutation.GetType(); ok {
		if err := buildcommit.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := bcuo.mutation.State(); ok {
		if err := buildcommit.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if _, ok := bcuo.mutation.BuildCommitToBuildID(); bcuo.mutation.BuildCommitToBuildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"BuildCommitToBuild\"")
	}
	return nil
}

func (bcuo *BuildCommitUpdateOne) sqlSave(ctx context.Context) (_node *BuildCommit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   buildcommit.Table,
			Columns: buildcommit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: buildcommit.FieldID,
			},
		},
	}
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing BuildCommit.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, buildcommit.FieldID)
		for _, f := range fields {
			if !buildcommit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != buildcommit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: buildcommit.FieldType,
		})
	}
	if value, ok := bcuo.mutation.Revision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: buildcommit.FieldRevision,
		})
	}
	if value, ok := bcuo.mutation.AddedRevision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: buildcommit.FieldRevision,
		})
	}
	if value, ok := bcuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: buildcommit.FieldState,
		})
	}
	if value, ok := bcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: buildcommit.FieldCreatedAt,
		})
	}
	if bcuo.mutation.BuildCommitToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   buildcommit.BuildCommitToBuildTable,
			Columns: []string{buildcommit.BuildCommitToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.BuildCommitToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   buildcommit.BuildCommitToBuildTable,
			Columns: []string{buildcommit.BuildCommitToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bcuo.mutation.BuildCommitToPlanDiffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToPlanDiffsTable,
			Columns: []string{buildcommit.BuildCommitToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.RemovedBuildCommitToPlanDiffsIDs(); len(nodes) > 0 && !bcuo.mutation.BuildCommitToPlanDiffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToPlanDiffsTable,
			Columns: []string{buildcommit.BuildCommitToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.BuildCommitToPlanDiffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   buildcommit.BuildCommitToPlanDiffsTable,
			Columns: []string{buildcommit.BuildCommitToPlanDiffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plandiff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BuildCommit{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{buildcommit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
