package resources

import "reflect"

type BaseResourceInterface interface {
	toMap() []interface{}
}

type BaseResource struct{}

func (resource *BaseResource) toMap() map[string]interface{} {
	val := reflect.ValueOf(resource)

	if val.Kind() != reflect.Struct {
		panic("Given data is not a struct.")
	}

	numFields := val.NumField()
	keyValues := make(map[string]interface{})

	for i := 0; i < numFields; i++ {
		field := val.Type().Field(i)
		key := field.Name
		value := val.Field(i).String()
		keyValues[key] = value
	}

	return keyValues
}
