package embeddingiface_test

import (
	"articles/tests/mocks/embeddingiface"
	"strconv"
	"testing"
)

// embeding interface in a struct for huge interface problem,
// or when you need to mock only one or fews, but not all funcs
type tMock struct {
	embeddingiface.Test
}

func (t *tMock) T1(i int) string {
	return strconv.Itoa(i)
}

func Test(t *testing.T) {
	tm := &tMock{}
	if v := embeddingiface.TestFn(tm); v != nil {
		t.Log("pass", v)
	} else {
		t.Error("failed")
	}
}
