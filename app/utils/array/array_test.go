package array

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	data := []interface{}{
		1,
	}
	println(reflect.TypeOf(data[0]).Kind())
}
