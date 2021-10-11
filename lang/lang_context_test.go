package lang

import "testing"

func TestSimpleContext(t *testing.T) {
	err := tryTestSimpleContext(t)
	if err != nil {
		t.Error(err)
	}
}

func tryTestSimpleContext(t *testing.T) error {

	c1 := &SimpleContext{}

	err := SetupContext(c1)
	if err != nil {
		return err
	}

	c1.SetValue("foo", "bar")

	c2, err := GetContext(c1)
	if err != nil {
		return err
	}

	value := c2.Value("foo")
	t.Log("foo=", value)

	return nil
}
