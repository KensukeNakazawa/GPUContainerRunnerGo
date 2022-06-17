package gpu

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GetInfo() {
	out, err := exec.Command("nvidia-smi", "--query-gpu={'index','memory.used','memory.total'}", " --format=csv,noheader,nounits").Output()
	if err != nil {
		log.Fatal("not found")
	}

	arrayOutput := strings.Split(string(out), "\n")
	// headerと最後の1行が不要なので削除
	noHeaderAndTailArray := arrayOutput[1 : len(arrayOutput)-1]
	fmt.Println(noHeaderAndTailArray)
}
