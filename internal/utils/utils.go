package utils

import (
	"errors"
	"strings"
)

type PresentationPath struct {
	DirName  string // Directory name
	FileName string // File name
	FileData []byte // File data
}

func (f *PresentationPath) NewPresentationPath(dirName, fileName string, fileData []byte) {
	f.DirName = dirName
	f.FileName = fileName
	f.FileData = fileData
}

// RemoveSpaces removes all spaces from the fileName string.
func IgnoreSpaces(fileName string) string {
	var builder strings.Builder
	for _, char := range fileName {
		if char != ' ' {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}
func WrapError(message string, err error) error {
	return errors.New(message + ": " + err.Error())
}
