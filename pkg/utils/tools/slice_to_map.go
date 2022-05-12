package tools

func StringSliceToMap(str []string) map[string]string {
	m := make(map[string]string, len(str))
	for _, s := range str {
		m[s] = ""
	}
	return m
}
