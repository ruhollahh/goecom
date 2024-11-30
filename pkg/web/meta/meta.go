package httpmeta

import (
	"context"
	"time"
)

type ctxKey int

const key ctxKey = 1

// Meta represent state for each request.
type Meta struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// SetMeta sets the specified Meta in the context.
func Set(ctx context.Context, v *Meta) context.Context {
	return context.WithValue(ctx, key, v)
}

// GetMeta returns the values from the context.
func Get(ctx context.Context) *Meta {
	v, ok := ctx.Value(key).(*Meta)
	if !ok {
		return &Meta{
			TraceID: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}

	return v
}

// GetTraceID returns the trace id from the context.
func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(key).(*Meta)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.TraceID
}
