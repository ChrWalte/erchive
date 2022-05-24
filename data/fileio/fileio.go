package fileio

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// the file permissions for read, write, and execute.
const filePermissions = 0777

// Read the file at the given path and return the contents as bytes.
// pathToFile: path to the file to read.
func ReadFile(pathToFile string) ([]byte, error) {
	// get file contents
	contents, err := os.ReadFile(pathToFile)

	// return contents and err
	return contents, err
}

// Read the file at the given path and return the contents as a string.
// pathToFile: path to the file to read.
func ReadFileAsString(pathToFile string) (string, error) {
	// read the file
	contentBytes, err := ReadFile(pathToFile)

	// convert to string.
	content := string(contentBytes)

	// return contents and err
	return content, err
}

// Reads the file size from the file at the given path.
// Returns the size of the file in bytes.
// pathToFile: path to the file to read.
func ReadFileSize(pathToFile string) (int64, error) {
	// open the file
	file, err := os.Open(pathToFile)

	// check for errors
	if err != nil {
		return -1, err
	}

	// read the file info
	fileInfo, err := file.Stat()

	// check for errors
	if err != nil {
		return -1, err
	}

	// return the file size
	return fileInfo.Size(), nil
}

// pathToFile: path to the file to write.
// contents: data to write to the file.
func WriteFile(pathToFile string, contents []byte) error {
	// write to the file.
	err := os.WriteFile(pathToFile, contents, filePermissions)

	// return err.
	return err
}

// Write the contents as a string to the file at the given path.
// pathToFile: path to the file to write.
// contents: data to write to the file.
func WriteFileWithString(filePath string, contents string) error {
	// convert from string to bytes.
	contentBytes := []byte(contents)

	// write to file using bytes.
	return WriteFile(filePath, contentBytes)
}

// Compresses a directory or file to a zip file
// pathToFile: path to the file or directory to compress.
// pathToDestination: path to the destination zip file.
func ZIPFile(pathToFile string, pathToDestination string) error {

	// create the archive file
	archive, err := os.Create(pathToDestination)

	// check for errors
	if err != nil {
		return err
	}

	// close the archive file on scope closure
	// need to handle the error better
	defer archive.Close()

	// create the zip writer
	zipWriter := zip.NewWriter(archive)

	// close the archive file on scope closure
	// need to handle the error better
	defer zipWriter.Close()

	// walk the given path handling each file, returning an error if one occurs
	return filepath.WalkDir(pathToFile, func(path string, entry fs.DirEntry, err error) error {
		// check for errors
		if err != nil {
			return err
		}

		// get the file information
		fileInfo, err := entry.Info()

		// check for errors
		if err != nil {
			return err
		}

		// create the file header
		header, err := zip.FileInfoHeader(fileInfo)

		// check for errors
		if err != nil {
			return err
		}

		// set the method of compression
		header.Method = zip.Deflate

		// set the name of the file or directory
		header.Name, err = filepath.Rel(filepath.Dir(pathToFile), path)

		// check for errors
		if err != nil {
			return err
		}

		// if the file is a directory, set the header to directory
		if entry.IsDir() {
			header.Name += "/"
		}

		// create the header writer to write to the file
		headerWriter, err := zipWriter.CreateHeader(header)

		// check for errors
		if err != nil {
			return err
		}

		// if the file is a directory, return
		if entry.IsDir() {
			return nil
		}

		// open the file
		file, err := os.Open(path)

		// check for errors
		if err != nil {
			return err
		}

		// close the file on scope closure
		defer file.Close()

		// copy the file data to the archive
		_, err = io.Copy(headerWriter, file)

		// return any errors
		return err
	})
}

// Compresses a directory or file to a zip file
// pathToFile: path to the file to decompress.
// pathToDestination: path to the destination where the un-zipped files will be.
func UnZIPFile(pathToFile string, pathToDestination string) error {

	// create the zip reader
	zipReader, err := zip.OpenReader(pathToFile)

	// check for errors
	if err != nil {
		return err
	}

	// close the archive file on scope closure
	defer zipReader.Close()

	// loop through the files in the archive
	for _, file := range zipReader.File {

		// get the path to the file
		filePath := filepath.Join(pathToDestination, file.Name)

		// check for zip slip.
		// https://www.geeksforgeeks.org/zip-slip/
		if !strings.HasPrefix(filePath, filepath.Clean(pathToDestination)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid file path: %s", filePath)
		}

		// check if the file is a directory
		if file.FileInfo().IsDir() {

			// if directory, create the directory tree
			err := os.MkdirAll(filePath, os.ModePerm)

			// check for errors
			if err != nil {
				return err
			}

			// move along to next file
			continue
		}

		// create all
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}

		// create the file to write to
		destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())

		// check for errors
		if err != nil {
			return err
		}

		// close the destination file on scope closure
		defer destinationFile.Close()

		// unzip the file contents
		zippedFile, err := file.Open()

		// check for errors
		if err != nil {
			return err
		}

		// close the zipped file on scope closure
		defer zippedFile.Close()

		// copy the file contents to the destination file
		_, err = io.Copy(destinationFile, zippedFile)

		// check for errors
		if err != nil {
			return err
		}
	}

	// return nil
	return nil
}
