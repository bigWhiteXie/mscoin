package tools

import (
	"fmt"
	"reflect"
)

func Default(data any) error {
	typeOf := reflect.TypeOf(data)
	if typeOf.Kind() != reflect.Pointer {
		return fmt.Errorf("只有指针类型才能默认值")
	}
	v := reflect.ValueOf(data).Elem() // 获取结构体值的反射对象
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// 如果字段是零值，则初始化为默认值
		if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			zeroValue := reflect.New(field.Type()).Elem() // 获取字段类型的零值
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				zeroValue.SetInt(0)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				zeroValue.SetUint(0)
			case reflect.Float32, reflect.Float64:
				zeroValue.SetFloat(0)
			case reflect.String:
				zeroValue.SetString("")
			case reflect.Bool:
				zeroValue.SetBool(false)
			// 如果是其他类型，则无法初始化，跳过
			default:
				continue
			}
			field.Set(zeroValue) // 设置字段的值为零值
		}
	}
	return nil
}

func defaultString() reflect.Value {
	var i = ""
	return reflect.ValueOf(i)
}

func defaultInt() reflect.Value {
	var i int = -1
	return reflect.ValueOf(i)
}

func defaultInt32() reflect.Value {
	var i int32 = -1
	return reflect.ValueOf(i)
}
func defaultInt64() reflect.Value {
	var i int64 = -1
	return reflect.ValueOf(i)
}

func defaultFloat64() reflect.Value {
	var i float64 = -1
	return reflect.ValueOf(i)
}
func defaultFloat32() reflect.Value {
	var i float32 = -1
	return reflect.ValueOf(i)
}
