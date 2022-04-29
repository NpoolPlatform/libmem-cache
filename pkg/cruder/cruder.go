package cruder

import (
	"context"
	"github.com/google/uuid"
	"reflect"
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

func AnyTypeInt32(v Any) (int32, error) {
	switch v.(type) {
	case int32:
		return v.(int32), nil
	case int:
		return int32(v.(int)), nil
	}
	return -1, fmt.Errorf("invalid value type: %v", reflect.TypeOf(v))
}
