package wgithub

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ops42-org/warpd/pkg/util"
	"github.com/ops42-org/warpd/pkg/warpd"

	"gopkg.in/yaml.v2"
)

type MatrixOutput struct {
	Include []MatrixInclude `json:"include"`
}

type MatrixInclude struct {
	Path       *string  `json:"path"`
	Name       *string  `json:"name"`
	Buildpacks []string `json:"buildpacks"`
}

func NewMatrixOutput(config warpd.WarpdConfig, rootDir string) *MatrixOutput {
	var out MatrixOutput

	for _, b := range config.Build {
		matches, err := filepath.Glob(filepath.Join(rootDir, *b.Path))
		if err != nil {
			fmt.Println("Failed to glob path '" + *b.Path + "': " + err.Error())
			continue
		}

		var dirs []string
	matchLoop:
		for _, match := range matches {
			f, _ := os.Stat(match)
			if f.IsDir() {
				dirs = append(dirs, match)

				for i, v := range out.Include {
					if *v.Path == match {
						out.Include[i].Name = util.StrPtr(filepath.Base(match))
						out.Include[i].Buildpacks = b.Buildpacks
						continue matchLoop
					}
				}

				out.Include = append(out.Include, MatrixInclude{
					Path:       util.StrPtr(match),
					Name:       util.StrPtr(filepath.Base(match)),
					Buildpacks: b.Buildpacks,
				})
			}
		}
	}

	return &out
}

// TODO: Not used
func (m *MatrixOutput) ToYaml() (string, error) {
	d, err := yaml.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("Failed to marhsal MatrixOutput: ", err.Error())
	}
	return string(d), nil
}

func (m *MatrixOutput) ToJson() (string, error) {
	d, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("Failed to marhsal MatrixOutput: ", err.Error())
	}
	return string(d), nil
}
