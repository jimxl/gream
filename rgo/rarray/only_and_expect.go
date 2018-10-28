package rarray

func OnlyAndExpect(only, expect, all []string) (results map[string]bool) {
	results = map[string]bool{}
	defualtValue := (len(only) == 0)
	for _, item := range all {
		results[item] = defualtValue
	}

	for _, item := range only {
		if _, ok := results[item]; ok {
			results[item] = true
		}
	}

	for _, item := range expect {
		if _, ok := results[item]; ok {
			results[item] = false
		}
	}
	return
}
