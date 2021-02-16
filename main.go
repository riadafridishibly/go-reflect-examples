
package main

import (
	"fmt"
	"reflect"
)

type index int

type Results struct {
	Values *[]interface{}
	// ErrorWrapper errorwrapper.Wrapper
}

// SetAt sets data to Results
func (r *Results) SetAt(i index, holder interface{}) {
	if r.Values == nil {
		values := make([]interface{}, i+1)
		r.Values = &values
	}
	(*r.Values)[i] = holder
}

// GetAt set data to Results
func (r *Results) GetAt(i index, holder interface{}) {
	srcReflectValue := reflect.Indirect(reflect.ValueOf((*r.Values)[i]))
	dstReflectValue := reflect.Indirect(reflect.ValueOf(holder))

	if srcReflectValue.Kind() != dstReflectValue.Kind() {
		// type mismatch, abort
		return
	}

	switch srcReflectValue.Kind() {
	case reflect.Struct:
        dstReflectValue.Set(srcReflectValue)
		// for i := 0; i < srcReflectValue.NumField(); i++ {
		// 	dstReflectValue.Set(srcReflectValue)
		// }
	}
}

type newType struct {
	hasStr bool
	str    string
}

func main() {

	var newData newType
	var oldData newType

	var res = Results{}

	// modify oldData
	oldData.hasStr = true
	oldData.str = "Hello World"


	res.SetAt(0, &oldData)
	res.GetAt(0, &newData)

	fmt.Println(oldData, newData)
}
