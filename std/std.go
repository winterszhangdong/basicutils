package std

// DelDuplicate 删除字符串列表中的重复项目
func DelDuplicate(list []string) ([]string, bool) {
	result := make([]string, 0, len(list))
	var flag bool
	tmp := map[string]struct{}{}
	for _, item := range list {
		if _, ok := tmp[item]; !ok {
			tmp[item] = struct{}{}
			result = append(result, item)
		} else {
			flag = true
		}
	}

	return result, flag
}
