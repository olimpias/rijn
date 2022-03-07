package log

import (
	"fmt"
	"os"
)

func Fatal(val string) {
	fmt.Println(val)
	os.Exit(1)
}
