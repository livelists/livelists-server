package shared

func FalseIfNil(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}
