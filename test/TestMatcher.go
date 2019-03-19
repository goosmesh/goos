package main

import (
	"fmt"
	"github.com/goosmesh/goos/core/utils"
)

func main() {
	s := []string{"tes.*", "dfds"}
	fmt.Println(utils.Matcher("test.json", s))
}
