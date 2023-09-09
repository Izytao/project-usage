package test

import (
	"project-usage/tools"
	"testing"
)

func TestParseIp(t *testing.T) {
	a := tools.ParseIp("47.242.44.39")
	t.Log(a)
}
