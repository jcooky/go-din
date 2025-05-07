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
	res := fmt.Sprintf("%s/%s", t.PkgPath(), t.Name())

	return Name(res)
}

func NewRandomName() Name {
	return Name(uuid.NewString())
}
