package onetimepad

// EncryptString encrypts a string with a random, one time key.
//
// Returns: key, encrypted string, error
func EncryptString(str string) ([]byte, string, error) {
	key, enc, err := EncryptBytes([]byte(str))
	return key, string(enc), err
}

// DecryptString decrypts a string with a key.
//
// Returns: decrypted string, error
func DecryptString(str string, key []byte) (string, error) {
	dec, err := DecryptBytes([]byte(str), key)
	return string(dec), err
}
