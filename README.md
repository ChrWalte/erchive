# erchive/zep

erchive is a go program that compresses and encrypts files and entire directories into .zep files (encrypted zip files).
it compresses using go's built-in zip compression and it encrypts using go's built-in advanced encryption standard algorithm.
it was built to be able to encrypt entire directories using only a single password; the compress is just a bonus.

## erchive

### download

the only way to download the executable is to download the source code.
the compiled binaries can be seen in the .bin folder where each version will be listed along with a zip of it.
the binaries are the compiled program from my Windows machine using arm64.
other platform binaries will be released later on
you may also compile the source code yourself using the golang compiler.
a real download method will be developed and the binaries will be removed.

### compile

to compile the source code one must have a go.17+ compiler installed.
compile like a normal go program.

### usage

to use erchive one must pass in console arguments for the directory or file to erchive and the password to use during encryption.
example:

- erchive.exe ./backup backup-password
- erchive.exe ./secrets secrets-password
- erchive ./secrets secrets-password

## folder structure

the folder and project structure was designed when I was learning many different system design patterns.
this heavily influenced the decisions I made in this project and I may have been too eager to try some.
the structure should be rethought and redesigned when a bigger picture can be seen.

### .bin

the temporary location of the compiled binaries.
each version can be found here alongside a zip version of each.

### data

the data access layer of the application.
this application only needs to access the file system to read and write the files erchive.

### service

the service layer of the application.
an encrypter service is used to handle encryption and decryption of byte data.
a hasher service is used to generate hashed bytes from byte data

## files

### .zip

represents a file format that is compressed using several algorithms. deflate is used in this implementation

### .zap

represents a file format that is compressed using the deflate algorithm and encrypted using the advanced encryption standard algorithm. data is protected by a password
