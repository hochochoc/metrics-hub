// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"metric-hub/ent/report"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Report is the model entity for the Report schema.
type Report struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ServerName holds the value of the "server_name" field.
	ServerName string `json:"server_name,omitempty"`
	// StartTime holds the value of the "start_time" field.
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime holds the value of the "end_time" field.
	EndTime time.Time `json:"end_time,omitempty"`
	// DurationInSeconds holds the value of the "duration_in_seconds" field.
	DurationInSeconds float64 `json:"duration_in_seconds,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Report) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case report.FieldDurationInSeconds:
			values[i] = new(sql.NullFloat64)
		case report.FieldID:
			values[i] = new(sql.NullInt64)
		case report.FieldServerName:
			values[i] = new(sql.NullString)
		case report.FieldStartTime, report.FieldEndTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Report", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Report fields.
func (r *Report) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case report.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case report.FieldServerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field server_name", values[i])
			} else if value.Valid {
				r.ServerName = value.String
			}
		case report.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				r.StartTime = value.Time
			}
		case report.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				r.EndTime = value.Time
			}
		case report.FieldDurationInSeconds:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_in_seconds", values[i])
			} else if value.Valid {
				r.DurationInSeconds = value.Float64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Report.
// Note that you need to call Report.Unwrap() before calling this method if this Report
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Report) Update() *ReportUpdateOne {
	return (&ReportClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Report entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Report) Unwrap() *Report {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Report is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Report) String() string {
	var builder strings.Builder
	builder.WriteString("Report(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", server_name=")
	builder.WriteString(r.ServerName)
	builder.WriteString(", start_time=")
	builder.WriteString(r.StartTime.Format(time.ANSIC))
	builder.WriteString(", end_time=")
	builder.WriteString(r.EndTime.Format(time.ANSIC))
	builder.WriteString(", duration_in_seconds=")
	builder.WriteString(fmt.Sprintf("%v", r.DurationInSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// Reports is a parsable slice of Report.
type Reports []*Report

func (r Reports) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
