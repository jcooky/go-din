package din_test

import (
	"github.com/jcooky/go-din"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTypeNameByAnyThenShouldBeError(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("Expected panic, but got nil")
		}
	}()
	name := din.NewTypeName[any]()
	t.Fatalf("Expected panic, but got name: %s", name)
}

func TestNewTypeNameByTypeThanShouldBeOK(t *testing.T) {
	type TestStruct struct{}
	name := din.NewTypeName[TestStruct]()
	require.Equal(t, "github.com/jcooky/go-din_test/TestStruct", name.String())
}

func TestNewTypaNameByPtrTypeThenShouldBeOK(t *testing.T) {
	type TestStruct struct{}
	name := din.NewTypeName[*TestStruct]()
	require.Equal(t, "github.com/jcooky/go-din_test/*TestStruct", name.String())
}

func TestNewTypeNameByMapTypeThenShouldBeOK(t *testing.T) {
	type TestStruct struct{}
	name := din.NewTypeName[map[string]TestStruct]()
	require.Equal(t, "github.com/jcooky/go-din_test/map[string]TestStruct", name.String())
}

func TestNewTypeNameBySliceTypeThenShouldBeOK(t *testing.T) {
	type TestStruct struct{}
	name := din.NewTypeName[[]TestStruct]()
	require.Equal(t, "github.com/jcooky/go-din_test/[]TestStruct", name.String())
}

func TestNewTypeNameBySlicePtrTypeThenShouldBeOK(t *testing.T) {
	type TestStruct struct{}
	name := din.NewTypeName[[]*TestStruct]()
	require.Equal(t, "github.com/jcooky/go-din_test/[]*TestStruct", name.String())
}

func TestNewTypeNameByMapPtrTypeThenShouldBeOK(t *testing.T) {
	type TestStruct struct{}
	name := din.NewTypeName[map[string]*TestStruct]()
	require.Equal(t, "github.com/jcooky/go-din_test/map[string]*TestStruct", name.String())
}

func TestNewTypeNameByMapSliceAnyTypeThenShouldBeError(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("Expected panic, but got nil")
		}
	}()
	name := din.NewTypeName[[]any]()
	t.Fatalf("Expected panic, but got name: %s", name)
}
