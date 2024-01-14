package config

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	src := "bind 0.0.0.0\n" +
		"port 6399\n" +
		"appendonly yes\n"
	p := parse(strings.NewReader(src))
	fmt.Println(p)
	// os.Open(filePath)
	file, _ := os.Open("redis.conf")
	p1 := parse(file)
	fmt.Println(p1)
}
