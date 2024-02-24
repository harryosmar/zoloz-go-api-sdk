package enums

import "fmt"

func listStringerToMap[T fmt.Stringer](list []T) map[string]any {
	if list == nil || len(list) == 0 {
		return nil
	}
	m := map[string]any{}
	for _, v := range list {
		m["name"] = v.String()
	}
	return m
}

func listStringerToListString[T fmt.Stringer](list []T) []string {
	if list == nil || len(list) == 0 {
		return nil
	}
	var l []string
	for _, v := range list {
		l = append(l, v.String())
	}
	return l
}
