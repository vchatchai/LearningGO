package reflection

import (
	"fmt"
	"reflect"
)

type Member struct {
	Age int `something:"age"`
}

func TagValue() {
	member := Member{34}
	t := reflect.TypeOf(member)
	v := reflect.ValueOf(&member)
	field := t.Field(0)
	//field, _ := t.FieldByName("Age") //alternative
	fmt.Println(field.Tag.Get("something"))

	fmt.Println(v.Elem().Field(0).Interface())
}
