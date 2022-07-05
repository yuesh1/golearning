package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 9, 3, 4, 5, 6}
	//for i := 0; i < len(arr); i++ {
	//	t.Log(arr[i])
	//}
	//for idx, e := range arr {
	//	t.Log(idx, e)
	//}

	for _, e := range arr {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr_sec3 := arr[:3]
	t.Log(arr_sec3)

	arr_3 := arr[3:]
	t.Log(arr_3)

	arr_all := arr[:]
	t.Log(arr_all)

	// 不支持负数
	//arr_3 := arr[:-1]
	//t.Log(arr_all)
}
