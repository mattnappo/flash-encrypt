package common

var (
	// ProhibitedFiles represents the filenames that will not be encrypted or decrypted.
	ProhibitedFiles = []string{
		"System Volume Information",
		".Trashes",
		".Spotlight-V100",
		".fseventsd",
	}
)
