package mapping

import "encoding/json"

func MapToStruct[T any](data interface{}, value *T) error {
	rawData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawData, value)
	return err
}
