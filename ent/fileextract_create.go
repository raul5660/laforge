// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/fileextract"
	"github.com/gen0cide/laforge/ent/tag"
)

// FileExtractCreate is the builder for creating a FileExtract entity.
type FileExtractCreate struct {
	config
	mutation *FileExtractMutation
	hooks    []Hook
}

// SetSource sets the source field.
func (fec *FileExtractCreate) SetSource(s string) *FileExtractCreate {
	fec.mutation.SetSource(s)
	return fec
}

// SetDestination sets the destination field.
func (fec *FileExtractCreate) SetDestination(s string) *FileExtractCreate {
	fec.mutation.SetDestination(s)
	return fec
}

// SetType sets the type field.
func (fec *FileExtractCreate) SetType(s string) *FileExtractCreate {
	fec.mutation.SetType(s)
	return fec
}

// SetTags sets the tags field.
func (fec *FileExtractCreate) SetTags(m map[string]string) *FileExtractCreate {
	fec.mutation.SetTags(m)
	return fec
}

// AddFileExtractToTagIDs adds the FileExtractToTag edge to Tag by ids.
func (fec *FileExtractCreate) AddFileExtractToTagIDs(ids ...int) *FileExtractCreate {
	fec.mutation.AddFileExtractToTagIDs(ids...)
	return fec
}

// AddFileExtractToTag adds the FileExtractToTag edges to Tag.
func (fec *FileExtractCreate) AddFileExtractToTag(t ...*Tag) *FileExtractCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fec.AddFileExtractToTagIDs(ids...)
}

// Mutation returns the FileExtractMutation object of the builder.
func (fec *FileExtractCreate) Mutation() *FileExtractMutation {
	return fec.mutation
}

// Save creates the FileExtract in the database.
func (fec *FileExtractCreate) Save(ctx context.Context) (*FileExtract, error) {
	var (
		err  error
		node *FileExtract
	)
	if len(fec.hooks) == 0 {
		if err = fec.check(); err != nil {
			return nil, err
		}
		node, err = fec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileExtractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fec.check(); err != nil {
				return nil, err
			}
			fec.mutation = mutation
			node, err = fec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fec.hooks) - 1; i >= 0; i-- {
			mut = fec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fec *FileExtractCreate) SaveX(ctx context.Context) *FileExtract {
	v, err := fec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (fec *FileExtractCreate) check() error {
	if _, ok := fec.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New("ent: missing required field \"source\"")}
	}
	if _, ok := fec.mutation.Destination(); !ok {
		return &ValidationError{Name: "destination", err: errors.New("ent: missing required field \"destination\"")}
	}
	if _, ok := fec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if _, ok := fec.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New("ent: missing required field \"tags\"")}
	}
	return nil
}

func (fec *FileExtractCreate) sqlSave(ctx context.Context) (*FileExtract, error) {
	_node, _spec := fec.createSpec()
	if err := sqlgraph.CreateNode(ctx, fec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fec *FileExtractCreate) createSpec() (*FileExtract, *sqlgraph.CreateSpec) {
	var (
		_node = &FileExtract{config: fec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fileextract.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fileextract.FieldID,
			},
		}
	)
	if value, ok := fec.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := fec.mutation.Destination(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldDestination,
		})
		_node.Destination = value
	}
	if value, ok := fec.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldType,
		})
		_node.Type = value
	}
	if value, ok := fec.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: fileextract.FieldTags,
		})
		_node.Tags = value
	}
	if nodes := fec.mutation.FileExtractToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileextract.FileExtractToTagTable,
			Columns: []string{fileextract.FileExtractToTagColumn},
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
	return _node, _spec
}

// FileExtractCreateBulk is the builder for creating a bulk of FileExtract entities.
type FileExtractCreateBulk struct {
	config
	builders []*FileExtractCreate
}

// Save creates the FileExtract entities in the database.
func (fecb *FileExtractCreateBulk) Save(ctx context.Context) ([]*FileExtract, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fecb.builders))
	nodes := make([]*FileExtract, len(fecb.builders))
	mutators := make([]Mutator, len(fecb.builders))
	for i := range fecb.builders {
		func(i int, root context.Context) {
			builder := fecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileExtractMutation)
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
					_, err = mutators[i+1].Mutate(root, fecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (fecb *FileExtractCreateBulk) SaveX(ctx context.Context) []*FileExtract {
	v, err := fecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
