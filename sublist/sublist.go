package sublist

import "reflect"

// Relation
type Relation string

/*
	whether x and y are “deeply equal” or not
*/

func Sublist(a, b []int) Relation{
	switch {
	case reflect.DeepEqual(a, b) :
		return "equal"
	case Check(b, a) :
		return "sublist"
	case Check(a, b):
		return "superlist"
	}
	return "unequal"
}

func Check (a, b []int) (ok bool) {
	for i:=0; i <= len(a) - len(b); i++ {
		if reflect.DeepEqual(a[i:i+len(b)], b) {
			return true
		}
	}
	return
}