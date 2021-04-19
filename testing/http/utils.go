package http

import (
	"encoding/json"
	"fmt"
)

func getBodyString(arg interface{}) (value string, err error) {

	switch body := arg.(type) {
	case string:
		value = string(body)
		break

	case []byte:
		value = string(body)
		break

	case uint8, uint16, uint32, uint64:
	case int8, int16, int32, int64:
	case float32, float64:
		value = fmt.Sprintf("%d", body)
		break

	case bool:
		value = fmt.Sprintf("%t", body)
		break

	default:
		res, errVal := json.Marshal(body)
		value, err = string(res), errVal
		break
	}

	return
}

func getPrettyBodyString(arg interface{}) (value string, err error) {

	switch body := arg.(type) {
	case string:
		value = string(body)
		break

	case []byte:
		value = string(body)
		break

	case uint8, uint16, uint32, uint64:
	case int8, int16, int32, int64:
	case float32, float64:
		value = fmt.Sprintf("%d", body)
		break

	case bool:
		value = fmt.Sprintf("%t", body)
		break

	default:
		res, errVal := json.MarshalIndent(body, "", "  ")
		value, err = string(res), errVal
		break
	}

	return
}
