package utilx

import "strings"

func FirstUpper(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

func FirstLower(s string) string {
	if len(s) > 0 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return s
}

func RemoveDuplicateElement(slices []string) []string {
	result := make([]string, 0, len(slices))
	temp := map[string]struct{}{}
	for _, item := range slices {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
