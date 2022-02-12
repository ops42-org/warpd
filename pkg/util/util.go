package util

func StrPtr(s string) *string {
	return &s
}

func DefaultStrPtr(s *string, def string) *string {
	if s != nil {
		return s
	}
	return StrPtr(def)
}
