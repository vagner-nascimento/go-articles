package httpcalls_test

import (
	"articles/tests/mocks/httpcalls"
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name      string
		server    *httptest.Server
		expectRes *httpcalls.Response
		expectErr error
	}{
		{
			name:   "server response 200 - OK with content",
			server: nil, //httptest.NewServer(),
			expectRes: &httpcalls.Response{Id: 1,
				Name:        "Bob Marley",
				Description: "Jamaican Musician",
			},
			expectErr: nil,
		},
	}

	fmt.Println(testCases)
}
