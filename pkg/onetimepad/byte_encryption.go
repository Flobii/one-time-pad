package onetimepad

import "errors"

// EncryptBytes encrypts a byte array with a random, one time key.
//
// Returns: key, encrypted byte array, error
func EncryptBytes(b_str []byte) ([]byte, []byte, error) {
	key, err := GenerateNewKey(len(b_str))
	if isError(err) {
		return nil, nil, err
	}

	if len(b_str) != len(key) {
		return nil, nil, errors.New("key and byte string not same length")
	}

	enc, err := xorBytes(key, b_str)
	if isError(err) {
		return nil, nil, err
	}
	return key, enc, nil
}

// DecryptBytes decrypts a byte array with a key.
//
// Returns: decrypted byte array, error
func DecryptBytes(b_str []byte, key []byte) ([]byte, error) {
	dec, err := xorBytes(key, b_str)
	if isError(err) {
		return nil, err
	}
	return dec, nil
}

func xorBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("a must be same length as b")
	}
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c, nil
}
