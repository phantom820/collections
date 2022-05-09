package wrapper

type Integer int

func (i Integer) HashCode() int {
	return int(i)
}

func (i Integer) Equals(other Integer) bool {
	return i == other
}

type String string

func (s String) HashCode() int {
	const (
		m = 1e9 + 9
		p = 53
	)
	p_pow := 1
	runes := []rune(s)
	code := 0
	for _, r := range runes {
		code = (code + p_pow*(int(r)+1)) % m
		p_pow = (p_pow * p) % m
	}
	return code
}

func (s String) Equals(other String) bool {
	return s == other
}

func (s String) Less(other String) bool {
	return s < other
}

func (x Integer) Less(y Integer) bool {
	return x < y
}
