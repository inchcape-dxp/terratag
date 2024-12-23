package file

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"go.uber.org/multierr"
	"gopkg.in/yaml.v3"
)

func CreatingBackup(path string) error {
	backupFilename := path + ".bak"

	log.Print("[INFO] Backing up ", path, " to ", backupFilename)

	if err := os.Rename(path, backupFilename); err != nil {
		return err
	}

	return nil
}

func ReplaceWithTerratagFile(path string, textContent string, rename bool) error {
	if rename {
		taggedFilename := strings.TrimSuffix(path, filepath.Ext(path)) + ".terratag.tf"
		if err := CreateFile(taggedFilename, textContent); err != nil {
			return err
		}
	}

	if !rename {
		if err := CreateFile(path, textContent); err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(path string, textContent string) error {
	log.Print("[INFO] Creating file ", path)

	return os.WriteFile(path, []byte(textContent), 0644)
}

func GetFilename(path string) string {
	_, filename := filepath.Split(path)
	filename = strings.TrimSuffix(filename, filepath.Ext(path))
	filename = strings.ReplaceAll(filename, ".", "-")

	return filename
}

func ReadHCLFile(path string) (*hclwrite.File, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	file, diagnostics := hclwrite.ParseConfig(src, path, hcl.InitialPos)
	if err := multierr.Combine(diagnostics.Errs()...); err != nil {
		return nil, err
	}

	return file, nil
}

func ReadYAMLFile(filePath string) (map[string]interface{}, error) {
	// Read the YAML file content
	yamlContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the YAML into a dynamic map
	var result map[string]interface{}
	err = yaml.Unmarshal(yamlContent, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ConvertToJSON(data map[string]interface{}) (string, error) {
	// Marshal the map into JSON format
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
