package gesa

// String は s string のポインタを返す
func String(s string) *string {
	return &s
}

// StringValue は s *string の値を返す
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Bool は b bool のポインタを返す
func Bool(b bool) *bool {
	return &b
}

// BoolValue は b *bool の値を返す
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// Int は i int のポインタを返す
func Int(i int) *int {
	return &i
}

// IntValue は i *int の値を返す
func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
