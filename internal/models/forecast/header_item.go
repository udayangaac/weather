package forecast

type HeaderItem string

func (h HeaderItem) Length() int {
	return len(h) + 2
}
func (h HeaderItem) String() string {
	return string(h)
}
