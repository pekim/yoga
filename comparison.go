package yoga

import (
	"math"
)

func IsUndefined[T float32 | float64](value T) bool {
	return value != value
}

func IsDefined[T float32 | float64](value T) bool {
	return !IsUndefined(value)
}

func maxOrDefined[T float32 | float64](a, b T) T {
	if IsDefined(a) && IsDefined(b) {
		return max(a, b)
	}
	return If[T](IsUndefined(a), b, a)
}

func minOrDefined[T float32 | float64](a, b T) T {
	if IsDefined(a) && IsDefined(b) {
		return min(a, b)
	}
	return If[T](IsUndefined(a), b, a)
}

// InexactEquals
func inexactEqual[T float32 | float64](a, b T) bool {
	if IsDefined(a) && IsDefined(b) {
		return math.Abs(float64(a-b)) < 0.0001
	}
	return IsUndefined(a) && IsUndefined(b)
}

func inexactEquals(a, b any) bool {
	switch a := a.(type) {
	case float32:
		switch b := b.(type) {
		case float32:
			return inexactEqual(a, b)
		}
	case float64:
		switch b := b.(type) {
		case float64:
			return inexactEqual(a, b)
		}
	case YGValue:
		switch b := b.(type) {
		case YGValue:
			return a.equal(b)
		}
	case CompactValue:
		switch b := b.(type) {
		case CompactValue:
			return inexactEquals(a.YGValue(), b.YGValue())
		}
	case FloatOptional:
		switch b := b.(type) {
		case FloatOptional:
			return inexactEqual(a.unwrap(), b.unwrap())
		}
	case []CompactValue:
		switch b := b.(type) {
		case []CompactValue:
			for i := 0; i < len(a); i++ {
				if !inexactEquals(a[i].YGValue(), b[i].YGValue()) {
					return false
				}
			}
		}
		return true
	}
	return false
}
