package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.String("server_name"),
		field.Time("start_time"),
		field.Time("end_time"),
		field.Float("duration_in_seconds"),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return nil
}
