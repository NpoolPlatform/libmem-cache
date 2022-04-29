package cruder

import (
	"context"
	"fmt"
	"reflect"

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

func AnyTypeInt32(v Any) (int32, error) {
	switch v.(type) {
	case int32:
		return v.(int32), nil
	case int:
		return int32(v.(int)), nil
	}
	return -1, fmt.Errorf("invalid value type: %v (int32)", reflect.TypeOf(v))
}

func AnyTypeUint32(v Any) (uint32, error) {
	switch v.(type) {
	case uint32:
		return v.(uint32), nil
	case int32:
		return uint32(v.(int32)), nil
	case int:
		return uint32(v.(int)), nil
	}
	return 0, fmt.Errorf("invalid value type: %v (uint32)", reflect.TypeOf(v))
}

func AnyTypeUUID(v Any) (uuid.UUID, error) {
	if _, ok := v.(uuid.UUID); ok {
		return v.(uuid.UUID), nil
	}
	if _, ok := v.(string); !ok {
		return uuid.UUID{}, fmt.Errorf("invalid value type: %v (uuid)", reflect.TypeOf(v))
	}
	return uuid.Parse(v.(string))
}
