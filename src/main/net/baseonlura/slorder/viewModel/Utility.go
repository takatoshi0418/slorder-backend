package viewModel

import "encoding/json"

var DATE_FORMAT = "2006-01-02"

func DeepCopy(src interface{}, dst interface{}) error {
	byteData, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &dst)
	if err != nil {
		return err
	}
	return nil
}
