package logging

func convertToBool(val interface{}) bool {
	switch val.(type) {
	case bool:
		return val.(bool)

	case float64:
		if val.(float64) != 0 {
			return true
		}

	case int:
		if val.(int) != 0 {
			return true
		}

	case string:
		if val.(string) != "" {
			return true
		}

	default:
		// For any other type, assume they wanted debug mode, since the key/item was provided
		return true
	}

	return false
}
