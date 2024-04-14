package logs

func LogConfigsParams(prefix string, settings map[string]interface{}) {
	for key, value := range settings {
		switch value := value.(type) {
		case map[string]interface{}:
			// Если значение является картой, рекурсивно логируем её содержимое
			if prefix == "" {
				LogConfigsParams(key, value)
			} else {
				LogConfigsParams(prefix+"."+key, value)
			}
		default:
			// Логирование простых значений
			if prefix == "" {
				Log.Infof("%s: %v", key, value)
			} else {
				Log.Infof("%s.%s: %v", prefix, key, value)
			}
		}
	}
}
