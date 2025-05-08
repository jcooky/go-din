package din

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
)

type (
	Name string
)

func (n Name) String() string {
	return string(n)
}

func NewTypeName[T any]() Name {
	t := reflect.TypeFor[T]()

	var name string
	for {
		if t.Kind() == reflect.Ptr {
			name += "*"
			t = t.Elem()
		} else if t.Kind() == reflect.Slice {
			name += "[]"
			t = t.Elem()
		} else if t.Kind() == reflect.Map {
			name += fmt.Sprintf("map[%s]", t.Key().String())
			t = t.Elem()
		} else {
			break
		}
	}

	if pkgPath := t.PkgPath(); pkgPath != "" {
		name = pkgPath + "/" + name
	}

	if t.Kind() == reflect.Interface && t.Name() == "" {
		panic("Last Elem type should be not any")
	} else {
		name += t.Name()
	}

	return Name(name)
}

func NewRandomName() Name {
	return Name(uuid.NewString())
}
