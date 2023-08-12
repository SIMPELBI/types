package types

type Schemas[I any, O any] interface {
	GetById(id I) (O, error)
	Get() (O, error)
	Update(data I) error
	Delete(data I) error
}
