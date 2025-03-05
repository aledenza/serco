package serco

import "encoding/json"

func Dump[O any](input any) (output O, err error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return output, err
	}
	err = json.Unmarshal(bytes, &output)
	return output, err
}

func DumpMany[O any, I any](input []I) ([]O, error) {
	output := make([]O, len(input))
	for i, inp := range input {
		if out, err := Dump[O](inp); err != nil {
			return nil, err
		} else {
			output[i] = out
		}
	}
	return output, nil
}
