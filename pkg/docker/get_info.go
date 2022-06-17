package docker

import (
	"log"
	"os/exec"
	"strings"
)

func GetImageInfos() map[string][]string {
	out, err := exec.Command("docker", "images").Output()
	if err != nil {
		log.Fatal("not found")
	}

	arrayOutput := strings.Split(string(out), "\n")
	// headerと最後の1行が不要なので削除
	noHeaderAndTailArray := arrayOutput[1 : len(arrayOutput)-1]

	imageInfos := map[string][]string{}

	for _, value := range noHeaderAndTailArray {
		arrayValue := strings.Fields(value)
		imageName, imageTag := arrayValue[0], arrayValue[1]
		if _, ok := imageInfos[imageName]; ok {
			imageInfos[imageName] = append(imageInfos[imageName], imageTag)
			continue
		}
		imageInfos[imageName] = []string{imageTag}
	}

	return imageInfos
}

type ImageInfo struct {
	imageName string
	imageTags []string
}
