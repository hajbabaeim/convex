package helpers

import "encoding/json"

func ConvertJson[T any](a any, b T) (*T, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func ConvertSliceJson[T any](a any, result ...T) []T {
	data, err := json.Marshal(a)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil
	}
	return result
}

func GetPointer[T any](a T) *T {
	return &a
}
