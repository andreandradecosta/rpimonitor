package models

type Info map[string]interface{}

func getData(data interface{}, err error) interface{} {
	if err != nil {
		return err.Error()
	}
	return data
}
