package graph

import (
	null "github.com/volatiletech/null/v8"
)

func NullInt64ToIntPtr(v null.Int64) *int {
	if !v.Valid {
		return nil
	}
	vNew := int(v.Int64)
	return &vNew
}
