package helm

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/go-logr/logr"

	"github.com/ops42-org/warpd/pkg/util"
)

const (
	execRetryAttempts = 3
	execRetrySleep    = 3
)

func main() {
	fmt.Println("vim-go")
}

type Helm struct {
	executable *string
	dir        *string
	log        *logr.Logger
}

func NewHelm(log *logr.Logger) Helm {
	return Helm{
		executable: util.StrPtr("helm"),
		log:        log,
	}
}

func (h *Helm) ExecWithRetry(args []string, sleepDuration time.Duration, retryAttempts int) (string, error) {
	var err error = nil
	sleep := sleepDuration
	for i := 0; i < retryAttempts; i++ {
		if i > 0 {
			h.log.Error(err, "Retrying after error")
			time.Sleep(sleep)
			sleep *= 2
		}
		o, err := h.Exec(args)
		if err == nil {
			return o, nil
		}
	}
	return "", fmt.Errorf("Failed after %d attempts, last error: %s", retryAttempts, err)
}

func (h *Helm) Exec(args []string) (string, error) {
	outBytes, err := exec.Command(*h.executable, args...).Output()
	if err != nil {
		h.log.Error(err, "Failed to run helm", "executable", *h.executable, "args", args)
	}
	o := string(outBytes)
	h.log.Info(o)
	return string(o), err
}

func (h *Helm) Version() (string, error) {
	args := []string{"version"}
	return h.Exec(args)
}

func (h *Helm) DepUp(chart string) (string, error) {
	args := []string{"dependency", "update", chart}
	return h.ExecWithRetry(args, execRetrySleep*time.Second, execRetryAttempts)
}
