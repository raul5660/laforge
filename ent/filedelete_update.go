// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/filedelete"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// FileDeleteUpdate is the builder for updating FileDelete entities.
type FileDeleteUpdate struct {
	config
	hooks    []Hook
	mutation *FileDeleteMutation
}

// Where adds a new predicate for the builder.
func (fdu *FileDeleteUpdate) Where(ps ...predicate.FileDelete) *FileDeleteUpdate {
	fdu.mutation.predicates = append(fdu.mutation.predicates, ps...)
	return fdu
}

// SetPath sets the path field.
func (fdu *FileDeleteUpdate) SetPath(s string) *FileDeleteUpdate {
	fdu.mutation.SetPath(s)
	return fdu
}

// SetTags sets the tags field.
func (fdu *FileDeleteUpdate) SetTags(m map[string]string) *FileDeleteUpdate {
	fdu.mutation.SetTags(m)
	return fdu
}

// AddFileDeleteToTagIDs adds the FileDeleteToTag edge to Tag by ids.
func (fdu *FileDeleteUpdate) AddFileDeleteToTagIDs(ids ...int) *FileDeleteUpdate {
	fdu.mutation.AddFileDeleteToTagIDs(ids...)
	return fdu
}

// AddFileDeleteToTag adds the FileDeleteToTag edges to Tag.
func (fdu *FileDeleteUpdate) AddFileDeleteToTag(t ...*Tag) *FileDeleteUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fdu.AddFileDeleteToTagIDs(ids...)
}

// Mutation returns the FileDeleteMutation object of the builder.
func (fdu *FileDeleteUpdate) Mutation() *FileDeleteMutation {
	return fdu.mutation
}

// ClearFileDeleteToTag clears all "FileDeleteToTag" edges to type Tag.
func (fdu *FileDeleteUpdate) ClearFileDeleteToTag() *FileDeleteUpdate {
	fdu.mutation.ClearFileDeleteToTag()
	return fdu
}

// RemoveFileDeleteToTagIDs removes the FileDeleteToTag edge to Tag by ids.
func (fdu *FileDeleteUpdate) RemoveFileDeleteToTagIDs(ids ...int) *FileDeleteUpdate {
	fdu.mutation.RemoveFileDeleteToTagIDs(ids...)
	return fdu
}

// RemoveFileDeleteToTag removes FileDeleteToTag edges to Tag.
func (fdu *FileDeleteUpdate) RemoveFileDeleteToTag(t ...*Tag) *FileDeleteUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fdu.RemoveFileDeleteToTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fdu *FileDeleteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fdu.hooks) == 0 {
		affected, err = fdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileDeleteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fdu.mutation = mutation
			affected, err = fdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fdu.hooks) - 1; i >= 0; i-- {
			mut = fdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fdu *FileDeleteUpdate) SaveX(ctx context.Context) int {
	affected, err := fdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fdu *FileDeleteUpdate) Exec(ctx context.Context) error {
	_, err := fdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdu *FileDeleteUpdate) ExecX(ctx context.Context) {
	if err := fdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fdu *FileDeleteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filedelete.Table,
			Columns: filedelete.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filedelete.FieldID,
			},
		},
	}
	if ps := fdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fdu.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedelete.FieldPath,
		})
	}
	if value, ok := fdu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: filedelete.FieldTags,
		})
	}
	if fdu.mutation.FileDeleteToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedelete.FileDeleteToTagTable,
			Columns: []string{filedelete.FileDeleteToTagColumn},
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
	if nodes := fdu.mutation.RemovedFileDeleteToTagIDs(); len(nodes) > 0 && !fdu.mutation.FileDeleteToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedelete.FileDeleteToTagTable,
			Columns: []string{filedelete.FileDeleteToTagColumn},
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
	if nodes := fdu.mutation.FileDeleteToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedelete.FileDeleteToTagTable,
			Columns: []string{filedelete.FileDeleteToTagColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filedelete.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FileDeleteUpdateOne is the builder for updating a single FileDelete entity.
type FileDeleteUpdateOne struct {
	config
	hooks    []Hook
	mutation *FileDeleteMutation
}

// SetPath sets the path field.
func (fduo *FileDeleteUpdateOne) SetPath(s string) *FileDeleteUpdateOne {
	fduo.mutation.SetPath(s)
	return fduo
}

// SetTags sets the tags field.
func (fduo *FileDeleteUpdateOne) SetTags(m map[string]string) *FileDeleteUpdateOne {
	fduo.mutation.SetTags(m)
	return fduo
}

// AddFileDeleteToTagIDs adds the FileDeleteToTag edge to Tag by ids.
func (fduo *FileDeleteUpdateOne) AddFileDeleteToTagIDs(ids ...int) *FileDeleteUpdateOne {
	fduo.mutation.AddFileDeleteToTagIDs(ids...)
	return fduo
}

// AddFileDeleteToTag adds the FileDeleteToTag edges to Tag.
func (fduo *FileDeleteUpdateOne) AddFileDeleteToTag(t ...*Tag) *FileDeleteUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fduo.AddFileDeleteToTagIDs(ids...)
}

// Mutation returns the FileDeleteMutation object of the builder.
func (fduo *FileDeleteUpdateOne) Mutation() *FileDeleteMutation {
	return fduo.mutation
}

// ClearFileDeleteToTag clears all "FileDeleteToTag" edges to type Tag.
func (fduo *FileDeleteUpdateOne) ClearFileDeleteToTag() *FileDeleteUpdateOne {
	fduo.mutation.ClearFileDeleteToTag()
	return fduo
}

// RemoveFileDeleteToTagIDs removes the FileDeleteToTag edge to Tag by ids.
func (fduo *FileDeleteUpdateOne) RemoveFileDeleteToTagIDs(ids ...int) *FileDeleteUpdateOne {
	fduo.mutation.RemoveFileDeleteToTagIDs(ids...)
	return fduo
}

// RemoveFileDeleteToTag removes FileDeleteToTag edges to Tag.
func (fduo *FileDeleteUpdateOne) RemoveFileDeleteToTag(t ...*Tag) *FileDeleteUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fduo.RemoveFileDeleteToTagIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (fduo *FileDeleteUpdateOne) Save(ctx context.Context) (*FileDelete, error) {
	var (
		err  error
		node *FileDelete
	)
	if len(fduo.hooks) == 0 {
		node, err = fduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileDeleteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fduo.mutation = mutation
			node, err = fduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fduo.hooks) - 1; i >= 0; i-- {
			mut = fduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fduo *FileDeleteUpdateOne) SaveX(ctx context.Context) *FileDelete {
	node, err := fduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fduo *FileDeleteUpdateOne) Exec(ctx context.Context) error {
	_, err := fduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fduo *FileDeleteUpdateOne) ExecX(ctx context.Context) {
	if err := fduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fduo *FileDeleteUpdateOne) sqlSave(ctx context.Context) (_node *FileDelete, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filedelete.Table,
			Columns: filedelete.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filedelete.FieldID,
			},
		},
	}
	id, ok := fduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FileDelete.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fduo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedelete.FieldPath,
		})
	}
	if value, ok := fduo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: filedelete.FieldTags,
		})
	}
	if fduo.mutation.FileDeleteToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedelete.FileDeleteToTagTable,
			Columns: []string{filedelete.FileDeleteToTagColumn},
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
	if nodes := fduo.mutation.RemovedFileDeleteToTagIDs(); len(nodes) > 0 && !fduo.mutation.FileDeleteToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedelete.FileDeleteToTagTable,
			Columns: []string{filedelete.FileDeleteToTagColumn},
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
	if nodes := fduo.mutation.FileDeleteToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedelete.FileDeleteToTagTable,
			Columns: []string{filedelete.FileDeleteToTagColumn},
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
	_node = &FileDelete{config: fduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filedelete.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
