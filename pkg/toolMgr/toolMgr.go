package toolMgr

import (
	"fmt"
	"os/exec"

	"github.com/go-logr/logr"
)

func main() {
	fmt.Println("vim-go")
}

type ToolMgrIf interface {
	InstallByDir(dir string) error
}

type ToolMgr struct {
	log *logr.Logger
}

func NewToolMgr(log *logr.Logger) ToolMgr {
	return ToolMgr{
		log: log,
	}
}

// InstallByDir installed all the necessary tools defined in directory tool versions files
func (t *ToolMgr) InstallByDir(dir string) error {
	cmd := exec.Command("asdf", "install")
	cmd.Dir = dir
	outBytes, err := cmd.Output()
	if err != nil {
		t.log.Error(err, "Failed to run 'asdf install'")
	}
	o := string(outBytes)
	t.log.Info(o)
	return err
}
