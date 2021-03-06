package cui_input

import (
	"GPUContainerRunnerGo/pkg/docker"
	"GPUContainerRunnerGo/pkg/gpu"
	"fmt"
	"strconv"
)

func InputUser() ScriptInfo {

	var volumePath string
	var filePath string
	var gpuId string
	var imageName string
	var imageTag string
	var logPath string

	fmt.Println("Input your volume path: ")
	fmt.Scan(&volumePath)
	fmt.Println("" + volumePath)

	fmt.Print("Input your file(Python) path:")
	fmt.Scan(&filePath)

	gpuInfos := gpu.GetInfo()
	for _, gpuInfo := range gpuInfos {
		fmt.Println("ID: " + gpuInfo.Id + "(" + strconv.FormatFloat(gpuInfo.MemoryUsage, 'f', -1, 64) + "%)")
	}
	fmt.Print("Input your gpu id: ")
	fmt.Scan(&gpuId)

	imageInfos := docker.GetImageInfos()
	fmt.Println("exist docker images")
	for key := range imageInfos {
		fmt.Println("- ", key)
	}
	fmt.Print("Input your image name: ")
	fmt.Scan(&imageName)

	fmt.Println("exist docker image tags in ", imageName)
	for imageTag := range imageInfos[imageName] {
		fmt.Println("- ", imageTag)
	}
	fmt.Print("Input your image tag: ")
	fmt.Scan(&imageTag)

	fmt.Print("Input your log path: ")
	fmt.Scan(&logPath)

	return ScriptInfo{
		VolumePath: volumePath,
		FilePath:   filePath,
		GPUId:      gpuId,
		ImageName:  imageName,
		ImageTag:   imageTag,
		LogPath:    logPath,
	}
}
