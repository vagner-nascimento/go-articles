package embeddingiface

type Test interface {
	T1(int) string
	T2(string) int
}

func TestFn(t Test) any {
	return t.T1(1)
}
