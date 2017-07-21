package main

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	ss := *s
	l := len(ss)
	ret := ss[l-1]
	*s = ss[:l-1]
	return ret
}

func (s *Stack) Top() interface{} {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
