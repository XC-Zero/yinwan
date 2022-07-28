package math_plus

import "testing"

func TestName(t *testing.T) {
	_, err := NewFromString("15.5")
	if err != nil {
		panic(err)
	}
	_, err = NewFromString("15")
	if err != nil {
		panic(err)
	}
	_, err = NewFromString("15/5")
	if err != nil {
		panic(err)
	}
	//_, err = NewFromString("15/0.3")
	//if err != nil {
	//	panic(err)
	//}
}
