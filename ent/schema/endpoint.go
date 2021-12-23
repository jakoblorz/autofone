package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Endpoint holds the schema definition for the Endpoint entity.
type Endpoint struct {
	ent.Schema
}

// Fields of the Settings.
func (Endpoint) Fields() []ent.Field {
	return []ent.Field{
		field.String("host"),
		field.Uint16("port").Optional(),
		field.Uint16("packet_whitelist").Optional(),
		field.Uint16("packet_blacklist").Optional(),
	}
}

// Edges of the Settings.
func (Endpoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requestedBy", Settings.Type),
	}
}
