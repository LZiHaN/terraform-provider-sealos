// Copyright (c) eden.zh.li@outlook.com, Inc.
// SPDX-License-Identifier: MPL-2.0

package conversions

func ToStringSlice(value interface{}) []string {
	if value == nil {
		return nil
	}
	slice, ok := value.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = v.(string)
	}
	return result
}

func GetStringValue(data map[string]interface{}, key string) string {
	value, exists := data[key]
	if exists {
		return value.(string)
	}
	return ""
}

func GetUInt16Value(data map[string]interface{}, key string) uint16 {
	value, exists := data[key]
	if exists {
		return uint16(value.(int))
	}
	return 0
}

func GetBoolValue(data map[string]interface{}, key string) bool {
	value, exists := data[key]
	if exists {
		return value.(bool)
	}
	return false
}
