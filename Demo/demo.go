package Demo

//reference
func RemoveEmptyStringNoReturn(s *[]string) {
	res := make([]string, 0)
	for _, item := range *s {
		if item != "" {
			res = append(res, item)
		}
	}
	*s = res
}

func RemoveEmptyStringNoReturn0(s []string) { // nosence error methor
	res := make([]string, 0)
	for _, item := range s {
		if item != "" {
			res = append(res, item)
		}
	}
	s = res
}

func RemoveEmptyStringWithReturn(s []string) (res []string) {
	res = make([]string, 0)
	for _, item := range s {
		if item != "" {
			res = append(res, item)
		}
	}
	return res
}

func DistincStringList(s []string) []string {
	exitDict := make(map[string]bool)
	res := make([]string, 0)
	for _, item := range s {
		if _, ok := exitDict[item]; !ok {
			res = append(res, item)
			exitDict[item] = true
		}
	}
	return res
}
