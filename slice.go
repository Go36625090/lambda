/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"encoding/json"
	"reflect"
	"sort"
)

type Slice[T any] []T

func (m *Slice[T]) Values() any {
	if len(*m) == 0 {
		return nil
	}

	var slice = reflect.SliceOf(reflect.TypeOf((*m)[0]))
	o := reflect.MakeSlice(slice, 0, 0)
	var values []reflect.Value
	for _, t := range *m {
		values = append(values, reflect.ValueOf(t))
	}
	return reflect.Append(o, values...).Interface()
}

func (m *Slice[T]) Assignment(o any) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, o)
}

func (m *Slice[T]) Filter(f func(i T) bool) *Slice[T] {
	var out Slice[T]
	for _, t := range *m {
		if f(t) {
			out = append(out, t)
		}
	}
	return &out
}

func (m *Slice[T]) Clear() {
	*m = Slice[T]{}
}

func (m *Slice[T]) Contains(i T) bool {
	if len(*m) == 0 {
		return false
	}

	for _, t := range *m {
		if Compare(i, t, CompareModeEqual) {
			return true
		}
	}
	return false
}

func (m *Slice[T]) ContainsAll(i *Slice[T]) bool {
	if len(*m) == 0 || len(*i) == 0 {
		return false
	}

	for _, v := range *i {
		if !m.Contains(v) {
			return false
		}
	}
	return true
}

func (m *Slice[T]) IndexOf(val T) int {
	if len(*m) == 0 {
		return -1
	}

	for i, t := range *m {
		if Compare(t, val, CompareModeEqual) {
			return i
		}
	}
	return -1
}

func (m *Slice[T]) Append(values ...T) {
	if values == nil || len(values) == 0 {
		return
	}
	*m = append(*m, values...)
}

func (m *Slice[T]) Insert(values ...T) {
	if values == nil || len(values) == 0 {
		return
	}
	*m = append(values, *m...)
}

func (m *Slice[T]) Sort(cmp func(i, j T) bool) {
	s := &slice[T]{data: m}
	s.cmp = cmp
	sort.Sort(s)
	copy(*m, *s.data)
}

func (m *Slice[T]) Length() int {
	return len(*m)
}

func (m *Slice[T]) Sorted() {
	s := &slice[T]{data: m}
	sort.Sort(s)
	copy(*m, *s.data)
}

func (m *Slice[T]) Reverse() {
	s := &slice[T]{data: m}
	sort.Sort(sort.Reverse(s))
	copy(*m, *s.data)
}

func (m *Slice[T]) Remove(values ...T) int {
	if 0 == len(values) {
		return 0
	}
	slice := Stream(values).Slice()
	var result Slice[T]
	for _, e := range *m {
		if slice.Contains(e) {
			continue
		}
		result.Append(e)
	}
	*m = result
	return 0
}

// RetailAll 求交集
func (m *Slice[T]) RetailAll(i *Slice[T]) *Slice[T] {
	if 0 == len(*i) {
		return m
	}
	if len(*m) == 0 {
		return &Slice[T]{}
	}

	var tmp Slice[T]
	for _, v := range *m {
		if i.Contains(v) {
			tmp.Append(v)
		}
	}
	if len(tmp) == 0 {
		return &tmp
	}
	var result Slice[T]
	for _, v := range *i {
		if tmp.Contains(v) && !result.Contains(v) {
			result.Append(v)
		}
	}
	return &result
}

// RemoveAll 求差集
func (m *Slice[T]) RemoveAll(i *Slice[T]) *Slice[T] {
	if len(*m) == 0 {
		return &Slice[T]{}
	}
	if len(*i) == 0 {
		return m
	}

	var result Slice[T]
	for _, v := range *m {
		if !i.Contains(v) {
			result.Append(v)
		}
	}
	for _, v := range *i {
		if !m.Contains(v) {
			result.Append(v)
		}
	}
	return &result
}

// UnionAll 求并集
func (m *Slice[T]) UnionAll(i *Slice[T]) *Slice[T] {
	if 0 == len(*i) {
		return m
	}
	if len(*m) == 0 {
		return i
	}

	sl := make(Slice[T], len(*m)+len(*i))
	copy(sl, *m)
	copy(sl[len(*m):], *i)
	var result Slice[T]
	for _, v := range sl {
		if !result.Contains(v) {
			result.Append(v)
		}
	}
	return &result
}

func (m *Slice[T]) Walk(f func(i T)) {
	for _, t := range *m {
		f(t)
	}
}

type slice[T any] struct {
	data *Slice[T]
	cmp  func(i, j T) bool
}

func (s *slice[T]) Len() int {
	return len(*s.data)
}

func (s *slice[T]) Less(i, j int) bool {
	if s.cmp == nil {
		return Compare(&(*s.data)[i], &(*s.data)[j], CompareModeLess)
	}
	return s.cmp((*s.data)[i], (*s.data)[j])
}

func (s *slice[T]) Swap(i, j int) {
	(*s.data)[i], (*s.data)[j] = (*s.data)[j], (*s.data)[i]
}
