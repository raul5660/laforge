// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/dnsrecord"
)

// DNSRecord is the model entity for the DNSRecord schema.
type DNSRecord struct {
	config `hcl:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" hcl:"name,attr"`
	// Values holds the value of the "values" field.
	Values []string `json:"values,omitempty" hcl:"values,optional"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty" hcl:"type,attr"`
	// Zone holds the value of the "zone" field.
	Zone string `json:"zone,omitempty" hcl:"zone,attr" `
	// Vars holds the value of the "vars" field.
	Vars map[string]string `json:"vars,omitempty" hcl:"vars,optional"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty" hcl:"disabled,optional"`
	// Tags holds the value of the "tags" field.
	Tags map[string]string `json:"tags,omitempty" hcl:"tags,attr"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DNSRecordQuery when eager-loading is set.
	Edges                                             DNSRecordEdges `json:"edges"`
	provisioning_step_provisioning_step_to_dns_record *int
}

// DNSRecordEdges holds the relations/edges for other nodes in the graph.
type DNSRecordEdges struct {
	// DNSRecordToTag holds the value of the DNSRecordToTag edge.
	DNSRecordToTag []*Tag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DNSRecordToTagOrErr returns the DNSRecordToTag value or an error if the edge
// was not loaded in eager-loading.
func (e DNSRecordEdges) DNSRecordToTagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.DNSRecordToTag, nil
	}
	return nil, &NotLoadedError{edge: "DNSRecordToTag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DNSRecord) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&[]byte{},         // values
		&sql.NullString{}, // type
		&sql.NullString{}, // zone
		&[]byte{},         // vars
		&sql.NullBool{},   // disabled
		&[]byte{},         // tags
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*DNSRecord) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // provisioning_step_provisioning_step_to_dns_record
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DNSRecord fields.
func (dr *DNSRecord) assignValues(values ...interface{}) error {
	if m, n := len(values), len(dnsrecord.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	dr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		dr.Name = value.String
	}

	if value, ok := values[1].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field values", values[1])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &dr.Values); err != nil {
			return fmt.Errorf("unmarshal field values: %v", err)
		}
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[2])
	} else if value.Valid {
		dr.Type = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field zone", values[3])
	} else if value.Valid {
		dr.Zone = value.String
	}

	if value, ok := values[4].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field vars", values[4])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &dr.Vars); err != nil {
			return fmt.Errorf("unmarshal field vars: %v", err)
		}
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field disabled", values[5])
	} else if value.Valid {
		dr.Disabled = value.Bool
	}

	if value, ok := values[6].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field tags", values[6])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &dr.Tags); err != nil {
			return fmt.Errorf("unmarshal field tags: %v", err)
		}
	}
	values = values[7:]
	if len(values) == len(dnsrecord.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field provisioning_step_provisioning_step_to_dns_record", value)
		} else if value.Valid {
			dr.provisioning_step_provisioning_step_to_dns_record = new(int)
			*dr.provisioning_step_provisioning_step_to_dns_record = int(value.Int64)
		}
	}
	return nil
}

// QueryDNSRecordToTag queries the DNSRecordToTag edge of the DNSRecord.
func (dr *DNSRecord) QueryDNSRecordToTag() *TagQuery {
	return (&DNSRecordClient{config: dr.config}).QueryDNSRecordToTag(dr)
}

// Update returns a builder for updating this DNSRecord.
// Note that, you need to call DNSRecord.Unwrap() before calling this method, if this DNSRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (dr *DNSRecord) Update() *DNSRecordUpdateOne {
	return (&DNSRecordClient{config: dr.config}).UpdateOne(dr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (dr *DNSRecord) Unwrap() *DNSRecord {
	tx, ok := dr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DNSRecord is not a transactional entity")
	}
	dr.config.driver = tx.drv
	return dr
}

// String implements the fmt.Stringer.
func (dr *DNSRecord) String() string {
	var builder strings.Builder
	builder.WriteString("DNSRecord(")
	builder.WriteString(fmt.Sprintf("id=%v", dr.ID))
	builder.WriteString(", name=")
	builder.WriteString(dr.Name)
	builder.WriteString(", values=")
	builder.WriteString(fmt.Sprintf("%v", dr.Values))
	builder.WriteString(", type=")
	builder.WriteString(dr.Type)
	builder.WriteString(", zone=")
	builder.WriteString(dr.Zone)
	builder.WriteString(", vars=")
	builder.WriteString(fmt.Sprintf("%v", dr.Vars))
	builder.WriteString(", disabled=")
	builder.WriteString(fmt.Sprintf("%v", dr.Disabled))
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", dr.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// DNSRecords is a parsable slice of DNSRecord.
type DNSRecords []*DNSRecord

func (dr DNSRecords) config(cfg config) {
	for _i := range dr {
		dr[_i].config = cfg
	}
}
