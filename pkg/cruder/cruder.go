package cruder

import (
	"context"
	"github.com/google/uuid"
)

type Any interface{}

const (
	EQ   = "eq"
	GT   = "gt"
	LT   = "lt"
	LIKE = "like"
)

type Cond struct {
	Op  string
	Val Any
}

type Cruder interface {
	Create(ctx context.Context, in Any) (Any, error)
	CreateBulk(ctx context.Context, in []Any) ([]Any, error)

	Update(ctx context.Context, in Any) (Any, error)
	UpdateFields(ctx context.Context, id uuid.UUID, fields map[string]Any) (Any, error)
	AtomicUpdateFields(ctx context.Context, id uuid.UUID, fields map[string]Any) (Any, error)

	Row(ctx context.Context, id uuid.UUID) (Any, error)
	Rows(ctx context.Context, conds map[string]Cond, offset, limit uint32) ([]Any, error)

	Exist(ctx context.Context, conds map[string]Cond) (bool, error)
	Count(ctx context.Context, conds map[string]Cond) (uint32, error)
	// TODO: MAP | REDUCE | FILTER | SUM

	Delete(ctx context.Context, id uuid.UUID) (Any, error)
}