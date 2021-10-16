package contexts

import "testing"

func TestSimpleContext(t *testing.T) {
	err := tryTestSimpleContext(t)
	if err != nil {
		t.Error(err)
	}
}

func tryTestSimpleContext(t *testing.T) error {

	c1 := &SimpleContext{}

	err := SetupContextSetter(c1)
	if err != nil {
		return err
	}

	c1.SetValue("foo", "bar")

	c2, err := GetContextSetter(c1)
	if err != nil {
		return err
	}

	value := c2.GetContext().Value("foo")
	t.Log("foo=", value)

	return nil
}
