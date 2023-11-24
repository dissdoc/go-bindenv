package internal

import "reflect"

func ScanEnv(cfg interface{}) {
	elem := reflect.ValueOf(cfg).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		kind := field.Type().Kind()

		switch kind {
		case reflect.Struct:
			// Создается ссылка на структуру
			addr := field.Addr()
			ScanEnv(addr.Interface())
		case reflect.Slice:
			result := make([]interface{}, field.Len())
			for i := 0; i < field.Len(); i++ {
				addr := field.Index(i).Addr()
				result[i] = addr.Interface()
				ScanEnv(result[i])
			}
		default:
			readField(&field)
		}
	}
}

func readField(field *reflect.Value) {
	if !field.IsValid() {
		return
	}

	value := field.Interface()
	if reflect.TypeOf(value).Kind() != reflect.String {
		return
	}

	if GetRule().MatchString(value.(string)) {
		param := bindParam(value.(string))
		field.Set(reflect.ValueOf(param))
	}
}

func bindParam(field string) string {
	res := GetRule().FindAllStringSubmatch(field, -1)
	return MustParam(res[0][1])
}
