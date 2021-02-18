// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/dnsrecord"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// DNSRecordQuery is the builder for querying DNSRecord entities.
type DNSRecordQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.DNSRecord
	// eager-loading edges.
	withDNSRecordToTag *TagQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DNSRecordQuery builder.
func (drq *DNSRecordQuery) Where(ps ...predicate.DNSRecord) *DNSRecordQuery {
	drq.predicates = append(drq.predicates, ps...)
	return drq
}

// Limit adds a limit step to the query.
func (drq *DNSRecordQuery) Limit(limit int) *DNSRecordQuery {
	drq.limit = &limit
	return drq
}

// Offset adds an offset step to the query.
func (drq *DNSRecordQuery) Offset(offset int) *DNSRecordQuery {
	drq.offset = &offset
	return drq
}

// Order adds an order step to the query.
func (drq *DNSRecordQuery) Order(o ...OrderFunc) *DNSRecordQuery {
	drq.order = append(drq.order, o...)
	return drq
}

// QueryDNSRecordToTag chains the current query on the "DNSRecordToTag" edge.
func (drq *DNSRecordQuery) QueryDNSRecordToTag() *TagQuery {
	query := &TagQuery{config: drq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsrecord.Table, dnsrecord.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dnsrecord.DNSRecordToTagTable, dnsrecord.DNSRecordToTagColumn),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DNSRecord entity from the query.
// Returns a *NotFoundError when no DNSRecord was found.
func (drq *DNSRecordQuery) First(ctx context.Context) (*DNSRecord, error) {
	nodes, err := drq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dnsrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (drq *DNSRecordQuery) FirstX(ctx context.Context) *DNSRecord {
	node, err := drq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DNSRecord ID from the query.
// Returns a *NotFoundError when no DNSRecord ID was found.
func (drq *DNSRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dnsrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (drq *DNSRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := drq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DNSRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DNSRecord entity is not found.
// Returns a *NotFoundError when no DNSRecord entities are found.
func (drq *DNSRecordQuery) Only(ctx context.Context) (*DNSRecord, error) {
	nodes, err := drq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dnsrecord.Label}
	default:
		return nil, &NotSingularError{dnsrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (drq *DNSRecordQuery) OnlyX(ctx context.Context) *DNSRecord {
	node, err := drq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DNSRecord ID in the query.
// Returns a *NotSingularError when exactly one DNSRecord ID is not found.
// Returns a *NotFoundError when no entities are found.
func (drq *DNSRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = &NotSingularError{dnsrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (drq *DNSRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := drq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DNSRecords.
func (drq *DNSRecordQuery) All(ctx context.Context) ([]*DNSRecord, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return drq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (drq *DNSRecordQuery) AllX(ctx context.Context) []*DNSRecord {
	nodes, err := drq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DNSRecord IDs.
func (drq *DNSRecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := drq.Select(dnsrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (drq *DNSRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := drq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (drq *DNSRecordQuery) Count(ctx context.Context) (int, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return drq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (drq *DNSRecordQuery) CountX(ctx context.Context) int {
	count, err := drq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (drq *DNSRecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return drq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (drq *DNSRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := drq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DNSRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (drq *DNSRecordQuery) Clone() *DNSRecordQuery {
	if drq == nil {
		return nil
	}
	return &DNSRecordQuery{
		config:             drq.config,
		limit:              drq.limit,
		offset:             drq.offset,
		order:              append([]OrderFunc{}, drq.order...),
		predicates:         append([]predicate.DNSRecord{}, drq.predicates...),
		withDNSRecordToTag: drq.withDNSRecordToTag.Clone(),
		// clone intermediate query.
		sql:  drq.sql.Clone(),
		path: drq.path,
	}
}

// WithDNSRecordToTag tells the query-builder to eager-load the nodes that are connected to
// the "DNSRecordToTag" edge. The optional arguments are used to configure the query builder of the edge.
func (drq *DNSRecordQuery) WithDNSRecordToTag(opts ...func(*TagQuery)) *DNSRecordQuery {
	query := &TagQuery{config: drq.config}
	for _, opt := range opts {
		opt(query)
	}
	drq.withDNSRecordToTag = query
	return drq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty" hcl:"name,attr"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DNSRecord.Query().
//		GroupBy(dnsrecord.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (drq *DNSRecordQuery) GroupBy(field string, fields ...string) *DNSRecordGroupBy {
	group := &DNSRecordGroupBy{config: drq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return drq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty" hcl:"name,attr"`
//	}
//
//	client.DNSRecord.Query().
//		Select(dnsrecord.FieldName).
//		Scan(ctx, &v)
//
func (drq *DNSRecordQuery) Select(field string, fields ...string) *DNSRecordSelect {
	drq.fields = append([]string{field}, fields...)
	return &DNSRecordSelect{DNSRecordQuery: drq}
}

func (drq *DNSRecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range drq.fields {
		if !dnsrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if drq.path != nil {
		prev, err := drq.path(ctx)
		if err != nil {
			return err
		}
		drq.sql = prev
	}
	return nil
}

func (drq *DNSRecordQuery) sqlAll(ctx context.Context) ([]*DNSRecord, error) {
	var (
		nodes       = []*DNSRecord{}
		withFKs     = drq.withFKs
		_spec       = drq.querySpec()
		loadedTypes = [1]bool{
			drq.withDNSRecordToTag != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dnsrecord.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DNSRecord{config: drq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, drq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := drq.withDNSRecordToTag; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*DNSRecord)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.DNSRecordToTag = []*Tag{}
		}
		query.withFKs = true
		query.Where(predicate.Tag(func(s *sql.Selector) {
			s.Where(sql.InValues(dnsrecord.DNSRecordToTagColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.dns_record_dns_record_to_tag
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "dns_record_dns_record_to_tag" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "dns_record_dns_record_to_tag" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.DNSRecordToTag = append(node.Edges.DNSRecordToTag, n)
		}
	}

	return nodes, nil
}

func (drq *DNSRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := drq.querySpec()
	return sqlgraph.CountNodes(ctx, drq.driver, _spec)
}

func (drq *DNSRecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := drq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (drq *DNSRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsrecord.Table,
			Columns: dnsrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dnsrecord.FieldID,
			},
		},
		From:   drq.sql,
		Unique: true,
	}
	if fields := drq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dnsrecord.FieldID)
		for i := range fields {
			if fields[i] != dnsrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := drq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := drq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := drq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := drq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, dnsrecord.ValidColumn)
			}
		}
	}
	return _spec
}

func (drq *DNSRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(drq.driver.Dialect())
	t1 := builder.Table(dnsrecord.Table)
	selector := builder.Select(t1.Columns(dnsrecord.Columns...)...).From(t1)
	if drq.sql != nil {
		selector = drq.sql
		selector.Select(selector.Columns(dnsrecord.Columns...)...)
	}
	for _, p := range drq.predicates {
		p(selector)
	}
	for _, p := range drq.order {
		p(selector, dnsrecord.ValidColumn)
	}
	if offset := drq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := drq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DNSRecordGroupBy is the group-by builder for DNSRecord entities.
type DNSRecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (drgb *DNSRecordGroupBy) Aggregate(fns ...AggregateFunc) *DNSRecordGroupBy {
	drgb.fns = append(drgb.fns, fns...)
	return drgb
}

// Scan applies the group-by query and scans the result into the given value.
func (drgb *DNSRecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := drgb.path(ctx)
	if err != nil {
		return err
	}
	drgb.sql = query
	return drgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := drgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSRecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := drgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) StringX(ctx context.Context) string {
	v, err := drgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSRecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := drgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) IntX(ctx context.Context) int {
	v, err := drgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSRecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := drgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := drgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DNSRecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := drgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DNSRecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drgb *DNSRecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := drgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drgb *DNSRecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range drgb.fields {
		if !dnsrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := drgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drgb *DNSRecordGroupBy) sqlQuery() *sql.Selector {
	selector := drgb.sql
	columns := make([]string, 0, len(drgb.fields)+len(drgb.fns))
	columns = append(columns, drgb.fields...)
	for _, fn := range drgb.fns {
		columns = append(columns, fn(selector, dnsrecord.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(drgb.fields...)
}

// DNSRecordSelect is the builder for selecting fields of DNSRecord entities.
type DNSRecordSelect struct {
	*DNSRecordQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (drs *DNSRecordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := drs.prepareQuery(ctx); err != nil {
		return err
	}
	drs.sql = drs.DNSRecordQuery.sqlQuery(ctx)
	return drs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drs *DNSRecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := drs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSRecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drs *DNSRecordSelect) StringsX(ctx context.Context) []string {
	v, err := drs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drs *DNSRecordSelect) StringX(ctx context.Context) string {
	v, err := drs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSRecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drs *DNSRecordSelect) IntsX(ctx context.Context) []int {
	v, err := drs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drs *DNSRecordSelect) IntX(ctx context.Context) int {
	v, err := drs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSRecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drs *DNSRecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := drs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drs *DNSRecordSelect) Float64X(ctx context.Context) float64 {
	v, err := drs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DNSRecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drs *DNSRecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := drs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (drs *DNSRecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsrecord.Label}
	default:
		err = fmt.Errorf("ent: DNSRecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drs *DNSRecordSelect) BoolX(ctx context.Context) bool {
	v, err := drs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drs *DNSRecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := drs.sqlQuery().Query()
	if err := drs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drs *DNSRecordSelect) sqlQuery() sql.Querier {
	selector := drs.sql
	selector.Select(selector.Columns(drs.fields...)...)
	return selector
}
