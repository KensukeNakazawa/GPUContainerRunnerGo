package run

import (
	"GPUContainerRunnerGo/pkg/cui_input"
	"reflect"
	"testing"
)

func TestGenerateCommand(t *testing.T) {
	t.Run("docker コマンドが生成されること", func(t *testing.T) {
		scriptInfo := cui_input.ScriptInfo{
			ImageTag:   "3",
			ImageName:  "python",
			GPUId:      "1",
			FilePath:   "rungpu.py",
			VolumePath: "/home/rungpu",
			LogPath:    "./",
		}
		expectArgs := []string{
			"run",
			"-i",
			"--rm",
			"-gpus 1",
			"-v /home/rungpu/:/var/run_gpu",
			"-w /var/run_gpu",
			"python:3",
			"python rungpu.py",
		}
		args := GenerateArgs(scriptInfo)
		if !reflect.DeepEqual(args, expectArgs) {
			t.Errorf("args = %v, expect %v", args, expectArgs)
		}

	})
}
