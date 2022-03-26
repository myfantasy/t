package t

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T constraints.Ordered](a T, b T) T {
	if a < b {
		return a
	}

	return b
}

func Greatest[T constraints.Ordered](a T, b ...T) T {
	for _, c := range b {
		if c > a {
			a = c
		}
	}

	return a
}

func Least[T constraints.Ordered](a T, b ...T) T {
	for _, c := range b {
		if c < a {
			a = c
		}
	}

	return a
}

func In[T constraints.Ordered](a T, b ...T) bool {
	for _, c := range b {
		if c == a {
			return true
		}
	}

	return false
}

func Iif[T any](cond bool, ifT T, ifF T) T {
	if cond {
		return ifT
	}

	return ifF
}

func Choose[K Number, T any](idx K, a ...T) T {
	return a[int(idx)]
}

func Coalesce[T any](a ...*T) *T {
	for _, c := range a {
		if c != nil {
			return c
		}
	}
	return nil
}

// ToP - to Pointer
func ToP[T any](a T) *T {
	return &a
}

// FromP - from Pointer
func FromP[T any](a *T) T {
	var r T
	if a != nil {
		return *a
	}

	return r
}
