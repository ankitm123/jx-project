package draft

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/draft/pkg/linguist"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
)

// copied from draft so we can change to lookup jx draft packs
// credit original from: https://github.com/Azure/draft/blob/8e1a459/cmd/draft/create.go#L163

// DoPackDetectionForBuildPack performs detection of the language based on a sepcific build pack
func DoPackDetectionForBuildPack(out io.Writer, dir, packDir string) (string, error) {
	log.Logger().Infof("performing pack detection in folder %s", dir)
	langs, err := linguist.ProcessDir(dir)
	if err != nil {
		return "", fmt.Errorf("there was an error detecting the language: %s", err)
	}
	if len(langs) == 0 {
		return "", fmt.Errorf("there was an error detecting the language")
	}
	for _, lang := range langs {
		detectedLang := linguist.Alias(lang)
		fmt.Fprintf(out, "--> Draft detected %s (%f%%)\n", detectedLang.Language, detectedLang.Percent)
		packs, err := os.ReadDir(packDir)
		if err != nil {
			return "", fmt.Errorf("there was an error reading %s: %v", packDir, err)
		}
		for _, file := range packs {
			if file.IsDir() {
				if strings.EqualFold(detectedLang.Language, file.Name()) {
					packPath := filepath.Join(packDir, file.Name())
					return packPath, nil
				}
			}
		}
		fmt.Fprintf(out, "--> Could not find a pack for %s. Trying to find the next likely language match...\n", detectedLang.Language)
	}
	return "", fmt.Errorf("there was an error detecting the language using packs from %s", packDir)
}
