package util

import (
	"fmt"
	"math/rand"
	"time"
)

func MockedIp() string {
	var arr [4]int
	for i := 0; i < 4; i++ {
		rand.NewSource(time.Now().UnixNano())
		arr[i] = rand.Intn(256)
	}
	return fmt.Sprintf("https://%d.%d.%d.%d", arr[0], arr[1], arr[2], arr[3])
}
