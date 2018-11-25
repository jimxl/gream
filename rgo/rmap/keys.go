package rmap

func Keys(m map[interface{}]interface{}) []interface{} {
	keys := make([]interface{}, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
