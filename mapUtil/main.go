package mapUtil

/**
 * @description: 获取map所有字段
 * @param {map[string]interface{}} data
 * @return {*}
 */
func Fields(data map[string]string) ([]string, error) {

	var result = make([]string, 0)
	for key := range data {
		result = append(result, key)
	}
	return result, nil
}
