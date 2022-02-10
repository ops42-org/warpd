package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/ops42-org/warpd/pkg/warpd"
	"github.com/ops42-org/warpd/pkg/warpd/wgithub"

	"github.com/spf13/cobra"
)

var (
	config string = "warpd.yaml"
)

func init() {
	githubCmd.PersistentFlags().StringVarP(&config, "config", "c", config, "warpd.yaml config file path")
	githubCmd.AddCommand(githubBuildmatrixCmd)
}

var githubBuildmatrixCmd = &cobra.Command{
	Use:   "buildmatrix",
	Short: "Generate build matrix for GitHub",
	Long:  `Helper function to generate GitHub build matrix`,
	Run: func(cmd *cobra.Command, args []string) {
		absPath, err := filepath.Abs(config)
		if err != nil {
			panic("Failed to get config absolute path: " + err.Error())
		}

		warpdConfig, err := warpd.LoadConfig(absPath)
		if err != nil {
			panic("Failed to load config: " + err.Error())
		}
		// for _, v := range warpdConfig.Build {
		// 	fmt.Println(*v.Path)
		// }

		matrix := wgithub.NewMatrixOutput(*warpdConfig, filepath.Dir(absPath))
		d, err := matrix.ToJson()
		if err != nil {
			panic("Failed to generate matrix yaml: " + err.Error())
		}
		fmt.Printf("::set-output name=matrix::%s", d)
	},
}
