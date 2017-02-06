package base64

import (
	"io/ioutil"
	"path/filepath"
	"mime"
	"encoding/base64"
	"os"
)

type imageFile struct {
	path string
	ext  string
	mime string
	size int64
}

func NewImageFile(path string) (imageFile, error) {
	ext := filepath.Ext(path)
	mimeType := mime.TypeByExtension(ext)
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		return imageFile{}, err
	}

	fi, err := fp.Stat()
	if err != nil {
		return imageFile{}, err
	}

	return imageFile{path: path, ext: ext, mime: mimeType, size: fi.Size()}, err
}

/*
	Read Image file
 */
func (imageFile *imageFile) Content() ([]byte, string, error) {
	content, err := ioutil.ReadFile(imageFile.path)
	return content, imageFile.mime, err
}

/*
	Encode Image File to Base 64
 */
func (imageFile *imageFile) EncodeBase64() (string, string, error) {
	imageBytes, mimeType, err := imageFile.Content()
	if err != nil {
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(imageBytes), mimeType, err
}
