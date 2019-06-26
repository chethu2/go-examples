package advance

import (
	"fmt"
	"time"
)

func GoRoutine () {
	 go count(5)
	 time.Sleep(1000) // 1 sec
}

func count (n int) {
	for i :=0;i <n ; i++ {
		fmt.Println(i)
	}
}

