// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: copyfrom.go

package measurements

import (
	"context"
)

// iteratorForCreateMeasurement implements pgx.CopyFromSource.
type iteratorForCreateMeasurement struct {
	rows                 []CreateMeasurementParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateMeasurement) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateMeasurement) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ComponentID,
		r.rows[0].DateMeasured,
		r.rows[0].PetID,
		r.rows[0].Amount,
	}, nil
}

func (r iteratorForCreateMeasurement) Err() error {
	return nil
}

func (q *Queries) CreateMeasurement(ctx context.Context, arg []CreateMeasurementParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"panel_measurements"}, []string{"component_id", "date_measured", "pet_id", "amount"}, &iteratorForCreateMeasurement{rows: arg})
}
