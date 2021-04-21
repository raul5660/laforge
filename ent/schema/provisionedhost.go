package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProvisionedHost holds the schema definition for the ProvisionedHost entity.
type ProvisionedHost struct {
	ent.Schema
}

// Fields of the ProvisionedHost.
func (ProvisionedHost) Fields() []ent.Field {
	return []ent.Field{
		field.String("subnet_ip"),
	}
}

// Edges of the ProvisionedHost.
func (ProvisionedHost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ProvisionedHostToTag", Tag.Type),
		edge.To("ProvisionedHostToStatus", Status.Type),
		edge.To("ProvisionedHostToProvisionedNetwork", ProvisionedNetwork.Type),
		edge.To("ProvisionedHostToHost", Host.Type),
		edge.From("ProvisionedHostToProvisioningStep", ProvisioningStep.Type).Ref("ProvisioningStepToProvisionedHost"),
		edge.From("ProvisionedHostToAgentStatus", AgentStatus.Type).Ref("AgentStatusToProvisionedHost"),
		edge.From("ProvisionedHostToPlan", Plan.Type).Ref("PlanToProvisionedHost"),
	}
}
