/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"fmt"
	"reflect"
)

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

type CompareMode int

const (
	CompareModeLess         CompareMode = -1
	CompareModeLessEqual    CompareMode = -10
	CompareModeEqual        CompareMode = 0
	CompareModeGreater      CompareMode = 1
	CompareModeGreaterEqual CompareMode = 10
)

type Serializer interface {
	HashCode() int
}

func Compare[T any](i, j T, CompareMode CompareMode) bool {
	iValue := reflect.ValueOf(i)
	jValue := reflect.ValueOf(j)

	switch iValue.Kind() {
	case reflect.Pointer:
		if iValue.Elem().Kind() == reflect.Struct {
			switch ip := iValue.Interface().(type) {
			case Serializer:
				switch jp := jValue.Interface().(type) {
				case Serializer:
					return compare(ip.HashCode(), jp.HashCode(), CompareMode)
				}
			}
			switch ip := iValue.Interface().(type) {
			case fmt.Stringer:
				switch jp := jValue.Interface().(type) {
				case fmt.Stringer:
					return compare(ip.String(), jp.String(), CompareMode)
				}
			}
		}
		return Compare(iValue.Elem().Interface(), jValue.Elem().Interface(), CompareMode)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(iValue.Int(), jValue.Int(), CompareMode)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return compare(iValue.Uint(), jValue.Uint(), CompareMode)

	case reflect.Float32, reflect.Float64:
		return compare(iValue.Float(), jValue.Float(), CompareMode)

	case reflect.Complex128, reflect.Complex64:
		return Complex128Cmp(iValue.Complex(), jValue.Complex(), CompareMode)

	case reflect.String:
		return compare(iValue.String(), jValue.String(), CompareMode)

	case reflect.Bool:
		if CompareMode == CompareModeEqual {
			return iValue.Bool() == jValue.Bool()
		}

	default:
		return false
	}

	return false
}

func compare[T Comparable](i, j T, CompareMode CompareMode) bool {
	switch CompareMode {
	case CompareModeLess:
		return i < j
	case CompareModeLessEqual:
		return i <= j
	case CompareModeEqual:
		return i == j
	case CompareModeGreater:
		return i > j
	case CompareModeGreaterEqual:
		return i >= j
	}
	return false
}

func Complex64Cmp(i, j complex64, CompareMode CompareMode) bool {
	switch CompareMode {
	case CompareModeLess:
		return real(i) < real(j) || (real(i) == real(j) && imag(i) < imag(j))
	case CompareModeEqual:
		return real(i) == real(j) && imag(i) == imag(j)
	case CompareModeGreater:
		return real(i) > real(j) || (real(i) == real(j) && imag(i) > imag(j))
	}
	return false
}

func Complex128Cmp(i, j complex128, CompareMode CompareMode) bool {
	switch CompareMode {
	case CompareModeLess:
		return real(i) < real(j) || (real(i) == real(j) && imag(i) < imag(j))
	case CompareModeLessEqual:
		return real(i) <= real(j) || (real(i) == real(j) && imag(i) <= imag(j))
	case CompareModeEqual:
		return real(i) == real(j) && imag(i) == imag(j)
	case CompareModeGreater:
		return real(i) > real(j) || (real(i) == real(j) && imag(i) > imag(j))
	case CompareModeGreaterEqual:
		return real(i) >= real(j) || (real(i) == real(j) && imag(i) >= imag(j))
	}
	return false
}
