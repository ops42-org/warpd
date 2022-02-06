package deployer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/ops42-org/warpd/pkg/helm"
	"github.com/ops42-org/warpd/pkg/toolMgr"
)

func main() {
	fmt.Println("vim-go")
}

type Deployer struct {
	workDir string
	toolMgr *toolMgr.ToolMgr
	helm    *helm.Helm
	log     *logr.Logger
}

func NewDeployer(workDir string, log *logr.Logger) Deployer {
	toolMgr := toolMgr.NewToolMgr(log)
	helm := helm.NewHelm(log)
	return Deployer{
		workDir: workDir,
		toolMgr: &toolMgr,
		helm:    &helm,
		log:     log,
	}
}

func (d *Deployer) InstallTools() {
	d.log.Info("Installing tools...")
	installSubdirs := []string{"chart", "terraform"}
	files, err := ioutil.ReadDir(d.workDir)
	if err != nil {
		d.log.Error(err, "Failed to read directory contents wile installing tools")
	} else {
		for _, f := range files {
			if f.IsDir() {
				for _, subdir := range installSubdirs {
					dir := filepath.Join(d.workDir, f.Name(), subdir)
					dirStat, err := os.Stat(dir)
					if err != nil {
						continue
					}
					if dirStat.IsDir() {
						d.log.Info("Installing tools in dir", "dir", dir)
						d.toolMgr.InstallByDir(dir)
					}
				}
			}
		}
	}
}

func (d *Deployer) DeployComponents() {
	d.log.Info("Deploying components...")
	files, err := ioutil.ReadDir(d.workDir)
	if err != nil {
		d.log.Error(err, "Failed to read directory contents while deploying")
	} else {
		for _, f := range files {
			if f.IsDir() {
				d.log.Info("Deploying component", "component", f.Name())
				d.DeployHelm(f.Name())
			}
		}
	}
}

func (d *Deployer) DeployHelm(component string) error {
	log := d.log.WithValues("component", component)
	log.Info("Deploying Helm...")
	dir := filepath.Join(d.workDir, component, "chart")
	dirStat, err := os.Stat(dir)
	if err != nil {
		return nil
	}
	if dirStat.IsDir() {
		log.Info("Found Helm chart. Will deploy...")
		d.helm.DepUp(dir)
	}
	return nil
}
