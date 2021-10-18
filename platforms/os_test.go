package platforms

import "testing"

func TestOS(t *testing.T) {

	p := Current()

	arch := p.Arch()
	os := p.OS()
	ver := p.GetOS().Version()

	t.Log("arch=", arch)
	t.Log("os=", os)
	t.Log("os.ver=", ver)
}
