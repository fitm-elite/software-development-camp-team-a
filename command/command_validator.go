package command

import (
	"fmt"
	"strings"
)

func fileExtensionValidator(args []string) bool {
	extension := strings.Split(args[0], ".")

	fmt.Println("extension: ", extension[len(extension)-1])

	return extension[len(extension)-1] == "csv"
}
