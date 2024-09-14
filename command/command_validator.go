package command

import (
	"strings"
)

func fileExtensionValidator(args []string) bool {
	extension := strings.Split(args[0], ".")
	return extension[len(extension)-1] == "csv"
}
