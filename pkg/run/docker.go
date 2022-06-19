// Package run
// nvidia-docker を用いて、Python スクリプトを実行するためのパッケージ
package run

import (
	"GPUContainerRunnerGo/pkg/cui_input"
	"log"
	"os/exec"
	"strings"
)

// Run
// nvidia-docker を利用して、指定のdocker imageのコンテナを立ち上げ、
// python スクリプトを実行する
func Run(scriptInfo cui_input.ScriptInfo) []string {
	args := GenerateArgs(scriptInfo)
	output, err := exec.Command("docker", args...).Output()
	if err != nil {
		log.Fatal("not found")
	}

	return strings.Split(string(output), "\n")
}

// GenerateArgs
// nvidia-docker を利用して実行する際のスクリプトに渡すための引数を生成する
// @see https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/user-guide.html
func GenerateArgs(scriptInfo cui_input.ScriptInfo) []string {
	gpuId := scriptInfo.GPUId
	workingDir := "/var/run_gpu"
	volumes := scriptInfo.VolumePath + "/:" + workingDir
	dockerImageInfo := scriptInfo.ImageName + ":" + scriptInfo.ImageTag
	targetScript := scriptInfo.FilePath
	args := []string{
		"run",
		"-i",
		"--rm",
		"-gpus" + " " + gpuId,
		"-v" + " " + volumes,
		"-w" + " " + workingDir,
		dockerImageInfo,
		"python " + targetScript,
	}

	return args
}
