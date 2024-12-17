package ifacesubstitution_test

import (
	"articles/tests/mocks/ifacesubstitution"
	"io"
	"reflect"
	"testing"
)

type readCloser struct {
	expect    []byte
	expectErr error
}

// allow tests to fill expect and expectErr for each test
func (rc *readCloser) Read(p []byte) (n int, err error) {
	copy(p, rc.expect)
	return len(p), rc.expectErr
}

// do nothing
func (rc *readCloser) Close() error {
	return nil
}

func TestReadContents(t *testing.T) {
	tests := []struct {
		name        string
		description string
		file        io.ReadCloser
		expect      []byte
		expectErr   error
	}{
		{
			name:        "succesful read content",
			description: "content should be read and returned without errors",
			file: &readCloser{
				expect:    []byte("hello"),
				expectErr: nil,
			},
			expect:    []byte("hello"),
			expectErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.description)

			res, err := ifacesubstitution.ReadContents(test.file, 5)

			if err != test.expectErr {
				t.Errorf("expetec err %v, got %v", test.expectErr, err)
			}

			if !reflect.DeepEqual(test.expect, res) {
				t.Errorf("expetec %v, got %v", test.expect, res)
			}
		})
	}
}
