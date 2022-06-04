package main

import (
	"fmt"
	"github.com/v2rayA/beego/v2/logs"
)

func main() {
	l := logs.NewLogger(200)
	params := fmt.Sprintf(`{"filename": "%s", "maxdays": %d}`, "test.log", 3)
	l.SetLogger("file", params)
	l.Warn("test")
}
