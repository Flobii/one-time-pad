package onetimepad

import "io/ioutil"

// EncryptFile encrypts a file with a random, one time key.
//
// Returns: key, error
func EncryptFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if isError(err) {
		return nil, err
	}

	enc, key, err := EncryptBytes(data)
	if isError(err) {
		return nil, err
	}

	ioutil.WriteFile(path, enc, 0644)
	return key, nil
}

// DecryptFile decrypts a file with a key.
//
// Returns: error
func DecryptFile(path string, key []byte) error {
	data, err := ioutil.ReadFile(path)
	if isError(err) {
		return nil
	}

	dec, err := DecryptBytes(data, key)
	if isError(err) {
		return err
	}

	ioutil.WriteFile(path, dec, 0644)
	return nil
}
