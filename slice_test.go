/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"fmt"
	"testing"
)

type personal struct {
	age  int
	name string
}

func (p *personal) String() string {
	return fmt.Sprintf("%d-%s", p.age, p.name)
}

func (p *personal) HashCode() int {
	return p.age
}

func TestSlice(t *testing.T) {
	slice := Stream[int64]([]int64{1, 1, 2, 5}).Slice()
	a1 := Stream[int64]([]int64{1, 2, 3, 4}).Slice()
	a2 := a1.Remove(*slice...)
	t.Log(a2, a1)
}
