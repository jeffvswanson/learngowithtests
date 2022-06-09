package main

import (
	"os"
	"time"

	"github.com/jeffvswanson/learngowithtests/16_maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
