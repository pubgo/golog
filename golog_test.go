package golog

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	cfg := New()
	cfg.Init()

	fmt.Println(GetCfg())
}
