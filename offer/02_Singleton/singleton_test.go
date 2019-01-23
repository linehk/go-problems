package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	e1 := GetInstance()
	t.Log(e1.name)

	e2 := GetInstance()
	t.Log(e2.name)
}
