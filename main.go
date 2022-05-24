package main

import (
	"encoding/hex"
	"erchive/constants"
	"erchive/data/fileio"
	"erchive/service/encrypter"
	"erchive/service/hasher"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// compression and encryption needs to happen in this order:
// [ compress ] -> [  encrypt     ]
// [ decrypt  ] -> [  decompress  ]
// https://stackoverflow.com/questions/4676095/when-compressing-and-encrypting-should-i-compress-first-or-encrypt-first?noredirect=1&lq=1

// the main entry function
func main() {
	// handle arguments to obtain the path to the file and the key
	useDir, key := handleArgs()

	// if useDir, is Empty_String: panic
	if useDir == constants.Empty_String {
		panic(constants.Error_Directory_Not_Provided)
	}

	// get the current extension for the given dir
	extension := filepath.Ext(useDir)

	// get the hashed version of the password
	keyBytes := hasher.SHA256([]byte(key))

	// switch on the extension
	switch extension {
	case constants.Dot_Zep: // [file.zep]

		// decrypt the file
		fmt.Println(constants.Decryption_Message)
		start := time.Now()
		handleDecryption(useDir, keyBytes)
		elapsed := time.Since(start)
		fmt.Println("finished in [", elapsed, "].")

		// drop the .zep extension
		removedZep := strings.TrimSuffix(useDir, constants.Dot_Zep)

		// add the .zip extension
		zipDir := removedZep + constants.Dot_Zip

		// decompress the file
		fmt.Println(constants.Decompression_Message)
		start = time.Now()
		handleDecompression(zipDir)
		elapsed = time.Since(start)
		fmt.Println("finished in [", elapsed, "].")

		// remove the regular zip file
		fmt.Println(constants.Cleaning_Message)
		start = time.Now()
		os.Remove(zipDir)
		elapsed = time.Since(start)
		fmt.Println("finished in [", elapsed, "].")

	case constants.Dot_Zip: // [file.zip]

		// compress the file
		fmt.Println(constants.Only_Decompressing_Message)
		start := time.Now()
		handleDecompression(useDir)
		elapsed := time.Since(start)
		fmt.Println("finished in [", elapsed, "].")

	// if the extension is not .zep or .zip, then it is a directory
	default:

		// compress the file
		fmt.Println(constants.Compression_Message)
		start := time.Now()
		handleCompression(useDir)
		elapsed := time.Since(start)
		fmt.Println("finished in [", elapsed, "].")

		// add the .zip extension
		zipDir := useDir + constants.Dot_Zip

		// encrypt the file
		fmt.Println(constants.Encryption_Message)
		start = time.Now()
		handleEncryption(zipDir, keyBytes)
		elapsed = time.Since(start)
		fmt.Println("finished in [", elapsed, "].")

		// remove the regular zip file
		fmt.Println(constants.Cleaning_Message)
		start = time.Now()
		os.Remove(zipDir)
		elapsed = time.Since(start)
		fmt.Println("finished in [", elapsed, "].")
	}
	fmt.Println(constants.Exited_Message)
}

// handles the arguments passed to the program
// returns the path to the file and the key if given
func handleArgs() (string, string) {

	// if length is Required_Args
	if len(os.Args) == constants.Required_Args {
		// return the path to the file and the key
		return os.Args[1], os.Args[2]
	}

	// return nothing
	return "", ""
}

// handles the compression of the file
// ends with a new file with the .zip extension
// pathToFile: the path to the file to be compressed
func handleCompression(pathToFile string) {
	// compress the file
	err := fileio.ZIPFile(pathToFile, pathToFile+constants.Dot_Zip)

	// check for errors
	if err != nil {
		panic(err)
	}
}

// handles the decompression of the file
// ends with the decompressed data in a file or directory
// pathToFile: the path to the file to be decompressed
func handleDecompression(pathToFile string) {
	// get the destination path by removing the .zip extension
	destination := strings.TrimSuffix(pathToFile, constants.Dot_Zip)

	// decompress the file
	err := fileio.UnZIPFile(pathToFile, destination)

	// check for errors
	if err != nil {
		panic(err)
	}
}

// handles the encryption of the file
// ends with a new file with the .zep extension
// pathToFile: the path to the file to be encrypted
// key: the key to be used for the encryption
func handleEncryption(pathToFile string, key []byte) {
	// read the file contents
	contents, err := fileio.ReadFile(pathToFile)

	// check for errors
	if err != nil {
		panic(err)
	}

	// encrypt the file contents
	encryptedData := encrypter.EncryptData(key, contents)

	// get the destination path by removing the .zip extension
	destination := strings.TrimSuffix(pathToFile, constants.Dot_Zip)

	// write the encrypted data to the file
	err = fileio.WriteFile(destination+constants.Dot_Zep, encryptedData)

	// check for errors
	if err != nil {
		panic(err)
	}
}

// handles the decryption of the file
// ends with the decrypted data in a .zip file
// pathToFile: the path to the file to be decrypted
// key: the key to be used for the decryption
func handleDecryption(pathToFile string, key []byte) {
	// read the file contents
	contents, err := fileio.ReadFile(pathToFile)

	// check for errors
	if err != nil {
		panic(err)
	}

	// decrypt the file contents
	decryptedData := encrypter.DecryptData(key, contents)

	// get the destination path by removing the .zep extension
	destination := strings.TrimSuffix(pathToFile, constants.Dot_Zep)

	// write the decrypted data to the file
	err = fileio.WriteFile(destination+constants.Dot_Zip, decryptedData)

	// check for errors
	if err != nil {
		panic(err)
	}
}

// TODO: show hash of file, zip, and zep
func getFileHash(pathToFile string) string {
	// get the file contents
	contents, err := fileio.ReadFile(pathToFile)

	// check for errors
	if err != nil {
		panic(err)
	}

	// get the hash of the file contents
	hash := hasher.SHA256(contents)

	// return hex-string version of the hash
	return hex.EncodeToString(hash[:])
}

// TODO: show file size of file, zip, and zep
func GetFileSize(pathToFile string) string {
	// get the file size
	fileSize, err := fileio.ReadFileSize(pathToFile)

	// check for errors
	if err != nil {
		panic(err)
	}

	// return the file size as a string
	return strconv.FormatInt(fileSize, 10)
}
