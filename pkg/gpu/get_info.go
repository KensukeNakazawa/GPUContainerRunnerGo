package gpu

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

func GetInfo() []Info {
	out, err := exec.Command("nvidia-smi", "--query-gpu=index,memory.used,memory.total", "--format=csv,noheader,nounits").CombinedOutput()
	if err != nil {
		return []Info{{Id: "-1", MemoryUsage: 0.0}}
	}

	// カンマを消して改行コードで分割した後に、不要な行を削除する
	afterOutput := strings.ReplaceAll(string(out), ",", " ")
	arrayOutput := strings.Split(afterOutput, "\n")
	noHeaderAndTailArray := arrayOutput[:len(arrayOutput)-1]

	var GPUInfos []Info
	for _, value := range noHeaderAndTailArray {
		arrayValue := strings.Fields(value)
		memoryUsage, err := strconv.ParseFloat(arrayValue[1], 64)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		memoryTotal, err := strconv.ParseFloat(arrayValue[2], 64)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		GPUInfos = append(GPUInfos, Info{
			Id:          arrayValue[0],
			MemoryUsage: Round((memoryUsage/memoryTotal)*100, 2),
		})
	}

	return GPUInfos
}

type Info struct {
	Id          string
	MemoryUsage float64
}

func Round(num, places float64) float64 {
	shift := math.Pow(10, places)
	return roundInt(num*shift) / shift
}
func roundInt(num float64) float64 {
	t := math.Trunc(num)
	if math.Abs(num-t) >= 0.5 {
		return t + math.Copysign(1, num)
	}
	return t
}
