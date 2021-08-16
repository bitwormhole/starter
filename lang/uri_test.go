package lang

import "testing"

func TestCreateURI(t *testing.T) {

	s1 := "http://user@host.com:8888/path/a/b/c?q=1#fragment666"

	uri1, err := ParseURI(s1)
	if err != nil {
		t.Error(err)
	}

	uri2 := CreateURI(uri1.URL())
	s2 := uri2.String()

	if s2 != s1 {
		t.Error("s1 != s2, s1=[", s1, "], s2=[", s2, "]")
	}
}

func TestFileURI(t *testing.T) {

	s1 := "file:/C:/a/bc/de/b"

	uri1, err := ParseURI(s1)
	if err != nil {
		t.Error(err)
	}

	uri2 := CreateURI(uri1.URL())
	s2 := uri2.String()
	path1 := uri1.Path()
	path2 := uri2.Path()

	t.Log("uri.path1=", path1)
	t.Log("uri.path2=", path2)

	if path1 != path2 {
		t.Error("s1 != s2, s1=[", s1, "], s2=[", s2, "]")
	}
}
