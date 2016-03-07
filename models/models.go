package models

//Info is a basic data type
type Info map[string]interface{}

func getData(data interface{}, err error) interface{} {
	if err != nil {
		return err.Error()
	}
	return data
}
