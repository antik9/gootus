package max

import (
	"errors"
	"reflect"
)

type lessFunc func(int, int) bool

func FindMax(_slice interface{}, less lessFunc) (interface{}, error) {
	if _slice == nil {
		return nil, errors.New("nil is given")
	} else {
		slice := reflect.ValueOf(_slice)
		if reflect.TypeOf(_slice).Kind() != reflect.Slice || slice.Len() == 0 {
			return nil, errors.New("Provide not empty slice")
		} else {
			maxElementIndex := 0
			for idx := 1; idx < slice.Len(); idx++ {
				if less(maxElementIndex, idx) {
					maxElementIndex = idx
				}
			}
			return slice.Index(maxElementIndex).Interface(), nil
		}
	}
}
