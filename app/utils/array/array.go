package array

import (
	"errors"
	"reflect"
	"strings"
)

func Unique(data interface{}) []interface{} {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Slice {
		panic(errors.New("unique data type must slice"))
	}
	var rel []interface{}
	tmp := make(map[interface{}]bool)

	// TODO 有冗余重复代码
	switch typeInfo.String() {
	case "[]uint64":
		originData := data.([]uint64)
		for _, v := range originData {
			if _, ok := tmp[v]; !ok {
				tmp[v] = true
				rel = append(rel, v)
			}
		}
		break
	case "[]int64":
		originData := data.([]int64)
		for _, v := range originData {
			if _, ok := tmp[v]; !ok {
				tmp[v] = true
				rel = append(rel, v)
			}
		}
		break
	case "[]uint":
		originData := data.([]uint)
		for _, v := range originData {
			if _, ok := tmp[v]; !ok {
				tmp[v] = true
				rel = append(rel, v)
			}
		}
		break
	case "[]int":
		originData := data.([]int)
		for _, v := range originData {
			if _, ok := tmp[v]; !ok {
				tmp[v] = true
				rel = append(rel, v)
			}
		}
		break
	case "[]string":
		originData := data.([]string)
		for _, v := range originData {
			if _, ok := tmp[v]; !ok {
				tmp[v] = true
				rel = append(rel, v)
			}
		}
		break
	}

	return rel
}

func In(val string, arr []string) bool {
	for _, v := range arr {
		if strings.Contains(v, val) {
			return true
		}
	}

	return false
}


