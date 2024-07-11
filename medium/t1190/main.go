package main

func reverseParentheses(s string) string {
	teleports := getHoles(s)
	i := 0
	step := 1
	res := make([]byte, 0, len(s))
	for cnt := 0; cnt < len(s); cnt++ {
		if s[i] == '(' || s[i] == ')' {
			i = teleports[i]
			step = -step
		} else {
			res = append(res, s[i])
		}
		i += step
	}
	return string(res)
}

func getHoles(s string) map[int]int {
	holes := make(map[int]int, len(s))
	stack := stack{mem: make([]int, 0, len(s))}
	for i, c := range s {
		if c == '(' {
			stack.push(i)
		}
		if c == ')' {
			x := stack.pop()
			holes[x] = i
			holes[i] = x
		}
	}
	return holes
}

type stack struct {
	mem []int
}

func (s *stack) push(x int) {
	s.mem = append(s.mem, x)
}

func (s *stack) pop() int {
	x := s.mem[len(s.mem)-1]
	s.mem = s.mem[:len(s.mem)-1]
	return x
}
