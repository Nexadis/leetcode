package main

func isValid(s string) bool {
	st := newStack(len(s))
	for _, c := range s {
		c := byte(c)
		switch c {
		case '(', '{', '[':
			st.push(c)
		case ')':
			cl := st.pop()
			if cl != '(' {
				return false
			}
		case '}':
			cl := st.pop()
			if cl != '{' {
				return false
			}
		case ']':
			cl := st.pop()
			if cl != '[' {
				return false
			}
		}
	}
	return st.size == 0
}

func newStack(size int) stack {
	return stack{
		size: 0,
		mem:  make([]byte, 0, size),
	}
}

type stack struct {
	mem  []byte
	size int
}

func (s *stack) push(b byte) {
	s.size++
	s.mem = append(s.mem, b)
}

func (s *stack) pop() byte {
	if s.size == 0 {
		return 0
	}
	r := s.mem[s.size-1]
	s.size--
	s.mem = s.mem[:s.size]
	return r
}
