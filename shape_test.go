package arrgo

import "testing"

func TestHstack(t *testing.T) {
	var a = Arange(10)
	var b = Arange(10)
	t.Log(Hstack(a, b, a, b))
}

func TestConcat(t *testing.T) {
	var a = Arange(1, 11).Reshape(2, 5)
	var b = Arange(10).Reshape(2, 5)
	t.Log(Concat(0, a, b))
	t.Log(Concat(1, a, b))
}

func TestAtLeast2D(t *testing.T) {
	a := Arange(10)
	AtLeast2D(a)
	if !SameIntSlice(a.shape, []int{1, 10}) {
		t.Error("Expected [1, 10], got ", a.shape)
	}

	a.Reshape(1, 1, 10)
	AtLeast2D(a)
	if !SameIntSlice(a.shape, []int{1, 1, 10}) {
		t.Error("Expected [1, 1, 10], got ", a.shape)
	}

	if AtLeast2D(nil) != nil {
		t.Error("Expected nil, got ", AtLeast2D(nil))
	}
}
