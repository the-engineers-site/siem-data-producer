package fileUtils

import (
	"os"
	"strings"
)

const (
	FileUploadFolder = "/tmp/storage/logs/"
	Permissions      = 0777
)

func CreateOutputFolder(deviceType string, deviceVendor string) (string, error) {
	path := FileUploadFolder + formatPath(deviceType) + "/" + formatPath(deviceVendor)
	err := os.MkdirAll(path, Permissions)
	return path, err

}

func formatPath(path string) string {
	return strings.ReplaceAll(path, " |/|\\", "_")
}

func CreateFile(path string) (*os.File, error) {
	return os.Create(path)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}
