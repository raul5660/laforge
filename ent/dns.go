// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/dns"
)

// DNS is the model entity for the DNS schema.
type DNS struct {
	config `hcl:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty" hcl:"type,attr"`
	// RootDomain holds the value of the "root_domain" field.
	RootDomain string `json:"root_domain,omitempty" hcl:"root_domain,attr" `
	// DNSServers holds the value of the "dns_servers" field.
	DNSServers []string `json:"dns_servers,omitempty" hcl:"dns_servers,attr"`
	// NtpServers holds the value of the "ntp_servers" field.
	NtpServers []string `json:"ntp_servers,omitempty" hcl:"ntp_servers,optional"`
	// Config holds the value of the "config" field.
	Config map[string]string `json:"config,omitempty" hcl:"config,optional"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DNSQuery when eager-loading is set.
	Edges                          DNSEdges `json:"edges"`
	competition_competition_to_dns *int
}

// DNSEdges holds the relations/edges for other nodes in the graph.
type DNSEdges struct {
	// DNSToTag holds the value of the DNSToTag edge.
	DNSToTag []*Tag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DNSToTagOrErr returns the DNSToTag value or an error if the edge
// was not loaded in eager-loading.
func (e DNSEdges) DNSToTagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.DNSToTag, nil
	}
	return nil, &NotLoadedError{edge: "DNSToTag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DNS) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // type
		&sql.NullString{}, // root_domain
		&[]byte{},         // dns_servers
		&[]byte{},         // ntp_servers
		&[]byte{},         // config
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*DNS) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // competition_competition_to_dns
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DNS fields.
func (d *DNS) assignValues(values ...interface{}) error {
	if m, n := len(values), len(dns.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[0])
	} else if value.Valid {
		d.Type = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field root_domain", values[1])
	} else if value.Valid {
		d.RootDomain = value.String
	}

	if value, ok := values[2].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field dns_servers", values[2])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &d.DNSServers); err != nil {
			return fmt.Errorf("unmarshal field dns_servers: %v", err)
		}
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field ntp_servers", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &d.NtpServers); err != nil {
			return fmt.Errorf("unmarshal field ntp_servers: %v", err)
		}
	}

	if value, ok := values[4].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field config", values[4])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &d.Config); err != nil {
			return fmt.Errorf("unmarshal field config: %v", err)
		}
	}
	values = values[5:]
	if len(values) == len(dns.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field competition_competition_to_dns", value)
		} else if value.Valid {
			d.competition_competition_to_dns = new(int)
			*d.competition_competition_to_dns = int(value.Int64)
		}
	}
	return nil
}

// QueryDNSToTag queries the DNSToTag edge of the DNS.
func (d *DNS) QueryDNSToTag() *TagQuery {
	return (&DNSClient{config: d.config}).QueryDNSToTag(d)
}

// Update returns a builder for updating this DNS.
// Note that, you need to call DNS.Unwrap() before calling this method, if this DNS
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *DNS) Update() *DNSUpdateOne {
	return (&DNSClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *DNS) Unwrap() *DNS {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: DNS is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *DNS) String() string {
	var builder strings.Builder
	builder.WriteString("DNS(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", type=")
	builder.WriteString(d.Type)
	builder.WriteString(", root_domain=")
	builder.WriteString(d.RootDomain)
	builder.WriteString(", dns_servers=")
	builder.WriteString(fmt.Sprintf("%v", d.DNSServers))
	builder.WriteString(", ntp_servers=")
	builder.WriteString(fmt.Sprintf("%v", d.NtpServers))
	builder.WriteString(", config=")
	builder.WriteString(fmt.Sprintf("%v", d.Config))
	builder.WriteByte(')')
	return builder.String()
}

// DNSs is a parsable slice of DNS.
type DNSs []*DNS

func (d DNSs) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
