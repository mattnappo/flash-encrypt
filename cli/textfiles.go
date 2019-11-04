package cli

const (
	// Help stores the help text for the help menu.
	Help = `    encrypt(int driveID)
      Encrypt a flash drive with id driveID.

    decrypt(int driveID)
      Decrypt a flash drive with id driveID.

    drives()
      Returns a list of the driveIDs of all connected drives.

    help()
      Display this help menu.

    exit()
		Exit flash-encrypt.`

	// StandaloneHelp stores the help text for the standalone help menu.
	StandaloneHelp = `    encrypt()
	  Encrypt ./secure folder.

  	decrypt()
	  Decrypt the ./secure folder.

  	help()
	  Display this help menu.

  	exit()
	  Exit flash-encrypt.`
)
