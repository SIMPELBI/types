package types

import "context"

type Schemas[I any, O any] interface {
	GetById(context.Context, I) (O, error)
	Get(context.Context) ([]O, error)
	Update(context.Context, I) error
	Delete(context.Context, I) error
}
