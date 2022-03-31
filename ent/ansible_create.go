// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/ansible"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/user"
	"github.com/google/uuid"
)

// AnsibleCreate is the builder for creating a Ansible entity.
type AnsibleCreate struct {
	config
	mutation *AnsibleMutation
	hooks    []Hook
}

// SetHclID sets the "hcl_id" field.
func (ac *AnsibleCreate) SetHclID(s string) *AnsibleCreate {
	ac.mutation.SetHclID(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AnsibleCreate) SetDescription(s string) *AnsibleCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetSource sets the "source" field.
func (ac *AnsibleCreate) SetSource(s string) *AnsibleCreate {
	ac.mutation.SetSource(s)
	return ac
}

// SetPlaybookName sets the "playbook_name" field.
func (ac *AnsibleCreate) SetPlaybookName(s string) *AnsibleCreate {
	ac.mutation.SetPlaybookName(s)
	return ac
}

// SetMethod sets the "method" field.
func (ac *AnsibleCreate) SetMethod(a ansible.Method) *AnsibleCreate {
	ac.mutation.SetMethod(a)
	return ac
}

// SetInventory sets the "inventory" field.
func (ac *AnsibleCreate) SetInventory(s string) *AnsibleCreate {
	ac.mutation.SetInventory(s)
	return ac
}

// SetTags sets the "tags" field.
func (ac *AnsibleCreate) SetTags(m map[string]string) *AnsibleCreate {
	ac.mutation.SetTags(m)
	return ac
}

// SetID sets the "id" field.
func (ac *AnsibleCreate) SetID(u uuid.UUID) *AnsibleCreate {
	ac.mutation.SetID(u)
	return ac
}

// AddAnsibleToUserIDs adds the "AnsibleToUser" edge to the User entity by IDs.
func (ac *AnsibleCreate) AddAnsibleToUserIDs(ids ...uuid.UUID) *AnsibleCreate {
	ac.mutation.AddAnsibleToUserIDs(ids...)
	return ac
}

// AddAnsibleToUser adds the "AnsibleToUser" edges to the User entity.
func (ac *AnsibleCreate) AddAnsibleToUser(u ...*User) *AnsibleCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ac.AddAnsibleToUserIDs(ids...)
}

// SetAnsibleFromEnvironmentID sets the "AnsibleFromEnvironment" edge to the Environment entity by ID.
func (ac *AnsibleCreate) SetAnsibleFromEnvironmentID(id uuid.UUID) *AnsibleCreate {
	ac.mutation.SetAnsibleFromEnvironmentID(id)
	return ac
}

// SetNillableAnsibleFromEnvironmentID sets the "AnsibleFromEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (ac *AnsibleCreate) SetNillableAnsibleFromEnvironmentID(id *uuid.UUID) *AnsibleCreate {
	if id != nil {
		ac = ac.SetAnsibleFromEnvironmentID(*id)
	}
	return ac
}

// SetAnsibleFromEnvironment sets the "AnsibleFromEnvironment" edge to the Environment entity.
func (ac *AnsibleCreate) SetAnsibleFromEnvironment(e *Environment) *AnsibleCreate {
	return ac.SetAnsibleFromEnvironmentID(e.ID)
}

// Mutation returns the AnsibleMutation object of the builder.
func (ac *AnsibleCreate) Mutation() *AnsibleMutation {
	return ac.mutation
}

// Save creates the Ansible in the database.
func (ac *AnsibleCreate) Save(ctx context.Context) (*Ansible, error) {
	var (
		err  error
		node *Ansible
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnsibleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnsibleCreate) SaveX(ctx context.Context) *Ansible {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnsibleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnsibleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AnsibleCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := ansible.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnsibleCreate) check() error {
	if _, ok := ac.mutation.HclID(); !ok {
		return &ValidationError{Name: "hcl_id", err: errors.New(`ent: missing required field "hcl_id"`)}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if _, ok := ac.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "source"`)}
	}
	if _, ok := ac.mutation.PlaybookName(); !ok {
		return &ValidationError{Name: "playbook_name", err: errors.New(`ent: missing required field "playbook_name"`)}
	}
	if _, ok := ac.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "method"`)}
	}
	if v, ok := ac.mutation.Method(); ok {
		if err := ansible.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "method": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Inventory(); !ok {
		return &ValidationError{Name: "inventory", err: errors.New(`ent: missing required field "inventory"`)}
	}
	if _, ok := ac.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "tags"`)}
	}
	return nil
}

func (ac *AnsibleCreate) sqlSave(ctx context.Context) (*Ansible, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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

func (ac *AnsibleCreate) createSpec() (*Ansible, *sqlgraph.CreateSpec) {
	var (
		_node = &Ansible{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ansible.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ansible.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.HclID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ansible.FieldHclID,
		})
		_node.HclID = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ansible.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := ac.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ansible.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := ac.mutation.PlaybookName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ansible.FieldPlaybookName,
		})
		_node.PlaybookName = value
	}
	if value, ok := ac.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: ansible.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := ac.mutation.Inventory(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ansible.FieldInventory,
		})
		_node.Inventory = value
	}
	if value, ok := ac.mutation.Tags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ansible.FieldTags,
		})
		_node.Tags = value
	}
	if nodes := ac.mutation.AnsibleToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ansible.AnsibleToUserTable,
			Columns: []string{ansible.AnsibleToUserColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AnsibleFromEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ansible.AnsibleFromEnvironmentTable,
			Columns: []string{ansible.AnsibleFromEnvironmentColumn},
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
		_node.environment_environment_to_ansible = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AnsibleCreateBulk is the builder for creating many Ansible entities in bulk.
type AnsibleCreateBulk struct {
	config
	builders []*AnsibleCreate
}

// Save creates the Ansible entities in the database.
func (acb *AnsibleCreateBulk) Save(ctx context.Context) ([]*Ansible, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Ansible, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnsibleMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnsibleCreateBulk) SaveX(ctx context.Context) []*Ansible {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnsibleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnsibleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}