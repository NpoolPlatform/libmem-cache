package cruder

import (
	"context"
)

type Any interface{}

const (
	EQ   = 1
	GT   = 2
	LT   = 3
	LIKE = 4
)

type Cond struct {
	Op  uint32
	Val Any
}

type Cruder interface {
	Create(ctx context.Context, in Any) (Any, error)
	CreateBulk(ctx context.Context, in []Any) ([]Any, error)

	Update(ctx context.Context, in Any) (Any, error)
	UpdateFields(ctx context.Context, id string, fields map[string]Any) (Any, error)

	AtomicInc(ctx context.Context, id string, fields []string) (Any, error)
	AtomicSub(ctx context.Context, id string, fields []string) (Any, error)
	AtomicSet(ctx context.Context, id string, fields map[string]Any) (Any, error)

	Row(ctx context.Context, id string) (Any, error)
	Rows(ctx context.Context, conds map[string]Cond, offset, limit uint32) ([]Any, error)

	Exist(ctx context.Context, conds map[string]Cond) (bool, error)
	Count(ctx context.Context, conds map[string]Cond) (uint32, error)
	// TODO: MAP | REDUCE | FILTER | SUM

	Delete(ctx context.Context, id string) (Any, error)
}
