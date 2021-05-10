package internal

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)


type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (r RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}
func NewFakeClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestHTTPWrap_MakeRequest(t *testing.T) {
	t.Parallel()

	type testData struct {
		tcase string
		turl string
		roundTripFn RoundTripFunc
		expected []byte
	}

	testTableData := []testData{
		{
			tcase: "success",
			roundTripFn: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					Body: io.NopCloser(strings.NewReader("=")),
				}, nil
			},
			expected: []byte{61},
		},
		{
			tcase: "request error",
			roundTripFn: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("some error")
			},
			expected: []byte{},
		},
	}

	for _, testUnit := range testTableData {
		wr := NewHTTPWrap(NewFakeClient(testUnit.roundTripFn))
		actual := wr.MakeRequest(testUnit.turl)
		if !bytes.Equal(actual, testUnit.expected) {
			t.Errorf("\n...expected = %v\n...obtained = %v", testUnit.expected, actual)
		}
	}
}
