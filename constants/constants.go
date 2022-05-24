package constants

const (
	// file extensions
	// file extension for compressed and encrypted files
	Dot_Zep = ".zep"
	// file extension for compressed files
	Dot_Zip = ".zip"

	// user messages
	// messaging showing the start of the compression
	Compression_Message = "handing compression..."
	// messaging showing the start of the encryption
	Encryption_Message = "handing encryption..."
	// message showing the start of decryption
	Decryption_Message = "handing decryption..."
	// message showing the start of decompression
	Decompression_Message = "handing decompression..."
	// message showing cleaning up temp files
	Cleaning_Message = "handing cleaning..."
	// message showing only decompressing files
	Only_Decompressing_Message = "only handing decompression..."
	// message showing exiting
	Exited_Message = "exited."

	// user errors
	// error message for invalid arguments
	Error_Directory_Not_Provided = "directory not provided\nusage: erchive <directory> <secret>"

	// magic numbers:
	// the number of required arguments
	Required_Args = 3

	// empty sting const
	// it is what it says
	Empty_String = ""
)
