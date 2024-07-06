package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	i := len(a) - 1
	j := len(b) - 1

	c := make([]byte, 0, len(a)+1)
	add := byte(0)
	for j >= 0 {
		x, y := a[i]-'0', b[j]-'0'
		c = append(c, (x+y+add)%2+'0')
		add = (x + y + add) / 2
		i--
		j--

	}
	for i >= 0 {
		x := a[i] - '0'
		n := (x + add) % 2
		c = append(c, n+'0')
		add = (x + add) / 2
		i--
	}

	if add == 1 {
		c = append(c, '1')
	}

	l := 0
	r := len(c) - 1
	for l < r {
		c[l], c[r] = c[r], c[l]
		l++
		r--
	}

	return string(c)
}
