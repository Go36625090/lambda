/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package lambda

import (
	"errors"
	"go/types"
	"testing"
)

type User struct {
	Id   int
	Name string
}

func (u *User) HashCode() int {
	return u.Id
}

func TestStream(t *testing.T) {

}

type Lottery struct {
	Code int `json:"code"`
	Data struct {
		OpenCode string `json:"openCode"`
		Expect   string `json:"expect"`
	} `json:"data"`
}
type Balls struct {
	Red  []int `json:"red"`
	Blue int   `json:"blue"`
}

type Error struct {
	Msg string
}

func (a Error) Error() string {
	return "error"
}
func TestArgumentErrorUnwrapping(t *testing.T) {
	var err error = &types.ArgumentError{
		Index: 1,
		Err:   Error{Msg: "test"},
	}
	//e := &Error{
	//Msg: "err",
	//}
	var e interface{} = new(chan error)
	if !errors.As(err, e) {
		t.Logf("error %v does not wrap types.Error", err)
	}
	//{
	//}
	//if e.Msg != "test" {
	//	t.Errorf("e.Msg = %q, want %q", e.Msg, "test")
	//}
}
