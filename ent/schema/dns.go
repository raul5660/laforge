package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// DNS holds the schema definition for the DNS entity.
type DNS struct {
	ent.Schema
}

// Fields of the DNS.
func (DNS) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			StructTag(`hcl:"type,attr"`),
		field.String("root_domain").
			StructTag(`hcl:"root_domain,attr" `),
		field.JSON("dns_servers", []string{}).
			StructTag(`hcl:"dns_servers,attr"`),
		field.JSON("ntp_servers", []string{}).
			StructTag(`hcl:"ntp_servers,optional"`),
		field.JSON("config", map[string]string{}).
			StructTag(`hcl:"config,optional"`),
	}
}

// Edges of the DNS.
func (DNS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("DNSToTag", Tag.Type),
	}
}
