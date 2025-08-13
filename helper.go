package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	suffix     string = "~"
	nbsp       string = " / "
	suffixLine string = strings.Repeat(suffix, 35)
	suffixEnd  string = "\033[0m\n"
)

// color: 1 red, 2 green, 3 yello, 4 blue, 5 purple, 6 blue
func p(color int, str ...any) {
	suffixColor := "\033[3" + strconv.Itoa(color) + "m"
	fmt.Printf("%s%s%s", suffixColor, fmt.Sprint(str...), suffixEnd)
}
