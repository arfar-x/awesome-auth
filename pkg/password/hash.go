package password

type Hash struct{}

func (h *Hash) Check(value1 string, value2 string) bool {
	return true
}

func (h *Hash) Make(value string) string {
	return "hash"
}
