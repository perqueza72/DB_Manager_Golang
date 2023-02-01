package filehandlers

import (
	"fmt"
	"reflect"
	"strings"
)

func ExistProperty(item interface{}, property string) bool {

	metaValue := reflect.ValueOf(item).Elem()
	field := metaValue.FieldByName(property)
	if field == (reflect.Value{}) {
		return false
	}

	return true
}

func ExistJsonProperty(item interface{}, fieldName string) (bool, map[string]int, error) {
	v := reflect.ValueOf(item).Elem()
	if !v.CanAddr() {
		return false, nil, fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	findJsonName := func(t reflect.StructTag) (string, error) {
		if jt, ok := t.Lookup("json"); ok {
			return strings.Split(jt, ",")[0], nil
		}
		return "", fmt.Errorf("tag provided does not define a json tag, %v", fieldName)
	}
	fieldNames := map[string]int{}
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		jname, _ := findJsonName(tag)
		fieldNames[jname] = i
	}

	_, ok := fieldNames[fieldName]

	return ok, fieldNames, nil
}

func SetField(item interface{}, fieldName string, value interface{}) error {
	ok, fieldNames, err := ExistJsonProperty(item, fieldName)

	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("field %s does not exist within the provided item", fieldName)
	}

	v := reflect.ValueOf(item).Elem()
	fieldNum, ok := fieldNames[fieldName]
	fieldVal := v.Field(fieldNum)
	fieldVal.Set(reflect.ValueOf(value))
	return nil
}
