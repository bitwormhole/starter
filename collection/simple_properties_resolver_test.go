package collection

import "testing"

func TestResolvePropertyVar(t *testing.T) {

	props := CreateProperties()

	props.SetProperty("a", "hello")
	props.SetProperty("b", "{{a }}, world")
	props.SetProperty("c", "{{ b  }}!")
	props.SetProperty("d", "{{  a}}, say: {{ c }}")

	table := props.Export(nil)
	for k, v := range table {
		t.Log("property:    ", k, " = ", v)
	}

	err := ResolvePropertiesVarWithTokens("{{", "}}", props)
	if err != nil {
		t.Error(err)
		return
	}

	table = props.Export(nil)
	for k, v := range table {
		t.Log("property:    ", k, " = ", v)
	}
}
