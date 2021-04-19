package collection

import "testing"

func TestSimplePropertiesAll(t *testing.T) {

	props := &SimpleProperties{}

	props.SetProperty("i", "12")
	props.SetProperty("j", "22")
	props.SetProperty("k", "32")

	props.SetProperty("a.r", "233")
	props.SetProperty("a.s", "2333")
	props.SetProperty("a.t", "2233")

	props.SetProperty("test.b.num", "2")
	props.SetProperty("test.xyz.num", "99")
	props.SetProperty("test.xyz.name", "foo")
	props.SetProperty("test.c.num", "3")
	props.SetProperty("test.a.num", "1")
	props.SetProperty("test.xyz.type", "bar")

	props.SetProperty("test.x.y.z.a", "yes")
	props.SetProperty("test.x.y.z.b", "no")

	text1 := FormatPropertiesWithSegment(props)
	props2, err := ParseProperties(text1, nil)
	text2 := FormatPropertiesWithSegment(props2)

	if err != nil {
		t.Error(err)
	}

	if text1 == text2 {
		t.Log(text1)
	} else {
		t.Error("text1 != text2")
	}

}
