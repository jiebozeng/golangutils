package intAndByte

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"runtime/debug"
	"strconv"
	"time"
)

//uint16 to bytes
func Uint16ToBytes(n uint16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

//int32 to bytes
func Int32ToBytes(n int32) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//bytes to int32
func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

//value interfase{}
func ToInt32(value interface{}) int32 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case int32:
		return value.(int32)
	case int64:
		return int32(value.(int64))
	case int:
		return int32(value.(int))
	case float64:
		return int32(value.(float64))
	case []uint8:
		v1, err := strconv.ParseInt(string(value.([]uint8)), 10, 64)
		if err != nil {
			return 0
		}
		return int32(v1)
	case string:
		if value.(string) == "" {
			return 0
		}
		ret, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			debug.PrintStack()
		}
		return int32(ret)
	}
	return 0
}

//value interfase{}
func ToInt64(value interface{}) int64 {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case int32:
		return int64(value.(int32))
	case int64:
		return value.(int64)
	case int:
		return int64(value.(int))
	case float64:
		return int64(value.(float64))
	case []uint8:
		v1, err := strconv.ParseInt(string(value.([]uint8)), 10, 64)
		if err != nil {
			return 0
		}
		return int64(v1)
	case string:
		if value.(string) == "" {
			return 0
		}
		ret, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			debug.PrintStack()
		}
		return ret
	}
	debug.PrintStack()
	return 0
}

//value interfase{}
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch value.(type) {
	case int32:
		return strconv.Itoa(int(value.(int32)))
	case int:
		return strconv.Itoa(int(value.(int)))
	case uint16:
		return strconv.Itoa(int(value.(uint16)))
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case []uint8:
		return string(value.([]uint8))
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case time.Time:
		return value.(time.Time).Format("2006-01-02 15:04:05")
	case string:
		return value.(string)
	default:
		debug.PrintStack()
		ret, err := json.Marshal(value)
		if err != nil {
		}
		return string(ret)
	}
	return ""
}

//value interfase{}
func ToFloat64(value interface{}) float64 {
	switch value.(type) { //多选语句switch
	case int32:
		return float64(value.(int32))
	case int:
		return float64(value.(int))
	case string:
		ret, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {

			debug.PrintStack()
		}
		return ret
	}
	debug.PrintStack()
	return 0
}
