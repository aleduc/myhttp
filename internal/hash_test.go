package internal

import "testing"

func TestHash_GetMD5(t *testing.T) {
	t.Parallel()

	h := &Hash{}
	actual := h.GetMD5([]byte("password"))
	expected := "5f4dcc3b5aa765d61d8327deb882cf99"
	if actual != expected {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, actual)
	}
}
