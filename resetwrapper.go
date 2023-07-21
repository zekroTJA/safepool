package safepool

// ResetWrapper
type ResetWrapper[T any] struct {
	Inner     T
	resetHook func(v T)
}

var _ (ResetState) = (*ResetWrapper[any])(nil)

func Wrap[T any](v T, reset func(v T)) *ResetWrapper[T] {
	return &ResetWrapper[T]{
		Inner:     v,
		resetHook: reset,
	}
}

func (t ResetWrapper[T]) ResetState() {
	t.resetHook(t.Inner)
}
