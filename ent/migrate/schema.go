// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ReportsColumns holds the columns for the "reports" table.
	ReportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "server_name", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "duration_in_seconds", Type: field.TypeFloat64},
	}
	// ReportsTable holds the schema information for the "reports" table.
	ReportsTable = &schema.Table{
		Name:       "reports",
		Columns:    ReportsColumns,
		PrimaryKey: []*schema.Column{ReportsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ReportsTable,
	}
)

func init() {
}
