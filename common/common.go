package common

var (
	// ProhibitedFiles represents the filenames that will not be encrypted or decrypted.
	ProhibitedFiles = []string{
		"System Volume Information",
		".Trashes",
		".Spotlight-V100",
		".fseventsd",
		".git",
	}

	// OSSlash represents the filepath slash delimiter for the current OS.
	OSSlash = "/"
)
