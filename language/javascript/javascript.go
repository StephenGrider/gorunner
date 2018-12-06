package javascript

import (
	"github.com/stephengrider/gorunner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, "node", files[0])
}
