package download

import (
	"github.com/bmuschko/letsgopher/template/storage"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// TemplateDownloader retrieves a template archive from an URL.
type TemplateDownloader struct {
	Home   storage.Home
	Getter Getter
}

// Download downloads a template archive from an URL.
func (td *TemplateDownloader) Download(url string) (string, error) {
	data, err := td.Getter.Get(url)
	if err != nil {
		return "", err
	}

	destfile := filepath.Join(td.Home.ArchiveDir(), extractTemplateName(url))
	if err := ioutil.WriteFile(destfile, data.Bytes(), 0644); err != nil {
		return "", err
	}

	return destfile, nil
}

func extractTemplateName(url string) string {
	lastDotSlash := strings.LastIndex(url, "/")
	r := []rune(url)
	return string(r[lastDotSlash:len(url)])
}
