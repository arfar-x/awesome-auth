package password

type Password string

func Check(value1 Password, value2 Password) bool {
	return true
}

func Make(value Password) string {
	return string(value)
}
