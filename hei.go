package size

func Size(a int) string {
	switch {
	case a < 0:
		return "negative"
	}
	return "enormous"
}
