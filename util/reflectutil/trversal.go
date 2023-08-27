package reflectx

import (
	"fmt"
	"reflect"
)

type TypeInfo struct {
	Path        string
	ElementType string
	IsSlice     bool
	IsStruct    bool
	IsPtr       bool
	IsLeaf      bool
}

type TypeInfoVisitor interface {
	Visit(ti TypeInfo)
}

func TraverseType(path TypeInfo, typ reflect.Type, v TypeInfoVisitor) {

	// t.Logf("val kind of path %s is %v", path, val.Kind())
	switch typ.Kind() {
	case reflect.Struct:
		//typ := val.Type()

		path.ElementType = typ.String()
		path.IsStruct = true
		v.Visit(path)

		for i := 0; i < typ.NumField(); i++ {

			sformat := "%s.%s"
			isSlice := false
			isPtr := false

			switch typ.Field(i).Type.Kind() {
			case reflect.Slice:
				isSlice = true
				sformat = "%s.[]%s"
				if typ.Field(i).Type.Elem().Kind() == reflect.Ptr {
					isPtr = true
					sformat = "%s.[]*%s"
				}

			case reflect.Ptr:
				isPtr = true
				sformat = "%s.*%s"
			}
			fieldPath := TypeInfo{Path: fmt.Sprintf(sformat, path.Path, typ.Field(i).Name), IsSlice: isSlice, IsPtr: isPtr}
			TraverseType(fieldPath, typ.Field(i).Type, v)
		}
	case reflect.Ptr:
		/*		ti := TypeInfo{TargetPath: path, IsPtr: true, ElementType: typ.Elem().String()}
				v.Visit(ti)*/

		TraverseType(path /*fmt.Sprintf("(*%s)", path)*/, typ.Elem(), v)

	case reflect.Slice:
		/*		ti := TypeInfo{TargetPath: path, IsSlice: true, ElementType: typ.Elem().String()}
				v.Visit(ti)*/
		TraverseType(path /*fmt.Sprintf("([]%s)", path) */, typ.Elem(), v)
	default:
		path.ElementType = typ.String()
		path.IsLeaf = true
		v.Visit(path)
	}

}
