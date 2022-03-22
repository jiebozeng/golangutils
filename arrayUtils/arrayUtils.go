package arrayUtils

import "reflect"
// add slice 
func ItemExists(targetArray interface{}, item interface{}) bool {
	arr := reflect.ValueOf(targetArray)
	if arr.Kind() != reflect.Array || arr.Kind() != reflect.Slice{
		panic("targetArray Invalid data-type")
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func ArrayDeleteUint32(arr []uint32, tar uint32) ([]uint32, bool) {
	del := false
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == tar {
			del = true
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	return arr, del
}

func ArrayDeleteInt32(arr []int32, tar int32) ([]int32, bool) {
	del := false
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == tar {
			del = true
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	return arr, del
}

func ArrayDeleteInt64(arr []int64, tar int64) ([]int64, bool) {
	del := false
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == tar {
			del = true
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	return arr, del
}

func ArrayDeleteString(arr []string, tar string) ([]string, bool) {
	del := false
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == tar {
			del = true
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	return arr, del
}

/**
 指定元素在数组中存在个数
 return Array Exist item numbers
 */
func ArrayExistsCount(arrayOrSlice interface{}, target interface{}) int32 {
	count := int32(0)
	targetValue := reflect.ValueOf(arrayOrSlice)
	switch reflect.TypeOf(arrayOrSlice).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			i2 := targetValue.Index(i).Interface()
			if i2 == target {
				count++
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(target)).IsValid() {
			return 1
		}
	}
	return count
}
