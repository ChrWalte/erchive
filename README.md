# erchive/zep

erchive is a go program that compresses and encrypts files and entire directories into .zep files (encrypted zip files).
it compresses using go's built-in zip compression and it encrypts using go's built-in advanced encryption standard algorithm.
it was built to be able to encrypt entire directories using only a single password; the compress is just a bonus.

## erchive

### usage

to use erchive one must pass in console arguments for the directory or file to erchive and the password to use during encryption.
example:

- erchive.exe ./backup backup-password
- erchive.exe ./secrets secrets-password
- erchive ./secrets secrets-password

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

## files

### .zip

represents a file format that is compressed using several algorithms. deflate is used in this implementation

### .zap

represents a file format that is compressed using the deflate algorithm and encrypted using the advanced encryption standard algorithm. data is protected by a password
