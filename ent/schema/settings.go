package schema

import "entgo.io/ent"

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return nil
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return nil
}
