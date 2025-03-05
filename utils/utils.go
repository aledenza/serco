package utils

func GetOptionalValue[T any](values ...T) (ret T) {
	if len(values) != 0 {
		return values[0]
	}
	return
}

func GetOptionalValueOrNil[T any](values ...T) *T {
	if len(values) != 0 {
		return &values[0]
	}
	return nil
}

func Ptr[T any](value T) *T {
	return &value
}
