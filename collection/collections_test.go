package collection

import "testing"

func TestMapCopy(t *testing.T) {

	map1 := make(map[string]string)
	map2 := map1
	map3 := map2

	map1["a"] = "1"
	map2["b"] = "2"
	map3["c"] = "3"

	t.Log(map1)
	t.Log(map2)
	t.Log(map3)

}

func TestArrayCopy(t *testing.T) {

	array1 := make([]string, 0)
	array2 := array1
	array3 := array2

	array1 = append(array1, "a")
	array2 = append(array2, "b")
	array3 = append(array3, "c")

	array4 := array1
	array4[0] = "x"

	t.Log(array1)
	t.Log(array2)
	t.Log(array3)
	t.Log(array4)

}
