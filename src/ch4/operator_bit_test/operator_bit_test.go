package operator_bit_test

import "testing"

func TestOperatorBit(t *testing.T) {
	a := 7
	t.Log(a &^ 1)
	t.Log(a &^ 0)
	//	按位清零
}
