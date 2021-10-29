package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
)

// hashes the given bytes using the SHA512 algorithm
// returns the resulting hash
// bytes: the bytes to hash
func SHA512(bytes []byte) []byte {
	// hash the bytes
	sha512 := sha512.Sum512(bytes)

	// return the hash
	return sha512[:]
}

// hashes the given bytes using the SHA256 algorithm
// returns the resulting hash
// bytes: the bytes to hash
func SHA256(bytes []byte) []byte {
	// hash the bytes
	sha256 := sha256.Sum256(bytes)

	// return the hash
	return sha256[:]
}

// hashes the given bytes using the MD5 algorithm
// returns the resulting hash
// bytes: the bytes to hash
func MD5(bytes []byte) []byte {
	// hash the bytes
	md5 := md5.Sum(bytes)

	// return the hash
	return md5[:]
}
