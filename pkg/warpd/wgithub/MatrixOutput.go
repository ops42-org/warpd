package wgithub

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ops42-org/warpd/pkg/util"
	"github.com/ops42-org/warpd/pkg/warpd"

	"gopkg.in/yaml.v2"
)

const (
	DEFAULT_BUILDER = "gcr.io/buildpacks/builder:v1"
)

type MatrixOutput struct {
	Include []MatrixInclude `json:"include"`
}

type MatrixInclude struct {
	Path       *string  `json:"path"`
	Name       *string  `json:"name"`
	Builder    *string  `json:"builder"`
	Buildpacks []string `json:"buildpacks"`
	Env        *string  `json:"env"`
}

func strMapToStr(m map[string]string, separator string) string {
	out := ""

	for k, v := range m {
		out += fmt.Sprintf("%s%s=%s", separator, k, v)
	}
	return strings.TrimLeft(out, separator)
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

				// Override image name with repo name for root path
				name := util.StrPtr(filepath.Base(match))
				if w, wPresent := os.LookupEnv("GITHUB_WORKSPACE"); wPresent {
					if match == w {
						if r, rPresent := os.LookupEnv("GITHUB_REPOSITORY"); rPresent {
							name = util.StrPtr(filepath.Base(r))
						}
					}
				}

				for i, v := range out.Include {
					if *v.Path == match {
						out.Include[i].Name = name
						out.Include[i].Builder = util.DefaultStrPtr(v.Builder, DEFAULT_BUILDER)
						out.Include[i].Buildpacks = b.Buildpacks
						out.Include[i].Env = util.StrPtr(strMapToStr(b.Env, " "))
						continue matchLoop
					}
				}

				out.Include = append(out.Include, MatrixInclude{
					Path:       util.StrPtr(match),
					Name:       name,
					Builder:    util.DefaultStrPtr(b.Builder, DEFAULT_BUILDER),
					Buildpacks: b.Buildpacks,
					Env:        util.StrPtr(strMapToStr(b.Env, " ")),
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
