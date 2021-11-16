package arrayUtils

import "reflect"

func itemExists(targetArray interface{}, item interface{}) bool {
	arr := reflect.ValueOf(targetArray)
	if arr.Kind() != reflect.Array {
		panic("targetArray Invalid data-type")
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
