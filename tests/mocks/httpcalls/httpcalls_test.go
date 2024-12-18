package httpcalls_test

import (
	"articles/tests/mocks/httpcalls"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name        string
		description string
		server      *httptest.Server
		expect      *httpcalls.Response
		expectErr   error
	}{
		{
			name:        "server call success",
			description: "should return expected Response",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id": 1, "name": "Bob Marley", "description": "Jamaican Musician"}`))
			})),
			expect: &httpcalls.Response{
				Id:          1,
				Name:        "Bob Marley",
				Description: "Jamaican Musician",
			},
			expectErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.description)

			res, err := httpcalls.GetHttp(test.server.URL)
			if !reflect.DeepEqual(test.expect, res) {
				t.Errorf("expected %v, got %v", test.expect, res)
			}

			if err != test.expectErr {
				t.Errorf("expected error %v, got %v", test.expectErr, err)
			}
		})
	}
}
