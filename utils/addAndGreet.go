package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// prompt for an integer
func promptInt(p string) int {
	var result int64
	fmt.Println(p)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	s := strings.Split(text, "\r")
	result, _ = strconv.ParseInt(s[0], 10, 16)

	return int(result)
}
