package filehandlers

import (
	"bufio"
	models "models_zinc"
	"os"
	"reflect"
)

func GetLines(path *string) (*[]*string, error) {
	readFile, err := os.Open(*path)

	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	response := []*string{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		response = append(response, &text)
	}
	readFile.Close()

	return &response, nil
}

func ParseStrings2Email(strings *[]*string) (*models.Email, error) {

	email := models.Email{}

	lastProperty := false

	LAST_TYPE := "X-FileName"
	last_type_setted := false
	prev_string := ""
	actual_string := ""
	property := ""

	v := reflect.ValueOf(&email).Elem()

	for _, line := range *strings {
		if lastProperty && !last_type_setted {
			SetField(&email, LAST_TYPE, actual_string)

			actual_string = ""
			last_type_setted = true
		}

		if !lastProperty {
			prev_string += actual_string
			actual_string = ""
		}
		for _, c := range *line {
			if c == ':' {

				exist, _, _ := ExistJsonProperty(&email, actual_string)
				if exist {

					if property != "" {

						if actual_string == LAST_TYPE {
							lastProperty = true
						}

						if exist {
							SetField(&email, property, prev_string)
						} else {
							v.FieldByName(property)
							v.SetString(prev_string)
						}
					}

					property = actual_string
					actual_string = ""
					prev_string = ""
				} else {
					actual_string += string(c)
				}

				continue
			}
			actual_string += string(c)
		}

		if lastProperty {
			actual_string += "\n"
		}
	}

	if actual_string != "" {
		email.Content = actual_string
	}

	return &email, nil
}
