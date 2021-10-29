package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// encrypt given data using the AES cipher and the given key
// returns the encrypted data
// key: the key to use for encryption
// data: the data to encrypt
// Big-O-Notation: N/A Due to the Unknown Nature of the Algorithm
func EncryptData(key []byte, data []byte) []byte {
	// create a new GCM cipher using the given key
	gcmCipher := createGCMCipher(key)

	// create a new nonce of default size
	nonce := make([]byte, gcmCipher.NonceSize())

	// read random data into the nonce
	_, err := io.ReadFull(rand.Reader, nonce)

	// check for errors
	if err != nil {
		panic(err)
	}

	// encrypt the data using the GCM cipher
	encryptedFileBytes := gcmCipher.Seal(nonce, nonce, data, nil)

	// return the encrypted data
	return encryptedFileBytes
}

// decrypt given data using the AES cipher and the given key
// returns the decrypted data
// key: the key to use for decryption
// data: the data to decrypt
// Big-O-Notation: N/A Due to the Unknown Nature of the Algorithm
func DecryptData(key []byte, data []byte) []byte {
	// create a new GCM cipher using the given key
	gcmCipher := createGCMCipher(key)

	// create a new nonce of default size
	nonceSize := gcmCipher.NonceSize()

	// split the data into the nonce and the encrypted data
	nonce, encryptedBytes := data[:nonceSize], data[nonceSize:]

	// decrypt the data using the GCM cipher
	decryptedFilesBytes, err := gcmCipher.Open(nil, nonce, encryptedBytes, nil)

	// check for errors
	if err != nil {
		panic(err)
	}

	// return the decrypted data
	return decryptedFilesBytes
}

// create a new GCM cipher using the given key
// returns the cipher
// key: the key to use for the AES cipher
// Big-O-Notation: N/A Due to the Unknown Nature of the Algorithm
func createGCMCipher(key []byte) cipher.AEAD {
	// create a new AES cipher using the given key
	aesCipher, err := aes.NewCipher(key)

	// check for errors
	if err != nil {
		panic(err)
	}

	// create a new GCM cipher using the AES cipher
	gcmCipher, err := cipher.NewGCM(aesCipher)

	// check for errors
	if err != nil {
		panic(err)
	}

	// return the GCM cipher
	return gcmCipher
}
