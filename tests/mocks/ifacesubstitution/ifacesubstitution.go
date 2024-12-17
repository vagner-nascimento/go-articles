package ifacesubstitution

import "io"

//before refactor: func ReadContents(f *os.File, byts int) ([]byte, error)
func ReadContents(f io.ReadCloser, byts int) ([]byte, error) {
	defer f.Close()
	data := make([]byte, byts)

	if _, err := f.Read(data); err != nil {
		return nil, err
	}

	return data, nil
}
